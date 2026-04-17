from app.core.llm import chat_completion
from app.prompts.system_prompt import get_system_prompt
from app.prompts.evaluation_prompt import get_evaluation_prompt
from app.utils.helpers import safe_json_parse


class ConversationService:

    def process_message(self, session, user_input):

        # First message → start interview
        if not session.messages:
            question = self._generate_question(session)
            session.last_question = question

            ai_msg = {
                "role": "assistant",
                "type": "question",
                "content": question
            }

            session.messages.append(ai_msg)
            return ai_msg
        
        session.messages.append({
        "role": "user",
        "type": "text",
        "content": user_input
        })


        # Evaluate answer
        evaluation = self._evaluate(session.last_question, user_input)
        
        # Feedback
        feedback = self._generate_feedback(session.last_question, user_input, evaluation)

        # Generate next question
        next_question = self._generate_next_question(session, evaluation)

        session.last_question = next_question

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

        # session.messages.append({
        #     "role": "user",
        #     "type": "text",
        #     "content": user_input
        # })

        session.messages.append(ai_msg)

        return ai_msg

    def _generate_question(self, session):
        messages = [
            {"role": "system", "content": get_system_prompt(session)},
            {"role": "user", "content": "Ask the next interview question."}
        ]

        return chat_completion(messages)

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
        answer_type = evaluation.get("answer_type")

        prompt = f"""
        You are an interviewer.

        Previous Question: {session.last_question}
        Candidate Score: {evaluation.get("overall_score")}
        Answer Type: {answer_type}

        Rules:

        - If answer_type = wrong or unclear:
            → Repeat the SAME question in simpler words

        - If answer_type = dont_know:
            → Give hint + ask same question again

        - If answer_type = partial:
            → Ask follow-up on SAME topic

        - If answer_type = correct:
            → Move to next concept

        Keep it concise and natural.

        """

        response = chat_completion([
            {"role": "system", "content": get_system_prompt(session)},
            {"role": "user", "content": prompt}
        ], max_tokens=120)

        return response