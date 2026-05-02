from openai import OpenAI
from app.core.config import settings

client = OpenAI(api_key=settings.OPENAI_API_KEY)

def chat_completion(messages, temperature=0.7, max_tokens=200):
    response = client.chat.completions.create(
        model=settings.MODEL,
        messages=messages,
        temperature=temperature,
        max_tokens=max_tokens
    )
    return response.choices[0].message.content.strip()