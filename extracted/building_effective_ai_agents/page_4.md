# Page 4

## Text Content

```
Building Effective AI Agents \ Anthropic

7/9/25, 4:18 PM

The basic building block of agentic systems is an LLM enhanced with
augmentations such as retrieval, tools, and memory. Our current
models can actively use these capabilities—generating their own
search queries, selecting appropriate tools, and determining what
information to retain.

The augmented LLM

We recommend focusing on two key aspects of the implementation:
tailoring these capabilities to your specific use case and ensuring
they provide an easy, well-documented interface for your LLM. While
there are many ways to implement these augmentations, one
approach is through our recently released Model Context Protocol,
which allows developers to integrate with a growing ecosystem of
third-party tools with a simple client implementation.
For the remainder of this post, we'll assume each LLM call has access
to these augmented capabilities.

Workflow: Prompt chaining
Prompt chaining decomposes a task into a sequence of steps, where
each LLM call processes the output of the previous one. You can add
programmatic checks (see "gate” in the diagram below) on any
intermediate steps to ensure that the process is still on track.

https://www.anthropic.com/engineering/building-effective-agents

Page 4 of 16


```

## Images

![Image from page 4](images/page_4_img_001.ppm)

![Image from page 4](images/page_4_img_002.ppm)

