# Page 6

## Text Content

```
Definition 2 (Agentic Structure AG). An Agentic structure AG is a meta-type that bundles a type schema satype 2 and a
corresponding list of instances, referred to as states sstates :


satype : Θ,
AG :=
sstates : List[Θ]
For clarity, we write AG[X] for an Agentic structure with satype = X, and x = AG[X] for its instance. The i-th state is
accessed via xi = x[i] = x.sstates [i].
In Logical Transduction Algebra, we focus on structured data and its transformation around agents encapsulating LLMs.
The algebraic structure of composing two Agentic instances of the same type can be shown as follows.
Proposition 1 (Monoid of Agentic Instances). Let AG[X] be an Agentic structure and let ξ be the set of all instances
of AG[X]. Define a binary operation ◦ on ξ such that for any x1 , x2 ∈ ξ, their composition x = x1 ◦ x2 is an Agentic
instance whose state list is the concatenation: x.sstates := x1 .sstates ◦ x2 .sstates . Then, the pair (ξ, ◦) forms a monoid,
where the identity element is the Agentic instance with an empty state list: e.sstates := [].
The standard operators follow standard algebraic principles such as the product, x × y for x ∈AG[X] and y ∈AG[Y ],
the equivalence, x1 ∼ x2 for x1 , x2 ∈ AG[X], and the quotient, z/y for z ∈ AG[X ×Y ] and y ∈ AG[Y ], as detailed
in the Appendix.
Equipped with Agentic structures that form a monoid, we obtain a sound abstraction for composing the data workflows
in a functional programming style. This foundation enables the introduction of the logical transduction operator, which
utilizes LLMs as transductive inference engines.
3.3.1

Transduction Operator

We now define a series of overloaded logical transduction operators, which are designed with explicit consideration of
the Agentic structure.
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
In Appendix, we also show more overloaded transduction operators in the case of the few-shot transduction, y[i] =
Y ≪ (P (x[i]) ⊕ F S(x, y)) with a few-shot function F S(x, y) := P ((x′ , y′ )), and a syntactic sugar such as selftransduction.
Next, we formalize key properties of the transduction operator. These properties are foundational for enabling scalable,
parallel, and composable computation.
2

Given two types X and Y , the standard set operations such as union, intersection, complement, and product can be defined
component-wise.

6


```

