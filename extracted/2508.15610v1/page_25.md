# Page 25

## Text Content

```
option_ids: List[str] = Field(description="List of option identifiers")
asset_name: str = Field(description="Name of the industrial asset")
relevancy: str = Field(description="Relevancy context or metadata")
question_type: str = Field(description="Type of question")
subject: str = Field(description="Subject or topic of the question")
system_answer: Answer = Field(description="Model-generated answer object")

Listing 5: Data Model for Domain-Specific Multiple Choice QA
Main Algorithm The workflow begins by instantiating an AG object with the FailureSensorIQ schema and
a specified batch size. Each example from the dataset is parsed into a structured FailureSensorIQ instance and
appended to the agent’s internal state.
The core operation is the self transduction, which performs schema-guided inference over the specified input fields—such
as question, options, and asset name—and generates structured outputs in the system answer field. The transduction is
guided by a natural language instruction that defines the task: selecting the most plausible answer from a set of options,
along with a confidence score and rationale.
# Initialize the \texttt{Agentics} benchmark with the FailureSensorIQ schema and batch
size
fsiq_benchmark = AG(atype=FailureSensorIQ, batch_size=40)
# Load dataset and populate agent states
dataset = load_dataset("cc4718/FailureSensorIQ")
for example in dataset:
fsiq_benchmark.states.append(FailureSensorIQ(**example))
# Run self-transduction with structured input and output fields
fsiq_benchmark = await fsiq_benchmark.self_transduction(
input_fields=[
"question", "options", "option_ids",
"asset_name", "relevancy", "question_type", "subject"
],
output_fields=["system_answer"],
instructions=(
"Read the input questions, all possible answers, and background task
information. "
"This is a multiple choice test, where one of the options is true and the
others are false. "
"Select the answer with the highest likelihood of being correct, and return it
along with "
"a confidence score and a verbal assessment explaining your judgment."
)
)

Listing 6: Pseudo Code for Domain-Specific Multiple Choice QA
B.2

Prompt Optimization

Prompt optimization is a critical component in leveraging large language models (LLMs) for complex tasks. The
performance of LLMs is highly sensitive to variations in prompt structure, tone, and even the positioning of textual
components. Minor changes—such as rephrasing imperative sentences or reordering blocks—can significantly impact
the model’s output.
Prompts typically follow a structured format, often divided into system and user sections. The system prompt provides
general context, instructions, and expectations for the output, while the user prompt contains task-specific information.
Common elements include task descriptions, constraints, few-shot examples, input/output format specifications, and
guiding phrases like “Let’s think step by step”. These components are frequently organized using structured formats
such as Markdown or JSON schemas.
General Framework for Prompt Optimization Ramnath et al. (2025) and Li et al. (2025) have summarized existing
prompt optimization techniques into a generic prompt optimization framework. Algorithm 1 presented in Ramnath et al.
25


```

