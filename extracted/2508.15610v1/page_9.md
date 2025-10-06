# Page 9

## Text Content

```
{
"Answer": "Rome",
"Justification": "The capital of Italy is a well-known fact, and Rome is the
correct answer.",
"Confidence": 1.0
}
We elaborate various design patterns and use cases for domain-specific multiple choice QA, text-to-SQL pipelines, and
automatic prompt optimization in the Appendix.
3.4.3

Execution Context

An Agentics structure can serve as either the source or the target of a transduction operation, depending on the
context. This polymorphic behavior could be enhanced by a set of additional parameters that define the execution
context, such as specialized prompt templates, external tool access during transduction, and the retrieval augmented
memory.
These additional parameters collectively guide the semantic negotiation required for logical transduction. By wrapping
Pydantic models within the Agentics class, the framework creates a highly customizable and modular execution
environment. Prompt templates, in particular, offer fine-grained control over transformation logic, allowing developers
to specify how data should be interpreted and structured. This design abstracts away the need for domain-specific prompt
engineering. Instead, developers define type systems and the agents operate independently of the content, allowing the
transduction process to dynamically infer the necessary transformations. As a result, the ≪ operator remains simple and
expressive, enabling transduction chains to be modeled as directed graphs with uniform edge semantics. This makes the
underlying algebra intuitive, and easy to implement.

4

Experiments

In this section, we demonstrate that the proposed framework achieves state-of-the-art performance on domain-specific
multiple-choice QA, enables flexible composition of text-to-SQL pipelines with improved performance on the BIRDbench evaluation, and supports efficient automatic prompt optimization in logical transduction algebra, enhancing both
downstream task performance and runtime.
In the experiments, we benchmark open-weight instruct tuned models ranging from larger or smaller parameter version
of Llama-3, Llama-4, Qwen-3, and Mistral. For the experiment that measures the running time, we host LLMs in local
vLLM (Kwon et al., 2023) server with four A-100-80GB GPUs for Llama-3-3-70B model, and one A-100-80GB GPU
for other 8B parameter models.
4.1

Computing Infrastructure
• In the MCQA experiment, we used instruction-tuned models such as Qwen3-8B, Llama-3.3-70B, MistralLarge, and Llama-3-405B. We locally hosted the Qwen3-8B model to measure running time, while the other
three models were used in a cloud computing environment.
• In automatic prompt optimization experiment, we locally hosted Qwen3-8B and Llama-3.3-70B models.
• In the Text-to-SQL experiment, we used Llama-3.3-70B, Mistral-Large, and Llama-4-Maverick-17B models.
These models were also run in a cloud computing environment.

The following shows the parameters for hosing Qwen3-8B and Llama-3.3-70B models with vLLM (Kwon et al., 2023).
GPUS=4
CPUS=16
MEM=200GB
MODEL="meta-llama/Llama-3.3-70B-Instruct"
LEN=16000
vllm serve ${MODEL} \
--max-model-len ${LEN} \
--tensor-parallel-size ${GPUS} \
--gpu-memory-utilization 0.9

9


```

