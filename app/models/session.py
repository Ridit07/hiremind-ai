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
        self.asked_questions = set()
        self.last_activity_time = time.time()
        self.silence_count = 0