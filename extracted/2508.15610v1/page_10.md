# Page 10

## Text Content

```
GPUS=1
CPUS=8
MEM=64GB
LEN=8000
MODEL="Qwen/Qwen3-8B"
vllm serve ${MODEL} \
--max-model-len ${LEN} \
--tensor-parallel-size ${GPUS} \
--gpu-memory-utilization 0.9

Listing 4: vLLM parameters and computing resources

4.2

Domain-Specific Multi-Choice Question Answer

FailureSensorIQ benchmark (Constantinides et al., 2025) is recently proposed domain-specific multiple choice QA
benchmark to test LLMs’ ability to reason about failure modes and sensor relationships. The leaderboard shows that the
best performing openai/o1 model scores 60.4% 3 , unlike many saturated datasets.
Dataset We evaluated 2,667 single-correct MCQA instances spanning various industrial assets, with questions around
identifying the right sensor which can detect a given failure mode for a given asset, or identifying the right failure mode
that a given asset and sensor can detect. This requires nuanced understanding of sensor behavior, failure propagation,
and asset-specific operational logic, and performing logical deductions across the different knowledge about the asset.
An example query may be:
{
"Question": "For electric motor, if a failure event rotor windings fault
occurs, which sensor out of the choices is the most relevant sensor
regarding the occurrence of the failure event?",
"Options":["A. partial discharge", "B. resistance", "C. oil debris", "D.
current", "E. voltage"]
}

Methods Baseline results are obtained from the leaderboard that evaluates the standardized prompts with at most
three trials for invalid responses.
Our approach leverages the Agentics framework to perform schema-constrained transduction from structured
input to structured output. Each input instance is represented using a subset of fields from the FailureSensorIQ
schema—specifically, the question, asset name, option ids, options, and subject—which are sufficient to ground the
reasoning process in both linguistic and domain-specific context.
To improve inference efficiency, Agentics supports parallel batch execution via the aMap operation. This distributes
multiple structured prompts across concurrent model invocations, significantly reducing total runtime. Unlike sequential
prompting, which processes one question at a time, batch transduction enables scalable evaluation and deployment.
Experimental Configuration We evaluate four models ranging from 8B to 405B parameters: Qwen3-8B (Yang
et al., 2025), Llama-3.3-70B-Instruct (Touvron et al., 2023), Mistral-Large-Instruct-2407 (Jiang et al., 2024),
Llama-3-405B-Instruct (Touvron et al., 2023). Models are tested using both the original FailureSensorIQ baseline
pipeline and the Agentics framework. The baseline uses loosely formatted natural language prompts and retries
up to three times if the output is invalid. In contrast, Agentics uses structured prompting and schema-constrained
decoding. To measure execution time, we host Qwen3-8B on a dedicated node with an A100 80GB GPU running
VLLM (Kwon et al., 2023). Other models are accessed via cloud computing platform. We vary batch sizes to assess
scalability and throughput.
Accuracy Improvement Table 1 shows the accuracy comparison. Agentics improves the performance on all
evaluated models, with smaller models benefiting the most. Notably, Qwen3-8B achieves a major improvement of
+14.32%, getting right behind openai-o1, 60.4%. This suggests that prompting through logical transduction helps
unlock latent reasoning capabilities, even in models with limited parameter counts.
3

https://alpha-ollama.hf-mirror.com/spaces/cc4718/FailureSensorIQ

10


```

