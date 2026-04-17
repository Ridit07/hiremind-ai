from app.core.llm import chat_completion
from app.prompts.system_prompt import get_system_prompt
from app.prompts.question_prompt import get_question_prompt

def generate_question(context):
    last_entry = context.history[-1] if context.history else None

    messages = [
        {"role": "system", "content": get_system_prompt(context)},
        {"role": "user", "content": get_question_prompt(context, last_entry)}
    ]

    question = chat_completion(messages)
    context.current_question = question
    return question