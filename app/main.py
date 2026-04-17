from fastapi import FastAPI
from app.models.session import InterviewSession
from app.services.conversation_service import ConversationService

app = FastAPI()
service = ConversationService()

sessions = {}


@app.post("/chat")
def chat(data: dict):
    interview_id = data["interview_id"]

    if interview_id not in sessions:
        sessions[interview_id] = InterviewSession(
            interview_id=interview_id,
            role=data["role"],
            company=data["company"],
            interview_type=data["type"],
            skills=data["skills"],
            level=data["level"]
        )

    session = sessions[interview_id]

    user_input = data["message"]

    response = service.process_message(session, user_input)

    return response