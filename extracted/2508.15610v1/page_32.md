# Page 32

## Text Content

```
Listing 15: Data model for the compositional Text2SQL task
Meta-Prompts and Prompt Templates We define the initial prompt that performs the main text-to-SQL task. The
following shows the prompt template used as input to the experiment, which may be optionally modified by the
automated prompt optimization flag.
prompt_template = "
Translate the input natural language **question** to a valid SQLite query that can be
executed on the following database in **dbs**.
Do your best to apply the following rules when generating SQL.
- Cleary understand the **question** and given database description in **dbs**.
- Database description is given in the **dbs** with the following fields.
- The database description under **dbs** contains DDL scripts or natural language
description of the database, tables, columns, and values.
- Only SELECT statements are allowed, do not produce any DDL or DML.
- When writing ‘SELECT <column>‘, only include the columns specifically mentioned in
the question.
- Use **evidence** to find correct column names or the values of the columns or other
expressions.
- If you see ’None’ or ‘None‘ in the [Value examples] for a column, prioritize using ‘
JOIN <table>‘ or ‘WHERE <column> IS NOT NULL‘ to handle potential missing data
effectively.
- Use ‘WHERE <column> IS NOT NULL‘ in ‘WHERE‘ if you are sorting with ‘<column>‘.
- Use alias in ‘SELECT‘ is consistently in the expressions.
- Use ‘WHERE <alias> IS NOT NULL‘ in ‘WHERE‘ if you are sorting with ‘<alias>‘.
Input is provided under SOURCE.
"{input_spec_str}"
Generate Output as JSON decodable format
"{{"
generated_sql_query: a valid SQLite query that translates nautral language question
"}}"\n
"

Listing 16: The Prompt Template and Meta-Prompt for Automatic Prompt Optimization
Main Algorithm Next, we show how to define the compositional text-to-SQL pipeline in Agentics. We see that
the entire pipeline can be implemented through compositions of logical transductions defined over the data models.
text2sql("enrichment.keywords") << text2sql("enrichment.description") << text2sql("
ddl_schema")
text2sql("few_shots.question", "few_shots.sql_query") << text2sql("ddl_schema", "
enrichment.description") # synthetic pair
text2sql("few_shots") = text2sql("few_shots") + k_shot # additional augmentations
text2sql("few_shots") = text2sql.filter(valid_sql, "few_shots") # executes sql
text2sql("enrichment.subquestions") << text2sql("question")
text2sql("enrichment.linked_schema") << text2sql("enrichment.ddl_schema", "enrichment.
subquestions", "question")
input_fields = ["question", "enrichment"]
text2sql.instructions = optimize(prompt_template, input_fields)
text2sql("sql_query") << text2sql(*input_fields)
text2sql("execution_result") = text2sql.amap(execute_sql_query)

Listing 17: Pseudo-code for the Compositional Text-to-SQL Workflow

32


```

