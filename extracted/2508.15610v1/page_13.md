# Page 13

## Text Content

```
Figure 7: Improvement of test score over iterations: The x-axis represents the number of iterations, and the y-axis shows
the test score evaluated using the best prompt template found up to that iteration.

Figure 8: Average running time per iteration: The x-axis represents the batch size of the asynchronous execution, and
the y-axis shows the average running time in seconds.

4.4

Text-to-SQL

We evaluated various text-to-SQL pipelines on the challenging Bird-bench (Li et al., 2023) dataset. We observe that by
composing various components such as few-shot examples, schema linker outputs, keywords from topic models, and
sub-questions with optimized prompts, each model significantly improves execution match results, achieving up to a
10.33% increase over the baseline performance of Llama-3.3-70B.
Summary of Experiments The resultant Table 2 highlights the performance gains on average by including additional
transductions on top of the base prompt. We aggregate five runs. This includes setting a high model temperature of 0.9,
thus diversifying the generated transduction samples.
• Individual Component: We note that every individual transduction (i.e. FS / KW / SQ / SL / OP) does not in
fact improve average performance on all models.
• Schema Linker and Propmt Optimization Gains: Techniques such as schema linking and promptoptimization yield greater improvements on a few models than the other approaches.
• Adding All: Interestingly, when including all the transductions together (i.e. FS + KW + SQ + SL + OP),
performance significantly improves in a manner that is greater than the sum of its parts. This result
indicates that models can be pushed into greater performance by stipulating the right prompt programs to
captures the task.
13


```

## Images

![Image from page 13](images/page_13_img_001.ppm)

![Image from page 13](images/page_13_img_002.ppm)

![Image from page 13](images/page_13_img_003.ppm)

![Image from page 13](images/page_13_img_004.ppm)

![Image from page 13](images/page_13_img_005.ppm)

![Image from page 13](images/page_13_img_006.ppm)

![Image from page 13](images/page_13_img_007.ppm)

![Image from page 13](images/page_13_img_008.ppm)

![Image from page 13](images/page_13_img_009.ppm)

![Image from page 13](images/page_13_img_010.ppm)

![Image from page 13](images/page_13_img_011.ppm)

![Image from page 13](images/page_13_img_012.ppm)

