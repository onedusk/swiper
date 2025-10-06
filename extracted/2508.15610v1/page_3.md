# Page 3

## Text Content

```
We argue that these limitations stem not from LLM capabilities, but from the design paradigms embedded in current
frameworks that prioritize human-like behavior over the strengths of traditional computing, which manages structured,
semantically rich data with precision and reliability. Agentics proposes a shift towards computation centered on data
semantics and type-driven transformations, enabling more robust, scalable, and interpretable workflows.
2.1

Concepts

Recent demands in Agentic AI have led to a surge of interest in frameworks and methodologies for building intelligent
systems composed of multiple interacting agents. Several foundational works provide context for this development.
• Han et al. (2024) propose a unified framework for LLM-based multi-agent systems, structured around five
components: profile, perception, self-action, mutual interaction, and evolution. They highlight the potential
of LLMs to simulate human-like intelligence through coordinated agent behavior, while identifying key
challenges such as memory management and scalability.
• Huang and Huang (2025) explore collaborative problem-solving in multi-agent networks, emphasizing the role
of LLMs in enabling dynamic communication and negotiation. Their study distinguishes between cooperative
and competitive agent behaviors and demonstrates the flexibility of LLMs in real-world coalition formation.
• Moshkovich and Zeltyn (2025) focus on the lifecycle of agentic systems, introducing tools for observability,
diagnostics, and optimization. They argue that as agentic workflows grow in complexity, traditional debugging
methods fall short, and propose structured metrics for alignment, coherence, and task success.
• Acharya et al. (2025) present a comprehensive survey of agentic AI, covering autonomy, memory, planning,
and tool use across various domains. They discuss coordination strategies, ethical considerations, and the
societal impact of deploying autonomous agents at scale.
Together, these works underscore the growing importance of agentic AI as a paradigm for building intelligent, goaldriven systems. They also highlight the limitations of current frameworks in terms of flexibility, transparency, and
integration with structured data workflows—motivating the need for alternatives such as Agentics, which emphasize
semantic modeling and logical transduction.
2.2

Software Frameworks

Recent years have seen the emergence of a variety of frameworks for building agentic AI systems, each adopting a
different philosophy in orchestrating agent behavior, workflow logic, and semantic alignment. These systems aim to
harness the capabilities of large language models (LLMs) for complex reasoning, structured data manipulation, and
autonomous task execution.
To address the limitations of prompt-centric systems, recent research has introduced techniques such as guardrails
(Ouyang et al., 2022; Dong et al., 2024; Zhang et al., 2024), self-reflection (Shinn et al., 2023; Asai et al., 2024), and
correction strategies (Madaan et al., 2023; Pan et al., 2024). While these methods enhance reliability, most frameworks
still depend on free-form text or loosely structured prompts that remain fragile and difficult to verify—especially in
tasks requiring high semantic precision.
In response, a new generation of frameworks has emerged that emphasize structured data computation, type safety, and
modular orchestration. These include:
• LangGraph (LangChain, 2025), a graph-based orchestration system supporting stateful agents and structured
data flows
• CrewAI (CrewAI Inc., 2025), a multi-agent coordination framework integrating Pydantic for structured task
definitions
• Pydantic-AI (Pydantic, 2025), a data-centric agent framework built around Pydantic schemas
• DSPy (Khattab et al., 2024), a declarative abstraction for prompt engineering and optimization
• AutoGen (Wu et al., 2024), a flexible multi-agent framework supporting tool use, planning, and structured
messaging
• SmolAgents (SmolAI, 2024), a lightweight agentic framework for simplicity
• Atomic Agents (AtomicAI, 2024), a modular, type-safe agentic system for real-time observability.
These frameworks reflect a growing movement toward more interpretable, modular, and data-aware agentic systems
that lay the foundation for robust, enterprise AI workflows.
3


```

