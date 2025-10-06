# Page 31

## Text Content

```
translation task, questions are posed to a database and first translated into SQL queries before being executed and
answers retrieved and answers retrieved. In practice, this task involves multiple stages of reasoning, whereby one has to
interact with the schema of the structured data, as well as understand and decompose the question into its constituent
parts.
Data Models Agentics supports this workflow by chaining multiple transduction steps and integrating them with
traditional Python logic. First, let’s consider setting up the Pydantic types with the required task information.
class Text2SQLTask(BaseModel):
question: str = Field(description="The input natural language question.")
ddl: str = Field(description="The database schema in DDL format (e.g., CREATE
TABLE statements).")
sql_query: Optional[str] = Field(description="The SQL query to be generated from
the question.")
execution_result: Optional[List[Dict[str, str]]] = Field(description="The
resulting table from executing the SQL query.")

Listing 14: Data model for the simplified Text2SQL task
From the above Pydantic types, the simplified task is to map the question and DDL to the sql_query. However, we can
break down this complex operation into declarative data modeling steps to improve task performance, including:
• enriching the database so that there are additional fields like description of the schema and business descriptions
• decompose the user question into constituent parts input that can be also a part of optimization!
• optimize the prompt template using the final input fields.
class Text2SQLTask(BaseModel):
question: str = Field(description="The input natural language question.")
ddl: str = Field(description="The database schema in DDL format (e.g., CREATE
TABLE statements).")
enrichment: Optional[DB] = Field(description="Additional database enrichments
that applies to all problem instances")
sql_query: Optional[str] = Field(description="The SQL query to be generated from
the question.")
execution_result: Optional[List[Dict[str, str]]] = Field(description="The
resulting table from executing the SQL query.")
class DB(BaseModel):
description: Optional[str] = Field(description="A Description of the business
purpose of the db, what use cases it is good for how what type of information
it contain")
keywords: Optional[list[str]] = Field(description="A list of keywords describing
the content of the database. Produce Keywords that are: Domain-Relevant:
Reflects the thematic area (e.g., education, healthcare, finance). PurposeOriented: Indicates the type of insights the database supports (e.g.,
performance tracking, demographic analysis). Unambiguous: Avoids generic or
overly broad terms. Interoperable: Aligns with standard taxonomies when
possible (e.g., DataCite or UNSDG classification). Examples of Strong Keywords:
student_outcomes, climate_metrics, financial_forecasting,
public_health_indicators, supply_chain_kpis")
few_shots: Optional[list[QuestionSQLPair]] = Field(description="A selection of the
generated question-sql pair to be used as examples of how to generate a sql
from a question.")
subquestions: Optional[list[str]] = Field(description="a list of subquestions
inside question")
schema_link: Optional[str] = Field(description="the output of the schema linker in
the form of DDL script showing relevant table, column, and values for the
question")
class QuestionSQLPair(BaseModel):
question: Optional[str]
sql_query: Optional[str]

31


```

