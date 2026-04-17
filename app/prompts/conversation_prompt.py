from app.prompts.system_prompt import get_system_prompt


def build_messages(session, user_message):
    messages = []

    messages.append({
        "role": "system",
        "content": get_system_prompt(session)
    })

    for msg in session.messages[-6:]:
        messages.append(msg)

    messages.append(user_message)

    return messages