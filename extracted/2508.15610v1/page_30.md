# Page 30

## Text Content

```
defined in GSM8K.grade, and scores are assigned. The best-performing prompts are retained using a filtering function
(keep_best_k). This loop continues until convergence or a maximum number of iterations is reached. The use of
transduction and asynchronous execution enables parallel evaluation of prompt candidates, improving scalability and
runtime efficiency.
# load GSM8K training data into \texttt{Agentics} abstraction
trainset = AG.from_json("gsm8k_train.jsonl")
# truncate to a subset for training
trainset = trainset.truncate_states(num_trains)
# create demo tasks from training examples
demosets = create_optimization_demos(trainset, num_demos)
# convert demo tasks into OptimizationTask instances
optimization_tasks = OptimizationTask.create_optimization_tasks(demosets)
# prepare validation set from remaining examples
validationset = trainset.truncate_states(num_trains, num_trains + num_devs)
# initialize optimizer AG[OptimizationTask] with demo tasks
optimizer = AG.from_states(optimization_tasks, atype=OptimizationTask)
# set default parameters and prompt configuration
set_default_params(optimizer)
optimizer.prompt_template = "{{"demo tasks":{demos}}}"
optimizer.prompt_params = {"role": "Prompt optimizer.", "goal": "Propose diverse
prompt templates that achieves high performance for the demo task given as input.",
"backstory": "Understand the problem domain given the demo task example and
propose what answer should be generated.", "expected_output": "the outputs are role
, goal, and the expected output description, and imperative sentence for solving
provided tasks."}
# initialize list to store optimized prompt tasks
optimized_tasks = []
for iter_ind in range(max_iter):
# generate candidate prompt templates using meta-prompt and transduction
# transduction from demos to prompt parameters
optimizer.instructions = OPT_META_INSTRUCTION.format(
optimization_history = get_history_string(optimized_tasks))
optimizer = asyncio.run(optimizer.self_transduction(
["demos"],
["role", "goal", "expected_output", "imperative"]))
# apply candidate prompts to validation set using user prompt format
opt_eval = optimizer * validationset
opt_eval.prompt_template = USER_PROMPT_TEMPLATE
opt_eval = asyncio.run(opt_eval.self_transduction(
["role", "goal", "expected_output", "imperative", "question"],
["response_think", "response_answer"]))
# evaluate responses using GSM8K grading function
executed_tasks = opt_eval / validationset
for ind, exectask in enumerate(executed_tasks):
exectask = asyncio.run(exectask.amap(GSM8K.grade))
setattr(optimizer[ind], "score", summary["score"])
# retain top-performing prompts for next iteration
optimized_tasks.extend(optimizer.states)
optimized_tasks = keep_best_k(optimized_tasks)

Listing 13: Pseudo Code for GSM8K Prompt Optimization

B.3

Semantic Parsing Text-to-SQL

Text-to-SQL is an essential task for broadening the accessibility of structured data interaction, allowing users to query
databases without needing to understand the underlying decisions made by data engineers. Loosely considered a
30


```

