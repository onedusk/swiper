# Page 8

## Text Content

```
class AG(BaseModel):
atype: Type[BaseModel]
states: List[BaseModel]
llm: Any

Listing 3: Definition of AG class
3.4.2

PydanticTransducer

Figure 5: Pydantic Transducer Architecture in Agentics
Logical transduction is implemented via the ≪ (left-shift) operator that performs transduction from a source state (or
list of states) to a target AG instance. This operator triggers the creation and execution of a PydanticTransducer,
which is a stateless agent whose sole purpose is to generate a valid instance of the target atype from a textual input.
As shown in Figure 5, it reads the input, performs inference, and populates the target schema with logically consistent
values.
The textual input can represent virtually any concept accessible to an LLM, while the structured output ensures reliability
in downstream tasks. Because Agentics avoids shared conversational memory, it naturally supports asynchronous
execution and efficient scale-out.
To ensure high throughput and responsiveness across datasets of varying sizes, the transduction operator processes
data in configurable batches, typically between 8 and 32 items. Then, the standard transduction pipeline consists of the
following steps: (1) Each source state is serialized according to its Pydantic schema, (2) The prompt is enriched with
few-shot examples and relevant memory passages to guide inference, (3) A PydanticTransducer is instantiated
for the target type, and (4) Transduction is performed in batches using asynchronous calls to the LLM. In the event of a
failure such as rate limits, network interruptions, or parsing errors, we implement a fallback mechanism. The affected
batch is retried synchronously, processing each state individually. This graceful degradation ensures that all data is
processed reliably.
Example 1. The following example illustrates how the AG class and the PydanticTransducer work together. The
former defines the target schema and execution context, while the latter performs the actual transformation.
class Answer(BaseModel):
answer: Optional[str] = None
justification: Optional[str] = None
confidence: Optional[float] = None
async def main():
answers = AG(atype=Answer)
answers = answers << ["What is the capital of Italy?"]
answers = await answers

This will return a structured state with an answer Rome with a justification and a confidence score.
8


```

## Images

![Image from page 8](images/page_8_img_001.ppm)

![Image from page 8](images/page_8_img_002.ppm)

