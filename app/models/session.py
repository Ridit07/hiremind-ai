import time

class InterviewSession:
    def __init__(self, interview_id, role, company, interview_type, skills, level):
        self.interview_id = interview_id
        self.role = role
        self.company = company
        self.interview_type = interview_type
        self.skills = skills
        self.level = level

        self.messages = []
        self.last_question = None
        self.question_attempts = {}
        self.topics_covered = set()
        self.subtopics_covered = set()
        self.last_activity_time = time.time()

        self.silence_count = 0
        # 🔥 Coding config
        self.coding_enabled = False
        self.coding_required = 0
        self.coding_asked = 0
        self.coding_topics = []
        self.coding_difficulty = []
        self.preferred_language = None
        # runtime
        self.current_mode = "normal"   # normal | coding | waiting_code
        self.current_coding_question = None


    #  Serialize for Redis
    def to_dict(self):
        return {
            "interview_id": self.interview_id,
            "role": self.role,
            "company": self.company,
            "interview_type": self.interview_type,
            "skills": self.skills,
            "level": self.level,
            "messages": self.messages,
            "last_question": self.last_question,
            "question_attempts": self.question_attempts,
            "topics_covered": list(self.topics_covered),
            "subtopics_covered": list(self.subtopics_covered),
            "last_activity_time": self.last_activity_time,
            "silence_count": self.silence_count,

            "coding_enabled": self.coding_enabled,
            "coding_required": self.coding_required,
            "coding_asked": self.coding_asked,
            "coding_topics": self.coding_topics,
            "coding_difficulty": self.coding_difficulty,
            "preferred_language": self.preferred_language,
            "current_mode": self.current_mode,
            "current_coding_question": self.current_coding_question

        }

    #  Deserialize from Redis
    @classmethod
    def from_dict(cls, data):
        obj = cls(
            data["interview_id"],
            data["role"],
            data["company"],
            data["interview_type"],
            data["skills"],
            data["level"]
        )

        obj.messages = data.get("messages", [])
        obj.last_question = data.get("last_question")
        obj.question_attempts = data.get("question_attempts", {})

        obj.topics_covered = set(data.get("topics_covered", []))
        obj.subtopics_covered = set(data.get("subtopics_covered", []))

        obj.last_activity_time = data.get("last_activity_time", time.time())
        obj.silence_count = data.get("silence_count", 0)


        obj.coding_enabled = data.get("coding_enabled", False)
        obj.coding_required = data.get("coding_required", 0)
        obj.coding_asked = data.get("coding_asked", 0)
        obj.coding_topics = data.get("coding_topics", [])
        obj.coding_difficulty = data.get("coding_difficulty", [])
        obj.preferred_language = data.get("preferred_language")
        obj.current_mode = data.get("current_mode", "normal")
        obj.current_coding_question = data.get("current_coding_question")


        return obj