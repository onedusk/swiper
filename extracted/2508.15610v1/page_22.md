# Page 22

## Text Content

```
Transduction Operator Overloading
Definition 3 (Transduction). Given an information object x and a target Agentic structure AG[Y ], the transduction of
x into AG[Y ] is defined as:
[
y := Y ≪ x = (yi , Tsi ),
i

where each yi ∈ Tsi is a value assigned to slot si of type Tsi , logically inferred from x. Here, the operator ≪ denotes a
logical transduction process, implemented via an LLM, that maps x to a structured output conforming to the type Y .
To support the definition of logical transduction operators, we introduce a generic function that renders typed objects
into textual representations suitable for LLM input.
Definition 4 (Prompt Function). Given a type T ∈ Θ, a prompt function P is a mapping that renders a list of states t : T
into an information object, leveraging the string-valued slot names associated with T . Formally, P : List[T ] → str.
Prompt functions serve as a bridge between structured data and natural language, enabling logical transduction operators
to interface with LLMs by converting typed instances into semantically meaningful prompts.
We now define two specific forms of logical transduction: zero-shot and few-shot.
Definition 5 (Zero-Shot Logical Transduction). Let x = AG[X] and y = AG[Y ] be Agentic structures over types X
and Y , respectively. A zero-shot logical transduction from x to y is defined component-wise as:
y[i] = Y ≪ P (x[i]),
where P : X → str is a prompt function that renders each instance x[i] into a textual prompt.
Next, we also show more overloaded transduction operators in the case of the few-shot transduction, y[i] =
Y ≪ (P (x[i]) ⊕ F S(x, y)) with a few-shot function F S(x, y) := P ((x′ , y′ )), and a syntactic sugar such as
self-transduction.
Definition 13 (Few-Shot Logical Transduction). Let x = AG[X] and y = AG[Y ] be Agentic structures over types X
and Y , respectively. A few-shot logical transduction from x to y is defined for all indices i such that y[i] = ∅ as:
y[i] = Y ≪ (P (x[i]) ⊕ F S(x, y)) ,
where:
• P : X → str is a prompt function that renders an instance x[i] into a textual prompt.
• ⊕ denotes prompt concatenation.
• F S(x, y) is the few-shot context, defined as:
F S(x, y) := P ((x′ , y′ )) ,
where (x′ , y′ ) is the projection of (x, y) onto the subset of indices for which y[i] ̸= ∅.
Finally, we introduce self-transduction as syntactic sugar within the programming model for logical transductions.
Definition 14 (Self Transduction). Let x ∈ AG[X] be an Agentic structure, and let Y, Z ⊂ X be two disjoint subsets
of types. A self transduction is a function that produces a modified Agentic structure x′ ∈ AG[X], defined as:
x′ = x ≪Y,Z := x ∪ (x[Y ] ≪ x[Z]) ,
where x[Y ] denotes the rebind operator, which extracts an AG[Y ] from x by retaining only the slots in Y that overlap
with X.
Properties of Transduction Operator Next, we formalize key properties of the transduction operator. These
properties are foundational for enabling scalable, parallel, and composable computation.
Proposition 2 (Conditional Determinism). Let σ denote a fixed transduction context, which may include components
such as a few-shot context, additional instructions, external tools, or memory. Let the LLM configuration—comprising
model weights, temperature, and decoding strategy—also be fixed. Then, for any input xi , the transduction yi := Y ≪
xi is deterministic.
Proof. When the model parameters and transduction context σ are fixed, and the LLM is configured with deterministic
settings (e.g., temperature set to zero and caching enabled), the output yi is uniquely determined by the input xi and the
context σ. Therefore, the transduction process is deterministic under these conditions.
22


```

