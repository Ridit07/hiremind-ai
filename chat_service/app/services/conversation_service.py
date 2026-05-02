from app.core.llm import chat_completion
from app.prompts.system_prompt import get_system_prompt
from app.prompts.evaluation_prompt import get_evaluation_prompt
from app.utils.helpers import extract_topic, safe_json_parse
import time



class ConversationService:

    def process_message(self, session, user_input):
        session.last_activity_time = time.time()
        session.silence_count = 0

        #  1. If waiting for code → treat EVERYTHING as clarification
        if session.current_mode == "waiting_code":
            return self._handle_clarification(session, user_input)


        # 🔥 2. First message → start interview
        if not session.messages:

            # 👉 Decide: coding or normal
            if session.coding_enabled and session.coding_asked < session.coding_required:
                return self._start_coding_question(session)

            question = self._generate_question(session)
            session.last_question = question

            ai_msg = {
                "role": "assistant",
                "type": "question",
                "content": question
            }

            session.messages.append(ai_msg)

            # 🧠 Track topic
            topic_data = extract_topic(question)
            if topic_data.get("topic"):
                session.topics_covered.add(topic_data["topic"])
            if topic_data.get("subtopic"):
                session.subtopics_covered.add(topic_data["subtopic"])

            return ai_msg


        # 👇 Normal flow continues
        session.messages.append({
            "role": "user",
            "type": "text",
            "content": user_input
        })


        # Evaluate answer
        evaluation = self._evaluate(session.last_question, user_input)

        question = session.last_question

        if question not in session.question_attempts:
            session.question_attempts[question] = 0

        if evaluation.get("answer_type") in ["wrong", "unclear", "dont_know"]:
            session.question_attempts[question] += 1

        attempts = session.question_attempts[question]

        feedback = self._generate_feedback(question, user_input, evaluation)


        # 🔥 3. Decide next question (coding vs normal)
        if attempts >= 2:
            feedback = "Let's move to a different topic."

            if session.coding_enabled and session.coding_asked < session.coding_required:
                return self._start_coding_question(session)

            next_question = self._generate_question(session)

        else:
            # normal follow-up logic
            next_question = self._generate_next_question(session, evaluation)


        session.last_question = next_question

        topic_data = extract_topic(next_question)
        if topic_data.get("topic"):
            session.topics_covered.add(topic_data["topic"])
        if topic_data.get("subtopic"):
            session.subtopics_covered.add(topic_data["subtopic"])


        ai_msg = {
            "role": "assistant",
            "type": "response",
            "content": {
                "evaluation": evaluation,
                "feedback": feedback,
                "answer_type": evaluation.get("answer_type", "unknown"),
                "next_question": next_question
            }
        }

        session.messages.append(ai_msg)

        return ai_msg

    def _generate_question(self, session):
        covered_topics = ", ".join(session.topics_covered) or "None"
        covered_subtopics = ", ".join(session.subtopics_covered) or "None"

        prompt = f"""
        Ask an interview question.

        Constraints:
        - Avoid topics already covered: {covered_topics}
        - Avoid subtopics already covered: {covered_subtopics}
        - Keep it relevant to role and skills
        - Keep it concise and realistic
        """

        question = chat_completion([
            {"role": "system", "content": get_system_prompt(session)},
            {"role": "user", "content": prompt}
        ])

        return question

    def _evaluate(self, question, answer):
        messages = [
            {"role": "system", "content": "You are an expert interviewer."},
            {"role": "user", "content": get_evaluation_prompt(question, answer)}
        ]

        response = chat_completion(messages, temperature=0.3)

        return safe_json_parse(response)
    

    def _generate_feedback(self, question, answer, evaluation):
            prompt = f"""
        You are a strict but fair interviewer.

        Question: {question}
        Candidate Answer: {answer}
        Answer Type: {evaluation.get("answer_type")}


        Evaluation:
        Score: {evaluation.get("overall_score")}
        Strengths: {evaluation.get("strengths")}
        Improvements: {evaluation.get("improvements")}

        Generate feedback in MAX 1-2 lines.

        Rules:
        - Be direct and concise
        - No teaching tone (avoid "we can go over together")
        - If wrong → point out mismatch
        - If dont_know → encourage briefly
        - If unclear → say answer is not aligned
        - If correct → short appreciation

        Examples:
        - "That's not aligned with the question. Focus on X."
        - "You're partially correct, but you missed Y."
        - "Good answer, but you missed Z."
        """

            response = chat_completion([
                {"role": "system", "content": "You are a professional interviewer."},
                {"role": "user", "content": prompt}
            ], temperature=0.3, max_tokens=80)

            return response
    
    def _generate_next_question(self, session, evaluation):
        question = session.last_question
        attempts = session.question_attempts.get(question, 0)
        answer_type = evaluation.get("answer_type")

        covered_topics = ", ".join(session.topics_covered) or "None"
        covered_subtopics = ", ".join(session.subtopics_covered) or "None"

        # 🚨 CASE 1: Skip after 2 failed attempts
        if attempts >= 2:
            prompt = f"""
    You are an interviewer.

    Covered Topics: {covered_topics}
    Covered Subtopics: {covered_subtopics}

    Previous Question: {question}
    Candidate failed twice.

    Instructions:
    - Move to a COMPLETELY NEW topic
    - Do NOT repeat or rephrase the previous question
    - Avoid already covered subtopics
    - Ask a fresh, relevant interview question
    """

        # 🚨 CASE 2: Candidate doesn't know
        elif answer_type == "dont_know":
            prompt = f"""
    You are an interviewer.

    Previous Question: {question}

    Instructions:
    - Candidate does not know the answer
    - Give a VERY SMALL hint (one short line)
    - Then ask a simpler related question
    - Do NOT repeat the exact same question
    - Keep it natural and conversational
    """

        # 🚨 CASE 3: Wrong / unclear
        elif answer_type in ["wrong", "unclear"]:
            prompt = f"""
    You are an interviewer.

    Previous Question: {question}

    Instructions:
    - Candidate misunderstood the question
    - Rephrase the SAME question in simpler words
    - Keep it concise and clear
    """

        # 🚨 CASE 4: Partial answer
        elif answer_type == "partial":
            prompt = f"""
    You are an interviewer.

    Previous Question: {question}

    Instructions:
    - Candidate is partially correct
    - Ask a follow-up question on the SAME topic
    - Go slightly deeper
    """

        # 🚨 CASE 5: Correct answer
        else:  # correct
            prompt = f"""
    You are an interviewer.

    Covered Topics: {covered_topics}
    Covered Subtopics: {covered_subtopics}

    Previous Question: {question}

    Instructions:
    - Candidate answered correctly
    - Move to a NEW concept or subtopic
    - Avoid repeating covered subtopics
    - Keep question realistic and concise
    """

        response = chat_completion([
            {"role": "system", "content": get_system_prompt(session)},
            {"role": "user", "content": prompt}
        ], max_tokens=120)

        return response
    


    def _generate_hint(self, question):
        prompt = f"""
    You are an interviewer.

    Question: {question}

    Give a VERY SHORT hint (1 line max).
    Do not reveal the full answer.
    """

        return chat_completion([
            {"role": "system", "content": "You are a helpful interviewer."},
            {"role": "user", "content": prompt}
        ], max_tokens=50)


    def check_inactivity(self, session):
        now = time.time()
        diff = now - session.last_activity_time

        question = session.last_question

        # 🟡 Stage 1: Thinking
        if diff > 15 and session.silence_count == 0:
            session.silence_count += 1
            return {
                "type": "nudge",
                "content": "Take your time."
            }

        # 🟠 Stage 2: Give ACTUAL hint
        elif diff > 30 and session.silence_count == 1:
            session.silence_count += 1

            hint = self._generate_hint(question)

            return {
                "type": "hint",
                "content": hint
            }

        # 🔴 Stage 3: Skip with NEW question
        elif diff > 60:
            session.silence_count = 0

            next_question = self._generate_question(session)
            session.last_question = next_question

            return {
                "type": "skip",
                "content": "Let's move to another question.",
                "next_question": next_question
            }

        return None
    
    def _start_coding_question(self, session):

        # 🎯 Pick difficulty dynamically
        difficulty = None
        if session.coding_difficulty:
            idx = min(session.coding_asked, len(session.coding_difficulty) - 1)
            difficulty = session.coding_difficulty[idx]

        topics = ", ".join(session.coding_topics) or "any relevant topic"

        prompt = f"""
    You are a FAANG-level interviewer.

    Role: {session.role}
    Level: {session.level}

    Ask a coding problem.

    Constraints:
    - Topic preference: {topics}
    - Difficulty: {difficulty or "choose appropriately"}
    - Language preference: {session.preferred_language or "any"}

    Rules:
    - Do NOT give solution
    - Include constraints
    - Keep problem realistic
    - Avoid repeating previous problems
    """

        question = chat_completion([
            {"role": "system", "content": get_system_prompt(session)},
            {"role": "user", "content": prompt}
        ])

        session.current_mode = "waiting_code"
        session.current_coding_question = question
        session.coding_asked += 1

        return {
            "role": "assistant",
            "type": "coding_question",
            "content": question,
            "meta": {
                "difficulty": difficulty,
                "topics": session.coding_topics
            }
        }
    

    def _handle_clarification(self, session, user_input):
        prompt = f"""
    You are a professional technical interviewer.

    Coding Question:
    {session.current_coding_question}

    Candidate Question:
    {user_input}

    Rules:
    - Answer ONLY clarification questions
    - NEVER reveal full algorithm
    - NEVER provide code
    - NEVER explicitly state the optimal approach
    - Small hints are allowed
    - Keep response under 2 lines
    - Maintain interviewer tone
    """

        response = chat_completion([
            {"role": "system", "content": "You are a strict interviewer."},
            {"role": "user", "content": prompt}
        ], max_tokens=80)

        return {
            "role": "assistant",
            "type": "clarification",
            "content": response
        }
    
    def evaluate_code(self, session, data):
        prompt = f"""
    You are a senior engineer evaluating code.

    Problem:
    {session.current_coding_question}

    Code:
    {data['code']}

    Language: {data['language']}

    Candidate Explanation:
    Approach: {data.get('approach')}
    Time Complexity: {data.get('time_complexity')}
    Space Complexity: {data.get('space_complexity')}

    Evaluate STRICTLY.

    Return JSON:
    {{
    "correctness": number (0-10),
    "optimization": number (0-10),
    "code_quality": number (0-10),
    "complexity_accuracy": number (0-10),
    "overall_score": number (0-10),
    "issues": ["..."],
    "improvements": ["..."],
    "verdict": "correct | partially_correct | wrong"
    }}

    Rules:
    - DO NOT rewrite full solution
    - DO NOT provide full optimized code
    - Point issues clearly
    - Be strict
    """

        response = chat_completion([
            {"role": "system", "content": "You are a strict code reviewer."},
            {"role": "user", "content": prompt}
        ], temperature=0.2)

        return safe_json_parse(response)