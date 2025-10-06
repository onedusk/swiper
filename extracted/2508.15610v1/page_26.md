# Page 26

## Text Content

```
(2025), formalizes the process of prompt optimization as follows. Given a task model Mtask , an initial prompt ρ ∈ V ,
the goal of an prompt optimization system MP O is to obtain the best performing prompt-template ρopt under a metric
f ∈ F and eval-set Dval that maximizes expected performance:
ρopt := arg max Ex∼Dval [f (Mtask (ρ ⊕ x))] .
ρ∈V

Since the objective function is not tractable due to the combinatorial nature of discrete token-sequence search spaces,
the optimization process typically follows a generate-select-evaluate cycle, akin to local search algorithms (Russell and
Norvig, 2020). Most methods begin with a predefined prompt template that specifies the structure and content to be
included. An initial prompt may be constructed manually or generated automatically using an LLM.
Given a specific task, the dataset is usually partitioned into training and validation sets. The training set is used to
optimize the prompt, while the validation set is employed for tuning hyperparameters. Additionally, a held-out test set
is used to evaluate final performance. Candidate prompts are generated, filtered based on performance metrics, and
refined over successive iterations, with incremental improvements guided by feedback from model outputs.
Algorithm 1: Prompt Optimization Framework
1: P0 := {ρ1 , ρ2 , . . . , ρk }
2: Dval := {(x1 , y1 )}n
i=1
3: f1 , . . . , fm ∈ F
4: for t = 1, 2, . . . , N do
5:
Gt := MP O (P, Dval , F )
6:
Pt := Select(Gt , Dval , F )
7:
if fconvergence ≤ ϵ then
8:
exit
9: return arg maxρ∈PN Ex∼Dval [f (Mtask (ρ ⊕ x))]

▷ Initial seed prompts
▷ Validation set
▷ Inference evaluation
▷ Iteration depth
▷ Generate prompt candidates with MP O
▷ Filter and retain candidates
▷ Optionally check for early convergence

Parallelizing Prompt Optimization with Transduction Algebra Within the Agentics framework, the prompt
function defined in Definition 4 maps type information into structured information objects for transduction. Prompt
optimization in this context is naturally expressed using transduction algebra. Since candidate prompts can be generated
by logical transduction, we adopt meta-prompt-based optimization strategies similar to those proposed in (Yang et al.,
2024; Ye et al., 2024; Opsahl-Ong et al., 2024). Our approach emphasizes two key aspects:
• The Agentics framework supports a declarative style of prompting, where prompts are constructed to
encode rich contextual and type-level information rather than procedural instructions.
• The optimization process follows the generate-select-evaluate cycle described in Algorithm 1. Importantly,
the transduction algebra enables this optimization loop to be expressed in a functional and parallelizable
manner. Prompt candidates can be generated and evaluated independently, allowing for efficient execution.
This abstraction not only improves scalability but also decouples the optimization logic from the underlying
execution strategy.
In summary, the Agentics framework provides a principled and extensible foundation for prompt optimization. It
integrates declarative prompt construction, transduction algebra, and parallel search strategies into a unified system that
supports both expressiveness and scalability. Next, we present a functional design pattern for implementing prompt
optimization using transduction algebra, abstracting away procedural details common in existing approaches. In Section
4.3, our experiments demonstrate that declarative context optimization improves performance and that parallelization
yields substantial runtime gains.
Data Models We begin by defining two data models using Pydantic: OptimizationTask and GSM8K. The
OptimizationTask schema captures the components of a prompt template—such as role, goal, expected output,
and imperative instructions—along with a score field for evaluation. The GSM8K schema represents the target task,
including the question, ground-truth answer, model-generated reasoning, and correctness flag.
class OptimizationTask(BaseModel):
# a list of demo tasks used to guide prompt generation
demos: list[Any] = Field(description="optimization demo tasks")
# role description to be embedded in the prompt (e.g., ’You are a math tutor’)
role: str = Field(description="New role instruction")

26


```

