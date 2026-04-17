import json
from app.redis.client import redis_client

SESSION_PREFIX = "session:"


def save_session(interview_id, session_dict):
    key = SESSION_PREFIX + interview_id
    redis_client.set(key, json.dumps(session_dict))


def get_session(interview_id):
    key = SESSION_PREFIX + interview_id
    data = redis_client.get(key)

    if not data:
        return None

    return json.loads(data)