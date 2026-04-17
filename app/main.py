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


@app.get("/check-activity")
def check_activity(interview_id: str):
    session = sessions.get(interview_id)

    if not session:
        return {}

    result = service.check_inactivity(session)

    if result:
        return {
            "role": "assistant",
            "type": result["type"],
            "content": result["content"],
            "next_question": result.get("next_question")
        }

    return {}