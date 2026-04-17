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