# Page 11

## Text Content

```
Model
# Params Baseline
Agentics
60.18 (+14.32)
Qwen3-8B
8B
45.86
50.73 (+9.04)
70B
41.69
Llama-3.3-70B
58.41 (+8.32)
Mistral-Large
123B
50.09
52.90 (+1.64)
405B
51.26
Llama-3-405B
Table 1: Accuracy (%) of on 2,667 FailureSensorIQ instances.

Running Time Figure 6a illustrates the average time (sec) per question for Qwen3-8B across varying batch sizes.
As shown, parallel batch execution yields substantial speedups, from 8 seconds per question at batch size 1 to less than
1 second per question at batch sizes greater than 16. This improvement is nearly linear as the batch size increases from 1
to 4, after which it begins to saturate.
Perturbation Study We follow FailureSensorIQ’s perturbation study using the Agentics framework to
study if there are any robustness benefits that comes with the structured workflows that our framework offers. We
experiment with the following knowledge invariant perturbations:
• Option letter renaming; changing the option letters from of A., B., C., to other letters like P., Q., R. We’ll call
this “Simple” Perturbation.
• Option letter renaming and question rephrasing done by an LLM. We’ll call this “Complex” Perturbation.
We use the already prepared perturbed datasets from the original paper.

(a) Average time per question for Qwen3-8B across varying
batch sizes.

(b) Performance remains high even after the perturbations with
minimal drop.

Figure 6: Domain Specific MCQA Running Time and Perturbation Results.
Summary of Experiments
• Accuracy Gains: All models show a major performance improvement under Agentics, with smaller models
benefiting most. This suggests that structured prompting helps unlock latent reasoning capabilities, even in
models with limited parameter counts.
• Efficiency Gains: Agentics supports scalable batch execution, significantly reducing inference time and
making large-scale evaluation practical.
• Perturbation Robustness: Agentics seems to be robust against knowledge invariant perturbations. In the
original FailureSensorIQ experiments, all the models experienced significant drop in performance. We
believe this is due to the decoupling that Agentics provides, where option ids and options are provided as
separate fields and this makes the models easier to attent to the right information.
4.3

Prompt Optimization

Automatic prompt optimization (APO) is essential, as LLM performance is highly sensitive to prompt structure,
tone, and formatting. The prompt function in Definition 4 plays a central role in logical transduction, which can be
conceptualized as a negotiation of meaning between agents. Among various APO approaches (Ramnath et al., 2025; Li
11


```

## Images

![Image from page 11](images/page_11_img_001.ppm)

![Image from page 11](images/page_11_img_002.ppm)

![Image from page 11](images/page_11_img_003.ppm)

![Image from page 11](images/page_11_img_004.ppm)

