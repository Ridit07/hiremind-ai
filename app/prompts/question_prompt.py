def get_question_prompt(context, last_entry=None):
    if not last_entry:
        return "Start the interview with a strong opening question."

    return f"""
Previous Question: {last_entry['question']}
Answer: {last_entry['answer']}
Score: {last_entry['score']}

Generate the next question.
If score < 5 → follow-up
If score > 8 → increase difficulty
Else → move to next skill
"""