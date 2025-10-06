# Page 18

## Text Content

```
Appendices
Appendix A
A.1

Logical Transduction Algebra

Formalization

We present an abridged version of the formal Logical Transduction Algebra in the main paper. Our work is closely
related to relational algebra (Codd, 1970) and the MapReduce programming model (Dean and Ghemawat, 2008). This
enables the composition of data transformation pipelines and supports an efficient programming model that leverages
the stateless and asynchronous nature of LLM inference. The full details are provided in the Appendix.
Algebraic Structures
Types and Agentic Structure We define types and meta-types, collectively referred to as the Agentic Structure (AG),
and establish a sound algebra over the types and states within it.
Definition 1 (Types). Let Θ denote the universe of all possible types, Θ = {X, Y, Z, T, . . . }, where each type T ∈ Θ
is a collection of named fields (si , Tsi ):
T := {(s1 , Ts1 ), (s2 , Ts2 ), . . . , (sn , Tsn )},
with each si representing a string-valued slot name, and each Tsi ∈ Θ denoting the corresponding type of that slot.
Definition 8. Given two types X and Y , we define standard set operations component-wise:
X ∪ Y = {(si , Tsi ) | (si , Tsi ) ∈ X or (si , Tsi ) ∈ Y },
X ∩ Y = {(si , Tsi ) | (si , Tsi ) ∈ X and (si , Tsi ) ∈ Y },
X \ Y = {(si , Tsi ) | (si , Tsi ) ∈ X and (si , Tsi ) ∈
/ Y },
X × Y = {((si , Tsi ), (sj , Tsj )) | (si , Tsi ) ∈ X, (sj , Tsj ) ∈ Y }
Definition 2 (Agentic Structure AG). An Agentic structure AG is a meta-type that bundles a type schema satype 6 and a
corresponding list of instances, referred to as states sstates :


satype : Θ,
AG :=
sstates : List[Θ]
Notation conventions: Types are denoted by uppercase letters. Instances of types are denoted by lowercase letters, with
t : T indicating that t is an instance of type T . Lists are written in boldface, so t : T represents a list of instances of
type T . We use the shorthand AG[X] to denote an Agentic structure with satype = X. A boldface lowercase symbol,
such as x = AG[X], represents an instance of AG[X]. We also overload the notation to access the list of states:
xi = x[i] = x.sstates [i] refers to the i-th state of the Agentic instance x.
In Logical Transduction Algebra, we focus on structured data and its transformation around agents encapsulating LLMs.
The algebraic structure of composing two Agentic instances of the same type can be shown as follows.
Proposition 1 (Monoid of Agentic Instances). Let AG[X] be an Agentic structure and let ξ be the set of all instances
of AG[X]. Define a binary operation ◦ on ξ such that for any x1 , x2 ∈ ξ, their composition x = x1 ◦ x2 is an Agentic
instance whose state list is the concatenation: x.sstates := x1 .sstates ◦ x2 .sstates . Then, the pair (ξ, ◦) forms a monoid,
where the identity element is the Agentic instance with an empty state list: e.sstates := [].
Proof. We verify the three monoid properties:
Closure: Let x1 , x2 ∈ ξ. Then x = x1 ◦ x2 has a state list formed by concatenating two valid state lists, which is itself
valid. Hence, x ∈ ξ.
Associativity: For any x1 , x2 , x3 ∈ ξ,
((x1 ◦ x2 ) ◦ x3 ).sstates = (x1 .sstates ◦ x2 .sstates ) ◦ x3 .sstates
= x1 .sstates ◦ (x2 .sstates ◦ x3 .sstates )
= (x1 ◦ (x2 ◦ x3 )).sstates .
6

Given two types X and Y , the standard set operations such as union, intersection, complement, and product can be defined
component-wise.

18


```

