def get_code_evaluation_prompt(question, code):
    return f"""
Evaluate this coding answer.

Question: {question}
Code:
{code}

Return JSON:
{{
  "score": number,
  "correctness": "good/average/poor",
  "time_complexity": "O(...)",
  "issues": ["..."],
  "improvements": ["..."]
}}
"""