# Page 2

## Text Content

```
(a) Interaction Pattern in Conventional Agentic AI

(b) Proposed Data-Centric Interaction Paradigm in Agentics

Figure 1: Comparison of Interaction Patterns in GenAI Solutions for Data Transformation Tasks

agents as conversational entities, we view them as stateless transducers that operate over well-defined data types.
The main concept we present is logical transduction—the transformation of a data object of type X into type Y by
inferring each field in Y from X under the constraints of a target schema. These transformations can be implemented
by lightweight, stateless agents that are composable, parallelizable, and verifiable.
As shown in Figure 1b, SMEs declaratively define workflows by specifying Pydantic schemas and linking them
via transductions. This no-code approach enables ingestion and transformation of data, optionally guided by fewshot examples, and direct integration with analytics pipelines and Python code. Unlike traditional agents operate
on serialized conversational dependencies, the agents in our approach are stateless and context-free, enabling fully
asynchronous execution. Built on the asynchronous nature of transduction, we also introduce the asynchronous
Map/Reduce programming model, which enables scalable execution of complex workflows that blend LLM inference
with program logic.
Contributions Our key contributions are:
1. formalizing logical transduction and its algebra as the foundation for data-centric agentic AI,
2. introducing Agentics, a framework implementing logical transduction and an asynchronous Map/Reduce
model,
3. demonstrating robustness, accuracy, and scalability across tasks such as domain-specific QA, text-to-SQL, and
prompt optimization for structured workflows,
4. releasing an open-source implementation 1 to support further research and adoption.

2

Agentic AI

For decades, AI has been framed as an effort to emulate human intelligence. The Turing Test exemplifies this anthropomorphic ideal: a machine is deemed intelligent if its conversation is indistinguishable from a human’s. The rise
of LLMs has amplified this framing. By enabling rapid prototyping of intelligent systems through natural language
prompts, developing AI agents has become more intuitive and accessible. Consequently, many agentic AI frameworks
adopt agentic metaphors such as memory, planning, and tool use, often realized via chat-based interfaces.
This approach has driven remarkable progress in consumer applications. Yet, as tasks demand greater semantic precision,
prompt-centric methods often prove brittle, opaque, and hard to scale, especially in structured data environments where
reproducibility and accuracy are paramount. Agentic frameworks typically position the agent as the locus of intelligence,
with data as passive input. While effective for open-ended tasks, this model struggles in deterministic workflows. In
enterprise settings where querying, transformation, and integration of structured data are common, conversational agents
can introduce error propagation and unpredictable behavior.
1

https://github.com/IBM/agentics

2


```

## Images

![Image from page 2](images/page_2_img_001.ppm)

![Image from page 2](images/page_2_img_002.ppm)

![Image from page 2](images/page_2_img_003.ppm)

![Image from page 2](images/page_2_img_004.ppm)

