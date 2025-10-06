# Page 15

## Text Content

```
away, can introduce fragility and reduce transparency. For complex applications such as text-to-SQL or
structured data transformation, developers may need to construct entire programs from low-level building
blocks, as native support for external execution logic or dynamic schema evolution is often lacking.
• The data models in these systems are typically rigid, with input and output fields predefined in static schemas.
This restricts dynamic type creation and makes it difficult to express transformations between fields or adapt to
evolving data structures. The lack of flexibility can be compared to a neural network framework that disallows
architectural changes during training—limiting experimentation and adaptability.
• While some frameworks enrich user-provided task descriptions with verbose metadata, they often do not
make explicit how these transformations are applied or interpreted. This opacity can hinder debugging,
reproducibility, and the ability to reason about the system’s behavior.
• Although modern LLM inference APIs support efficient features such as batch processing and multi-sample
generation, many frameworks may not natively expose these capabilities.
In contrast, Agentics adopts a data-centric and type-driven approach. Logical transductions and asynchronous
map/reduce operations are composed declaratively, allowing developers to focus on semantic alignment and structured
transformation. LLM calls are treated as native asynchronous functions, making it natural to embed them within
general-purpose Python programs. This design simplifies development while preserving flexibility, type safety, and
composability.

6

Conclusion

We present a principled framework for agentic AI, grounded in a novel logical transduction algebra and a scalable
programming model. The proposed framework redefines how agents interact with data through a declarative, type-driven
approach. Its successful application to tasks such as multiple-choice QA, text-to-SQL parsing, and prompt optimization
demonstrates state-of-the-art performance and scalability, offering a new perspective on building intelligent systems.

References
Acharya, D. B., Kuppan, K., and Divya, B. (2025). Agentic AI: Autonomous intelligence for complex goals–a
comprehensive survey. IEEE Access.
Asai, A., Wu, Z., Wang, Y., Sil, A., and Hajishirzi, H. (2024). Self-RAG: Learning to retrieve, generate, and critique
through self-reflection. In The Twelfth International Conference on Learning Representations.
AtomicAI (2024). Atomic Agents and Pydantic AI: Modular and type-safe agentic frameworks.
Cobbe, K., Kosaraju, V., Bavarian, M., Chen, M., Jun, H., Kaiser, L., Plappert, M., Tworek, J., Hilton, J., Nakano, R.,
Hesse, C., and Schulman, J. (2021). Training verifiers to solve math word problems. arXiv preprint arXiv:2110.14168.
Codd, E. F. (1970). A relational model of data for large shared data banks. Communications of the ACM, 13(6):377–387.
Constantinides, C., Patel, D., Lin, S., Guerrero, C., Patil, S. D., and Kalagnanam, J. (2025). FailuresensorIQ: A
multi-choice QA dataset for understanding sensor relationships and failure modes. arXiv preprint arXiv:2506.03278.
CrewAI Inc. (2025). CrewAI. 2025-07-15.
Dean, J. and Ghemawat, S. (2008). Mapreduce: simplified data processing on large clusters. Communications of the
ACM, 51(1):107–113.
Dong, Y., Mu, R., Jin, G., Qi, Y., Hu, J., Zhao, X., Meng, J., Ruan, W., and Huang, X. (2024). Building guardrails for
large language models. In Bennett, K. and Srikumar, V., editors, Proceedings of the 41st International Conference on
Machine Learning (ICML), volume 235 of Proceedings of Machine Learning Research, pages 235–249. PMLR.
Gammerman, A., Vovk, V., and Vapnik, V. (1998). Learning by transduction. In Proceedings of the Fourteenth
Conference on Uncertainty in Artificial Intelligence, pages 148–155, San Francisco, CA, USA. Morgan Kaufmann
Publishers Inc.
Han, S., Zhang, Q., Yao, Y., Jin, W., Xu, Z., and He, C. (2024). LLM multi-agent systems: Challenges and open
problems. CoRR, abs/2402.03578.
Heck, P. (2024). What about the data? a mapping study on data engineering for AI systems. In Proceedings of the
IEEE/ACM 3rd International Conference on AI Engineering-Software Engineering for AI, pages 43–52.
Hendrycks, D., Burns, C., Basart, S., Zou, A., Mazeika, M., Song, D., and Steinhardt, J. (2021). Measuring massive
multitask language understanding. In International Conference on Learning Representations.
15


```

