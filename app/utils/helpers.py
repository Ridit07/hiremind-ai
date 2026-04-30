import json
from app.core.llm import chat_completion



# def safe_json_parse(response: str, default=None):
#     try:
#         return json.loads(response)
#     except Exception:
#         return default or {}


def calculate_average_score(history):
    if not history:
        return 0
    return sum(item["score"] for item in history) / len(history)


def get_difficulty_mode(avg_score):
    if avg_score < 5:
        return "easy"
    elif avg_score > 8:
        return "hard"
    return "normal"


def get_recent_history(history, limit=3):
    return history[-limit:] if history else []


def format_history_for_prompt(history):
    if not history:
        return "No previous interactions."

    formatted = []
    for item in history:
        formatted.append(
            f"Q: {item['question']}\nA: {item['answer']}\nScore: {item['score']}"
        )

    return "\n\n".join(formatted)


def safe_json_parse(response):
    try:
        return json.loads(response)
    except:
        return {}
    
def extract_topic(question):
    prompt = f"""
Classify this interview question.

Question: {question}

Return JSON:
{{
  "topic": "DSA | Python | System Design | DBMS | OS",
  "subtopic": "specific concept like Linked List, HashMap, Tuple"
}}
"""

    response = chat_completion([
        {"role": "system", "content": "You are a classifier."},
        {"role": "user", "content": prompt}
    ], temperature=0)

    return safe_json_parse(response)




def apply_coding_config(session, coding_config):
    if not coding_config:
        return

    session.coding_enabled = coding_config.get("enabled", False)

    if not session.coding_enabled:
        return

    session.coding_required = coding_config.get("num_questions", 1)

    # normalize difficulty → always list
    difficulty = coding_config.get("difficulty", [])
    if isinstance(difficulty, str):
        difficulty = [difficulty]

    session.coding_difficulty = difficulty
    session.coding_topics = coding_config.get("topics", [])
    session.preferred_language = coding_config.get("language")