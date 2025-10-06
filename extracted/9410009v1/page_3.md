# Page 3

## Text Content

```
one can have $strong refer to the full entry describing ‘strong’ in say its ordinary use, and have
the values that are particular to the collocational
strong overwrite the values provided in the ordinary entry, as in Mel’čuk’s proposal.
Collocations, Rules and Principles So far,
we have not specified in what way one gets from
the lexical entries for the base and the collocate to
the representation of the collocational expression.
In HPSG, the descriptions of complex expressions are constrained by principles. We will assume that collocations are subject to the same constraints. The ordinary rules of combination (combining adjectives and nouns, for instance) thus account for most of the properties of the collocational combination. However, we are still left with
the typical ‘collocational restriction’ which needs
to be accounted for.
We have therefore added a principle which says
that constructions that are analysed as collocations (indicated by the type collocation) are either head-adjunct structure or head-complement
structures with specific restrictions holding between the head and the adjunct or the head and
the complement respectively. Let’s consider the
former case4 , illustrated by the heavy smoker example. The adjunct daughter will contain the adjective collocate. In such collocational constructions the collocate adjuncts have to be ‘licensed’
by the noun or the head daughter. This is implemented by requiring that the collocates field
(colls) of the head daughter contains a reference
to a lexical entry that is compatible with the adjunct daughter. In the literal reading of an expression such as heavy smoker , the phrase will not be
analysed as a collocation and the principle does
not apply.



⇒

 COLLOCATION
HEAD DTR

HEAD DTR
COMP DTRS

< ...[COLLS{... 1 ...}]... >

ADJ DTRS



3



COLLS {... 1 ...}
< ...W
1 COLLOCATE ... >
1 COLLOCATE




Issues in Translation

The project has tried to investigate the use of lexical functions as an interlingual device, i.e., one
which is shared by the semantic representations of
collocations in the language pairs5 .
4 To illustrate the case of head-complement

structures
one could take some support verb construction (also called
light verb construction).
5 For another application of LFs in a multilingual NLP
context see ([Heid and Raab, 1989]). For other treatments
of
collocations in language generation see ([Nirenburg et al., 1988])
and ([Smadja and McKeown, 1990]).

The typing of a collocation with such a function opens up the way to a treatment of collocations inside a given language module and hence to
a substantial reduction in the number of collocations explicitly handled in the multilingual transfer
dictionary. The existence of a collocation function
is established during analysis. This information
is used to generate the correct translation in the
target language. To illustrate, the English analysis module might analyse (1) as (2). The transfer
module maps (2) onto (3) which is then synthesised by the French module to (4).
(1) heavy smoker → (2) Magn(smoker)
→ (3) Magn(fumeur) → (4) grand
fumeur
The example points out that the translation
strategy is a mixture of transfer and interlingua.
The bases are transferred but the representation
of the collocate is shared between the source and
the target representation. This treatment of collocations rests, among others, on the assumptions
that there are only a limited number of lexical
functions, that lexical functions can be assigned
consistently, that all (or a significant number of)
collocations realise a lexical function, that lexical functions are not restricted to particular languages, etc. In the following paragraph we present
an outline of the translation process. Next, we discuss some of the problems which follow from our
approach and we propose some ways to solve them.

3.1

Lexical Functions as Interlingua

It was assumed that the starting point for transfer is the semantic representation of the phrase.
Using a semantic representation as input to transfer implies that we relate semantic values of words
and phrases. For our purposes this is very satisfying since we will now be using the semantics
of collocates instead of their orthography, in other
words: we use lexical functions and abstract away
from the particular realisation of a collocate in a
particular language.
We now state the relation between the semantic representations of the source language and target language. The semantic relation between the
phrase heavy smoker and its French counterpart
can be made explicit in the following bilingual sign:

 

VAR
1
REST {smoker( 1 ),Magn( 1 )}

EN—SEM IND






FR—SEM IND



VAR
1
REST {fumeur( 1 ),Magn( 1 )}



 


Typically, the lexicon will contain a bilingual
sign for each possible value of reln. Thus, for
translating heavy smoker into grand fumeur we will
need the obvious entry for smoker-fumeur plus the
entry below:


```

## Images

![Image from page 3](images/page_3_img_001.pbm)

![Image from page 3](images/page_3_img_002.pbm)

![Image from page 3](images/page_3_img_003.pbm)

![Image from page 3](images/page_3_img_004.pbm)

![Image from page 3](images/page_3_img_005.pbm)

![Image from page 3](images/page_3_img_006.pbm)

![Image from page 3](images/page_3_img_007.pbm)

![Image from page 3](images/page_3_img_008.pbm)

![Image from page 3](images/page_3_img_009.pbm)

![Image from page 3](images/page_3_img_010.pbm)

![Image from page 3](images/page_3_img_011.pbm)

![Image from page 3](images/page_3_img_012.pbm)

![Image from page 3](images/page_3_img_013.pbm)

![Image from page 3](images/page_3_img_014.pbm)

![Image from page 3](images/page_3_img_015.pbm)

![Image from page 3](images/page_3_img_016.pbm)

![Image from page 3](images/page_3_img_017.pbm)

![Image from page 3](images/page_3_img_018.pbm)

![Image from page 3](images/page_3_img_019.pbm)

![Image from page 3](images/page_3_img_020.pbm)

![Image from page 3](images/page_3_img_021.pbm)

![Image from page 3](images/page_3_img_022.pbm)

![Image from page 3](images/page_3_img_023.pbm)

![Image from page 3](images/page_3_img_024.pbm)

![Image from page 3](images/page_3_img_025.pbm)

![Image from page 3](images/page_3_img_026.pbm)

![Image from page 3](images/page_3_img_027.pbm)

![Image from page 3](images/page_3_img_028.pbm)

![Image from page 3](images/page_3_img_029.pbm)

![Image from page 3](images/page_3_img_030.pbm)

![Image from page 3](images/page_3_img_031.pbm)

![Image from page 3](images/page_3_img_032.pbm)

![Image from page 3](images/page_3_img_033.pbm)

![Image from page 3](images/page_3_img_034.pbm)

![Image from page 3](images/page_3_img_035.pbm)

![Image from page 3](images/page_3_img_036.pbm)

![Image from page 3](images/page_3_img_037.pbm)

![Image from page 3](images/page_3_img_038.pbm)

![Image from page 3](images/page_3_img_039.pbm)

![Image from page 3](images/page_3_img_040.pbm)

![Image from page 3](images/page_3_img_041.pbm)

![Image from page 3](images/page_3_img_042.pbm)

![Image from page 3](images/page_3_img_043.pbm)

![Image from page 3](images/page_3_img_044.pbm)

![Image from page 3](images/page_3_img_045.pbm)

![Image from page 3](images/page_3_img_046.pbm)

![Image from page 3](images/page_3_img_047.pbm)

