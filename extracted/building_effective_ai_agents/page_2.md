# Page 2

## Text Content

```
Building Effective AI Agents \ Anthropic

7/9/25, 4:18 PM

"Agent" can be defined in several ways. Some customers define
agents as fully autonomous systems that operate independently over
extended periods, using various tools to accomplish complex tasks.
Others use the term to describe more prescriptive implementations
that follow predefined workflows. At Anthropic, we categorize all
these variations as agentic systems, but draw an important
architectural distinction between workflows and agents:
Workflows are systems where LLMs and tools are orchestrated
through predefined code paths.
Agents, on the other hand, are systems where LLMs dynamically
direct their own processes and tool usage, maintaining control
over how they accomplish tasks.
Below, we will explore both types of agentic systems in detail. In
Appendix 1 (“Agents in Practice”), we describe two domains where
customers have found particular value in using these kinds of
systems.

When (and when not) to use
agents
When building applications with LLMs, we recommend finding the
simplest solution possible, and only increasing complexity when
needed. This might mean not building agentic systems at all. Agentic
systems often trade latency and cost for better task performance, and
you should consider when this tradeoff makes sense.
When more complexity is warranted, workflows offer predictability
and consistency for well-defined tasks, whereas agents are the better
option when flexibility and model-driven decision-making are
needed at scale. For many applications, however, optimizing single
LLM calls with retrieval and in-context examples is usually enough.

https://www.anthropic.com/engineering/building-effective-agents

Page 2 of 16


```

