# Page 1

## Text Content

```
T RANSDUCTION IS A LL YOU N EED FOR
S TRUCTURED DATA W ORKFLOWS

arXiv:2508.15610v1 [cs.AI] 21 Aug 2025

Alfio Gliozzo
IBM Research
gliozzo@us.ibm.com

Naweed Khan
IBM Research
naweed.khan@ibm.com

Nandana Mihindukulasooriya
IBM Research
nnandana@ibm.com

Christodoulos Constantinides
IBM Software, Expert Labs
christodoulos.constantinides@ibm.com

Nahuel Defosse
IBM Research
nahuel.defosse@ibm.com

Junkyu Lee
IBM Research
junkyu.lee@ibm.com

A BSTRACT
This paper introduces Agentics, a modular framework for building agent-based systems capable of
structured reasoning and compositional generalization over complex data. Designed with research
and practical applications in mind, Agentics offers a novel perspective on working with data and AI
workflows. In this framework, agents are abstracted from the logical flow and they are used internally
to the data type to enable logical transduction among data. Agentics encourages AI developers to
focus on modeling data rather than crafting prompts, enabling a declarative language in which data
types are provided by LLMs and composed through logical transduction, which is executed by
LLMs when types are connected. We provide empirical evidence demonstrating the applicability
of this framework across domain-specific multiple-choice question answering, semantic parsing
for text-to-SQL, and automated prompt optimization tasks, achieving state-of-the-art accuracy or
improved scalability without sacrificing performance. The open-source implementation is available at
https://github.com/IBM/agentics.

1

Introduction

Large Language Models (LLMs) have revolutionized how we think about computation, enabling machines to interpret
and generate natural language with remarkable fluency. Yet, despite their success in unstructured settings, LLMs
remain poorly suited for working with structured data which defined by explicit schemas, identifiers, and deterministic
semantics, as commonly found in enterprise systems, databases, and analytics pipelines (Hopkins et al., 2022). Current
agentic AI frameworks address this gap by embedding structured data into natural language prompts, which breaks
down in enterprise use cases such as business analytics or data transformation that demand precision, reproducibility,
and scale (Sarirete et al., 2022; Heck, 2024; Putri and Athoillah, 2024).
Figure 1a illustrates the common GenAI workflow involving Subject Matter Experts (SMEs), LLM agents, and Software
Engineers (SWEs) in a feedback loop: SMEs describe tasks and provide iterative feedback, while agents generate
outputs. The SMEs describes the data transformation task and negotiating requirements with LLM agents using natural
language in multi-turn conversions and provides feedback on the output quality, which iterates until acceptable results
are generated. However, this process is fragile in tasks like API calls, database queries, or data transformation pipelines
that demand precise, verifiable input at scale. Moreover, iterative clarification and validation often lead to errors and
inefficiency. Even in the best-case scenario, agents typically return human-readable summaries or small data samples,
not the structured tabular datasets that SMEs need for further analysis. In practice, business analysts use structured
data that they can manipulate within business intelligence (BI) tools, and therefore, they request supports from SWEs
to scale data workflows in production. Thus, contemporary agentic solutions remain ill-suited for enterprise-grade
structured data operations.
In this paper, we introduce a data-centric paradigm for structured workflows that departs from anthropomorphic
metaphors such as memory, tools, and conversation, which dominate current agentic frameworks. Instead of treating


```

