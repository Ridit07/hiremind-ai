import json
from app.core.llm import chat_completion
from app.prompts.evaluation_prompt import get_evaluation_prompt
from app.utils.helpers import safe_json_parse

def evaluate_answer(question, answer):
    messages = [
        {"role": "system", "content": "You are an expert interviewer."},
        {"role": "user", "content": get_evaluation_prompt(question, answer)}
    ]

    response = chat_completion(messages, temperature=0.3, max_tokens=300)

    parsed = safe_json_parse(response)

    if not parsed:
        return {
            "overall_score": 5,
            "technical_accuracy": 5,
            "communication": 5,
            "depth": 5,
            "strengths": [],
            "improvements": ["Could not evaluate properly"],
            "suggested_answer": ""
        }

    return parsed