# Page 23

## Text Content

```
Proposition 3 (Statelessness). Logical transduction operators are stateless. The output of a transduction yi := Y ≪ xi
depends only on xi and the transduction context, and not on any prior or future transductions.
Proof. By definition, the transduction operator ≪ does not rely on conversational memory or sequential state. Each yi
is computed independently from xi and σ, enabling asynchronous evaluation.
Proposition 4 (Compositionality). Let AG[X], AG[Y ], and AG[Z] be Agentic structures over types X, Y , and Z,
respectively. Suppose y′ = y ≪ x, and z′ = z ≪ y′ . Then the composite transduction holds, z′ = z ≪ y ≪ x.
Proof. Each transduction step applies ≪ component-wise to the state list of the input Agentic structure. Since the
output of y ≪ x is an Agentic structure AG[Y ], it can be used as input to the next transduction. Thus, the composition
is well-defined and yields z′ .
Implications for Distributed and Concurrent Computing These properties make the transduction operator ≪
particularly well-suited for distributed and concurrent computing paradigms. The Conditional Determinism ensures
reproducibility and traceability in distributed pipelines, the Statelessness enables parallel execution, allowing transductions to be mapped across shards of data without coordination or shared state, and the Compositionality supports
modular pipeline construction, akin to functional composition in MapReduce, where intermediate Agentic structures
can be chained and reused.
Asynchronous MapReduce
The programming model of the Agentics supports asynchronous execution of mapping and reduction operations
over Agentic structures, enabling scalable and composable data workflows. We formalize these operations as aMap and
aReduce, which extend the MapReduce by Dean and Ghemawat (2008).
Definition 6 (Asynchronous Map (aMap)). Let AG[X] be an Agentic structure over type X, and let f : X → List[Y ]
be an asynchronous mapping function. Then the asynchronous map operator is defined as:
aMap : (AG[X], f ) → AG[Y ],
S
where the output Agentic structure y = aMap(x, f ) satisfies: y.sstates = i f (xi ), and the union preserving the original
order of inputs.
The function f may return an empty list by removing xi from the output acting as a filter, map each xi to a single output
acting as a transformer, or map each xi to multiple outputs acting as fan-out. Note that aMap operator is executed
asynchronously across all input states, enabling parallelism and scalability in distributed environments.
The aMap operator is executed asynchronously across all input states, enabling parallelism and scalability in distributed
environments.
Definition 7 (Asynchronous Reduce (aReduce)). Let AG[X] be an Agentic structure and let f : List[X] → Y be an
asynchronous reduction function. Then the asynchronous reduce operator is defined as:
aReduce : (AG[X], f ) → AG[Y ],
where the output Agentic structure y = aReduce(x, f ) satisfies: y.sstates = f (x).
Unlike aMap, which applies f to each state individually, aReduce applies f to the entire states x at once. This is
useful for summarization or aggregation, such as generating a report or computing statistics over the full dataset. Since
LLMs have limited context windows, applying aReduce to a large dataset may be intractable. In such cases, scalable
strategies such as hierarchical or batched reduction can be employed by applying aReduce to random subsets and
merging the results.
Composability with Logical Transduction aMap and aReduce can be composed with the logical transduction
operator ≪ to build expressive and modular workflows. For example,
y = aMap(x, x 7→ Y ≪ x),

z = aReduce(y, f ),

where the transduction Y ≪ x is embedded within the mapping function. As we can see, Agentic structure enables
distributed, asynchronous, and semantically typed computation over structured data.
23


```

