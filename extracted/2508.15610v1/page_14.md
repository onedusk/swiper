# Page 14

## Text Content

```
Model
llama-3-3
llama-4-maverick
Method
mistral-large
70b-instruct
17b-128e-instruct-fp8
P
50.51 ± 0.71 47.20 ± 0.78
53.88 ± 1.37
50.32 ± 0.52 45.63 ± 0.9
53.47 ± 0.56
P+FS
(−0.19)
(−1.57)
(−0.41)
50.32 ± 0.32 45.09 ± 0.81
53.03 ± 0.82
P+KW
(−0.19)
(−2.11)
(−0.85)
51.58 ± 0.58 46.15 ± 0.44
52.33 ± 0.62
P+SQ
(+1.07)
(−1.05)
(−1.55)
52.75 ± 0.07 49.54 ± 1.23
50.46 ± 0.05
P+SL
(+2.24)
(+2.34)
(−3.42)
54.90 ± 0.53 48.93 ± 0.5
53.15 ± 0.65
P+OP
(+4.39)
(+1.73)
(−0.74)
60.84 ± 0.7
55.41 ± 1.2
56.94 ± 0.23
P+FS+KW+SQ+SL+OP
(+10.33)
(+8.21)
(+3.06)
Table 2: Execution accuracy on BIRD-dev, testing a prompt P with the simplified task vs. additional transductions in the
composite workflow. SL replaces the full DDL schema with a linked schema, KW includes keyword topic modeling,
FS randomly generates sql-validated few-shot question-query pairs, SQ extracts sub-questions, and OP optimizes the
prompt template.

Agentics framework allows for more sophisticated augmentations and feedback loops to improve the creation of
such programs. Future attempts at modeling more complex prompt programs include the addition of feedback and
self-correction loops by running failed samples on the database and re-prompting. Additionally, model-dependent
optimizations can be illuminated upon, since the performance of the executed program is dependent on the model itself
and should therefore be guided by the model more closely.

5

Comparison of Agentic AI Frameworks

We compare our system, Agentics, with five representative frameworks: LangGraph, CrewAI, PydanticAI,
AutoGen, and DSPy. Unlike these systems, which often foreground agent behavior and control flow, Agentics is
centered on semantic modeling of data and logical transduction.
• LangGraph provides fine-grained control over agent states and transitions via graph-based workflows, but
requires manual wiring and memory management. In contrast, Agentics eliminates explicit workflow
construction by leveraging Pydantic types and algebraic transduction operators (≪), modeling logic internally
through structured schemas.
• CrewAI introduces a team-based model with autonomous agents communicating via messages. While this
supports collaborative multi-agent systems, it introduces coordination overhead. Agentics avoids this by
using a type-driven architecture where transductions are declaratively executed based on semantic alignment
between source and target structures.
• PydanticAI focuses on type-safe LLM interaction using Pydantic models. Agentics builds on this
foundation but extends it into a full transduction system, treating datasets as semantic structures and enabling
composable pipelines for expressive data manipulation.
• AutoGen simplifies orchestration through high-level abstractions but sacrifices fine-grained data modeling and
type safety. Agentics offers stronger guarantees by validating data against Pydantic types and embedding
semantic hints directly into schemas, reducing reliance on prompt engineering.
• DSPy emphasizes prompt optimization and imperative control flows, but its fixed data model and reliance
on verbose string templates limit flexibility. Agentics, by contrast, supports dynamic type creation, batch
processing, and external execution logic—making it more suitable for structured tasks like text-to-SQL.
While recent agentic AI frameworks have made significant strides in enabling modular, prompt-driven, and multi-agent
orchestration, several limitations remain when compared to the design of Agentics.
• Many frameworks require developers to explicitly define logical control flow through specialized modules or
components. These often rely on string-based prompt manipulation, which, although sometimes abstracted
14


```

