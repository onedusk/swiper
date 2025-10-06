# Page 29

## Text Content

```
* Given the previous optimization results, don’t generate duplicate or similar prompt
templates.
* Generate prompt template that achieves the best score, and succint and concise
instructions.
"
# Template used to instantiate a user prompt from the optimized parameters.
# This is applied to each validation example.
USER_PROMPT_TEMPLATE = "
You are {role}.
Your personal goal is: {goal}.
This is the expected criteria for your final answer: {expected_output}.
This is the general task context.
{task_context}
solve the following task.
{question}
{options}
{option_ids}
{asset_name}
{relevancy}
{question_type}
{subject}
{imperative}
"

Listing 11: Meta Prompt and Template for MCQA Prompt Optimization
In the following, we show the system and user prompts returned by APO. The system prompt includes three components:
role, goal, and expected_output. In the user prompt, task_context appears before each question, and an imperative
statement follows each question. The test score was 54.
{
"role": "Industrial Asset Diagnostician",
"goal": "To accurately identify the most relevant sensors for detecting specific
failure modes in various industrial assets and determine the least relevant failure
events for abnormal sensor readings, optimizing asset performance and reducing
downtime",
"expected_output": "A concise description of the most relevant sensor for monitoring a
specific failure mode or the least relevant failure event that does not
significantly contribute to detecting a particular failure mode in the given asset,
including the option id or description",
"task_context": "The task involves analyzing the relationship between sensor readings,
failure modes, and industrial assets such as steam turbines, aero gas turbines,
electric generators, and compressors, to determine the most relevant sensors for
monitoring specific failure modes and the least relevant failure events for
abnormal sensor readings", "imperative": "Analyze the provided asset, sensor
readings, and failure modes to identify the most relevant sensor for monitoring a
specific failure mode or the least relevant failure event for an abnormal sensor
reading, considering the relationship between failure modes, sensors, and assets",
"score": 54}

Listing 12: Optimized Prompt for MCQA using Llama-3.3-70B Model

Main Algorithm The main optimization loop follows the generate-select-evaluate cycle described in Algorithm 1. It
begins by preparing the training and validation sets using the Agentics (AG[GSM8K]). Demo tasks are extracted
and transformed into OptimizationTask instances.
In each iteration, candidate prompt templates are generated via self-transduction using the meta-prompt. These templates
are applied to the validation set using the user prompt format. The responses are evaluated using the grading function
29


```

