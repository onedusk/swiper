# Page 10

## Text Content

```
Building Effective AI Agents \ Anthropic

7/9/25, 4:18 PM

Agents are emerging in production as LLMs mature in key
capabilities—understanding complex inputs, engaging in reasoning
and planning, using tools reliably, and recovering from errors.
Agents begin their work with either a command from, or interactive
discussion with, the human user. Once the task is clear, agents plan
and operate independently, potentially returning to the human for
further information or judgement. During execution, it's crucial for
the agents to gain “ground truth” from the environment at each step
(such as tool call results or code execution) to assess its progress.
Agents can then pause for human feedback at checkpoints or when
encountering blockers. The task often terminates upon completion,
but it’s also common to include stopping conditions (such as a
maximum number of iterations) to maintain control.
Agents can handle sophisticated tasks, but their implementation is
often straightforward. They are typically just LLMs using tools based
on environmental feedback in a loop. It is therefore crucial to design
toolsets and their documentation clearly and thoughtfully. We
expand on best practices for tool development in Appendix 2
("Prompt Engineering your Tools").

Autonomous agent

When to use agents: Agents can be used for open-ended problems
where it’s difficult or impossible to predict the required number of
https://www.anthropic.com/engineering/building-effective-agents

Page 10 of 16


```

## Images

![Image from page 10](images/page_10_img_001.ppm)

![Image from page 10](images/page_10_img_002.ppm)

