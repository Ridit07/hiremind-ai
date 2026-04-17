def get_evaluation_prompt(question, answer):
    return f"""
You are an expert technical interviewer.

Question: {question}
Candidate Answer: {answer}

Evaluate on:
1. Technical Accuracy (correctness)
2. Communication (clarity and structure)
3. Depth (level of understanding)

Return STRICT JSON:

{{
  "overall_score": number (0-10),
  "technical_accuracy": number (0-10),
  "communication": number (0-10),
  "depth": number (0-10),
  "strengths": ["point1", "point2"],
  "improvements": ["point1", "point2"],
  "suggested_answer": "Provide a better version of the answer in 3-4 lines"
  "answer_type": "correct | partial | wrong | dont_know | unclear"
}}

Rules:
- If answer is unrelated → answer_type = wrong
- If user misunderstood question → answer_type = unclear
- If user says "I don't know" → answer_type = dont_know
- If partially correct → answer_type = partial
- If good → answer_type = correct
- Be strict in classification
- Be concise
- Do not add extra text outside JSON
"""