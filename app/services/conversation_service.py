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

        # Evaluate answer
        evaluation = self._evaluate(session.last_question, user_input)

        # Generate next question
        next_question = self._generate_question(session)

        session.last_question = next_question

        ai_msg = {
            "role": "assistant",
            "type": "response",
            "content": {
                "evaluation": evaluation,
                "next_question": next_question
            }
        }

        session.messages.append({
            "role": "user",
            "type": "text",
            "content": user_input
        })

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