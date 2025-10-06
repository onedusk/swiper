# Page 8

## Text Content

```
Building Effective AI Agents \ Anthropic

7/9/25, 4:18 PM

requiring different vote thresholds to balance false positives
and negatives.

Workflow: Orchestrator-workers
In the orchestrator-workers workflow, a central LLM dynamically
breaks down tasks, delegates them to worker LLMs, and synthesizes
their results.

The orchestrator-workers workflow

When to use this workflow: This workflow is well-suited for complex
tasks where you can’t predict the subtasks needed (in coding, for
example, the number of files that need to be changed and the nature
of the change in each file likely depend on the task). Whereas it’s
topographically similar, the key difference from parallelization is its
flexibility—subtasks aren't pre-defined, but determined by the
orchestrator based on the specific input.
Example where orchestrator-workers is useful:
Coding products that make complex changes to multiple files
each time.
Search tasks that involve gathering and analyzing information
from multiple sources for possible relevant information.

https://www.anthropic.com/engineering/building-effective-agents

Page 8 of 16


```

## Images

![Image from page 8](images/page_8_img_001.ppm)

![Image from page 8](images/page_8_img_002.ppm)

