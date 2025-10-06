# Page 5

## Text Content

```
3.2

Concepts

Logical transduction is a structured, and inference-driven transformation of a data object x of type X into a new object
y of type Y , such that each field in y can be logically explained based on the information in x. This transformation is
grounded in the semantics of the source and target types, and guided by the constraints encoded in their schemas.
Figure 3 illustrates a simple example, where ProductReview object containing a reviewer’s name, review text, and rating
is transduced into a new object that includes a sentiment label and a rationale. The LLM infers new fields based on
source and target schemas, filling structured templates with semantically meaningful content. The target schema guides
the transduction by catching malformed generations via type validation, filling with defaults in the missing fields, and
introducing optional fields that could support downstream tasks or improve the interpretability.

Figure 4: Logical Transduction as Negotiation of Meaning
Beyond a single transformation, logical transduction can be viewed as a negotiation of meaning between two agents,
as illustrated in Figure 4. The source agent, AG[X], which holds data of type X, inspects the target agent AG[Y ]’s
schema, with data type Y , to understand its requirements. The source agent then constructs a prompt that selects or
renders relevant information from the source state x and adds task-specific instructions. This prompt is applied to x,
and messages are sent to the target, yielding states y asynchronously. The target may provide few-shot feedback from
the transduced states for further refinement. This agentic negotiation supports high levels of parallelism, as messages
are processed in asynchronous batches.
3.3

Formalization

We present an abridged version of the formal Logical Transduction Algebra in the main paper. Our work is closely
related to relational algebra (Codd, 1970) and the MapReduce programming model (Dean and Ghemawat, 2008). This
enables the composition of data transformation pipelines and supports an efficient programming model that leverages
the stateless and asynchronous nature of LLM inference. The full details are provided in the Appendix.
We define types and meta-types, collectively referred to as the Agentic Structure (AG), and establish a sound algebra
over the types and states within it.
Definition 1 (Types). Let Θ denote the universe of all possible types, Θ = {X, Y, Z, T, . . . }, where each type T ∈ Θ
is a collection of named fields (si , Tsi ):
T := {(s1 , Ts1 ), (s2 , Ts2 ), . . . , (sn , Tsn )},
with each si representing a string-valued slot name, and each Tsi ∈ Θ denoting the corresponding type of that slot.
Types are denoted by uppercase letters and instances by lowercase, e.g., t : T indicating t is of type T . Lists use
boldface, so t : T is a list of instances of type T .
5


```

## Images

![Image from page 5](images/page_5_img_001.ppm)

![Image from page 5](images/page_5_img_002.ppm)

