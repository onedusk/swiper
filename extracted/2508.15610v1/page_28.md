# Page 28

## Text Content

```
each question. A score of 89 was evaluated on the validation set, using 100 problems sampled from the training set,
which is higher than the test score of 85.
{
"role": "Elite Mathematical Problem Solver",
"goal": "To swiftly and accurately determine the precise numerical solution by
meticulously dissecting complex problem statements, identifying pivotal information
, and applying a comprehensive array of advanced mathematical concepts, formulas,
and logical reasoning to achieve an optimal solution, ensuring efficiency,
precision, and reliability in all calculations, while providing accurate and
relevant answers",
"expected_output": "A single, exact numeric value that directly addresses and solves
the given mathematical problem, ensuring all calculations are correctly executed,
based on the information provided, and presented in a clear, concise manner, with
strict adherence to mathematical rules and consideration of all given data",
"imperative": "Thoroughly analyze the problem statement, extract key information,
apply pertinent mathematical principles and logical reasoning to derive the
accurate numerical answer, and provide the final answer in the required numeric
format, while validating the solution through re-evaluation of calculations and
verification of the accuracy of the obtained result, and ensuring precision,
accuracy, and reliability in all steps of the calculation process.",
"score": 89
}

Listing 9: Optimized Prompt for GSM8K using Llama-3.3-70B Model
The following result is obtained from the Qwen3-8B model. A score of 98 was evaluated on the validation set, using
100 problems sampled from the training set. This score is higher than the test score of 92.
{
"role": "Math Problem Solver",
"goal": "Solve complex problems by decomposing them into sequential steps, applying
arithmetic operations (percentages, fractions, averages), ensuring unit consistency
, and presenting the final answer in a boxed format.",
"expected_output": "A single numeric value boxed (e.g., \\boxed{64}) representing the
solution, derived through precise step-by-step calculations with attention to
percentages, fractions, and averages.",
"imperative": "Analyze the problem statement, execute calculations step-by-step with
focus on percentages, fractions, and unit conversions, then box the final numeric
result.",
"score": 98
}

Listing 10: Optimized Prompt for GSM8K using Qwen3-8B Model
The following shows the meta prompt for optimizing the prompt template for MCQA dataset.
OPT_META_INSTRUCTION = "Your proposed prompt template will be used in the following
way.
* You are "role" -- this role must be suitable for solving the demo task.
* Your personal goal is: "goal" -- the goal achieves the outputs given inputs.
* This is the expected criteria for your final answer "expected_output" -- this
constrains the output format.
* Extract"task_context" from demo tasks to explain the problem context -- this comes
before the input of the task.
* You can add a short imperative instruction "imperative" -- this comes after the
input of the task.

[[Several demo tasks of input and outputs will be provided when you solve problem.]]
[[The previous optimized prompt templates with scores appear from the worst to the
best.]]
{optimization_history}

28


```

