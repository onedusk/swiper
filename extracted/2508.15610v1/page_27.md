# Page 27

## Text Content

```
# goal statement describing what the prompt aims to achieve
goal: str = Field(description="New goal instruction")
# criteria for what constitutes a good or acceptable output
expected_output: str = Field(description="New expected_output instruction")
# imperative phrase (e.g., ’Let’s think step by step’)
imperative: str = Field(description="New imperative")
# evaluation score assigned to the prompt after testing on validation data
score: int = Field(description="evaluation score")
class GSM8K(BaseModel):
question: str = Field(description="a grade school math question.")
answer: str = Field(description="the ground-truth answer")
response_think: str = Field(description="the step by step reasoning")
response_answer: str = Field(description="the final answer")
# boolean flag indicating whether the response answer is correct
correct: bool

Listing 7: Data Models for GSM8K Prompt Optimization
Meta Prompt and Optimized Prompts The meta-prompt (OPT_META_INSTRUCTION) guides the generation of new prompt templates by describing the structure and expectations for the optimizer. It includes historical context from previous iterations and instructs the model to avoid redundancy. The user prompt template
(USER_PROMPT_TEMPLATE) is instantiated with the optimized parameters and used to evaluate candidate prompts
on the validation set.
The following shows the meta-prompt used for optimizing the prompt template for the GSM8K dataset.
OPT_META_INSTRUCTION = "Your proposed prompt template will be used in the following
way.
* You are "role" -- this role must be suitable for solving the demo task.
* Your personal goal is: "goal" -- the goal achieves the outputs given inputs.
* This is the expected criteria for your final answer "expected_output" -- this
constrains the output format.
* You can add a short imperative instruction "imperative" -- this comes after the
input of the task.
[[Several demo tasks of input and outputs will be provided when you solve problem.]]
[[The previous optimized prompt templates with scores appear from the worst to the
best.]]
{optimization_history}
* Given the previous optimization results, don’t generate duplicate or similar prompt
templates.
* Generate prompt template that achieves the best score, and succint and concise
instructions.
"

USER_PROMPT_TEMPLATE = "
You are {role}.
Your personal goal is: {goal}.
This is the expected criteria for your final answer: {expected_output}.
solve the following task.
{question}
{imperative}
"

Listing 8: Meta Prompt and Template for GSM8K Prompt Optimization
In the following, we show the prompts returned by APO, including both system and user prompts. The system prompt
consists of three components: role, goal, and expected_output. In the user prompt, an imperative statement appears after
27


```

