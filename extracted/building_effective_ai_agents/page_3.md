# Page 3

## Text Content

```
Building Effective AI Agents \ Anthropic

7/9/25, 4:18 PM

When and how to use frameworks
There are many frameworks that make agentic systems easier to
implement, including:
LangGraph from LangChain;
Amazon Bedrock's AI Agent framework;
Rivet, a drag and drop GUI LLM workflow builder; and
Vellum, another GUI tool for building and testing complex
workflows.
These frameworks make it easy to get started by simplifying
standard low-level tasks like calling LLMs, defining and parsing
tools, and chaining calls together. However, they often create extra
layers of abstraction that can obscure the underlying prompts and
responses, making them harder to debug. They can also make it
tempting to add complexity when a simpler setup would suffice.
We suggest that developers start by using LLM APIs directly: many
patterns can be implemented in a few lines of code. If you do use a
framework, ensure you understand the underlying code. Incorrect
assumptions about what's under the hood are a common source of
customer error.
See our cookbook for some sample implementations.

Building blocks, workflows, and
agents
In this section, we’ll explore the common patterns for agentic
systems we’ve seen in production. We'll start with our foundational
building block—the augmented LLM—and progressively increase
complexity, from simple compositional workflows to autonomous
agents.

Building block: The augmented LLM
https://www.anthropic.com/engineering/building-effective-agents

Page 3 of 16


```

