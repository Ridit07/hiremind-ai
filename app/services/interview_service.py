from app.services.question_service import generate_question
from app.services.evaluation_service import evaluate_answer

class InterviewService:

    def start_interview(self, context):
        return generate_question(context)

    def submit_answer(self, context, answer):
        question = context.current_question

        evaluation = evaluate_answer(question, answer)

        context.history.append({
            "question": question,
            "answer": answer,
            "score": evaluation["overall_score"]
        })

        next_question = generate_question(context)

        return {
            "evaluation": evaluation,
            "next_question": next_question
        }