# Page 14

## Text Content

```
Building Effective AI Agents \ Anthropic

7/9/25, 4:18 PM

In our own implementation, agents can now solve real GitHub issues
in the SWE-bench Verified benchmark based on the pull request
description alone. However, whereas automated testing helps verify
functionality, human review remains crucial for ensuring solutions
align with broader system requirements.

Appendix 2: Prompt engineering
your tools
No matter which agentic system you're building, tools will likely be
an important part of your agent. Tools enable Claude to interact with
external services and APIs by specifying their exact structure and
definition in our API. When Claude responds, it will include a tool
use block in the API response if it plans to invoke a tool. Tool
definitions and specifications should be given just as much prompt
engineering attention as your overall prompts. In this brief appendix,
we describe how to prompt engineer your tools.
There are often several ways to specify the same action. For instance,
you can specify a file edit by writing a diff, or by rewriting the entire
file. For structured output, you can return code inside markdown or
inside JSON. In software engineering, differences like these are
cosmetic and can be converted losslessly from one to the other.
However, some formats are much more difficult for an LLM to write
than others. Writing a diff requires knowing how many lines are
changing in the chunk header before the new code is written. Writing
code inside JSON (compared to markdown) requires extra escaping
of newlines and quotes.
Our suggestions for deciding on tool formats are the following:
Give the model enough tokens to "think" before it writes itself
into a corner.
Keep the format close to what the model has seen naturally
occurring in text on the internet.
Make sure there's no formatting "overhead" such as having to
https://www.anthropic.com/engineering/building-effective-agents

Page 14 of 16


```

