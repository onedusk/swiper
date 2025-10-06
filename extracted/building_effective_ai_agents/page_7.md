# Page 7

## Text Content

```
Building Effective AI Agents \ Anthropic

7/9/25, 4:18 PM

The parallelization workflow

When to use this workflow: Parallelization is effective when the
divided subtasks can be parallelized for speed, or when multiple
perspectives or attempts are needed for higher confidence results.
For complex tasks with multiple considerations, LLMs generally
perform better when each consideration is handled by a separate
LLM call, allowing focused attention on each specific aspect.
Examples where parallelization is useful:
Sectioning:
Implementing guardrails where one model instance
processes user queries while another screens them for
inappropriate content or requests. This tends to perform
better than having the same LLM call handle both guardrails
and the core response.
Automating evals for evaluating LLM performance, where
each LLM call evaluates a different aspect of the model’s
performance on a given prompt.
Voting:
Reviewing a piece of code for vulnerabilities, where several
different prompts review and flag the code if they find a
problem.
Evaluating whether a given piece of content is inappropriate,
with multiple prompts evaluating different aspects or
https://www.anthropic.com/engineering/building-effective-agents

Page 7 of 16


```

## Images

![Image from page 7](images/page_7_img_001.ppm)

![Image from page 7](images/page_7_img_002.ppm)

