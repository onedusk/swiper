# Page 15

## Text Content

```
Building Effective AI Agents \ Anthropic

7/9/25, 4:18 PM

keep an accurate count of thousands of lines of code, or stringescaping any code it writes.
One rule of thumb is to think about how much effort goes into
human-computer interfaces (HCI), and plan to invest just as much
effort in creating good agent-computer interfaces (ACI). Here are
some thoughts on how to do so:
Put yourself in the model's shoes. Is it obvious how to use this
tool, based on the description and parameters, or would you
need to think carefully about it? If so, then it’s probably also true
for the model. A good tool definition often includes example
usage, edge cases, input format requirements, and clear
boundaries from other tools.
How can you change parameter names or descriptions to make
things more obvious? Think of this as writing a great docstring
for a junior developer on your team. This is especially important
when using many similar tools.
Test how the model uses your tools: Run many example inputs in
our workbench to see what mistakes the model makes, and
iterate.
Poka-yoke your tools. Change the arguments so that it is harder
to make mistakes.
While building our agent for SWE-bench, we actually spent more
time optimizing our tools than the overall prompt. For example, we
found that the model would make mistakes with tools using relative
filepaths after the agent had moved out of the root directory. To fix
this, we changed the tool to always require absolute filepaths—and
we found that the model used this method flawlessly.

Product

Research

Commitments

Learn

Help and security

Claude overview

Research
overview

Transparency

Anthropic
Academy

Status

Claude Code
Max plan

Economic
Index

Team plan
https://www.anthropic.com/engineering/building-effective-agents

Responsible
scaling policy
Security and

Customer stories

Availability
Support center

Engineering at
Page 15 of 16


```

