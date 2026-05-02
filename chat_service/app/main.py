from fastapi import FastAPI
from app.models import session
from app.models.session import InterviewSession
from app.services.conversation_service import ConversationService
from app.redis.session_store import delete_session, save_session, get_session
from app.redis.lock import acquire_lock, release_lock
import time
from app.utils.helpers import apply_coding_config


app = FastAPI()
service = ConversationService()

sessions = {}


@app.post("/chat")
def chat(data: dict):
    interview_id = data["interview_id"]

    lock = acquire_lock(interview_id)  

    if not lock:
        return {"error": "Too many requests. Please try again."}
    
    try:

        session_data = get_session(interview_id)


        if session_data:
            session = InterviewSession.from_dict(session_data)
        else:
            session = InterviewSession(
                interview_id=interview_id,
                role=data["role"],
                company=data["company"],
                interview_type=data["type"],
                skills=data["skills"],
                level=data["level"]
            )
        apply_coding_config(session, data.get("coding"))

        # session = sessions[interview_id]
        session.last_activity_time = time.time()
        session.silence_count = 0


        user_input = data["message"]

        response = service.process_message(session, user_input)

        save_session(interview_id, session.to_dict())


        return response
    
    finally:
        release_lock(lock)


@app.get("/check-activity")
def check_activity(interview_id: str):

    lock = acquire_lock(interview_id)  

    if not lock:
        return {}

    try:
        session_data = get_session(interview_id)


        if not session_data:
            return {}

        session = InterviewSession.from_dict(session_data)

        result = service.check_inactivity(session)

        if result:
            save_session(interview_id, session.to_dict())
            return {
                "role": "assistant",
                "type": result["type"],
                "content": result["content"],
                "next_question": result.get("next_question")
            }

        return {}
    finally:
        release_lock(lock)



@app.post("/submit-code")
def submit_code(data: dict):

    interview_id = data["interview_id"]

    session_data = get_session(interview_id)

    if not session_data:
        return {"error": "Session not found"}

    session = InterviewSession.from_dict(session_data)

    evaluation = service.evaluate_code(session, data)

    # coding round finished
    session.current_mode = "normal"
    session.current_coding_question = None

    # 🔥 Continue flow
    if (
        session.coding_enabled and
        session.coding_asked < session.coding_required
    ):
        next_question = service._start_coding_question(session)

    else:
        next_question = service._generate_question(session)

    save_session(interview_id, session.to_dict())

    return {
        "evaluation": evaluation,
        "next_question": next_question
    }



@app.delete("/end-session")
def end_session(interview_id: str):
    delete_session(interview_id)
    return {"message": "Session deleted"}

@app.get("/")
def home():
    return {"message": "App is running 🚀"}

@app.get("/health")
def health():
    return {"status": "healthy"}