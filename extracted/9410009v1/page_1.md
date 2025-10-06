# Page 1

## Text Content

```
LEXICAL FUNCTIONS AND MACHINE TRANSLATION
Dirk Heylen, Kerry G. Maxwell and Marc Verhagen

arXiv:cmp-lg/9410009v1 20 Oct 1994

OTS, Trans 10, 3512 JK Utrecht, Netherlands
CLMT Group, Essex University, Colchester, Essex CO4 3SQ, England
email: heylen@let.ruu.nl, kerry@essex.ac.uk, verhm@essex.ac.uk
cmp-lg/9410009
This paper discusses the lexicographical concept of
lexical functions ([Mel’čuk and Žolkovsky, 1984])
and their potential exploitation in the development of a machine translation lexicon designed
to handle collocations.
We show how lexical functions can be thought to reflect crosslinguistic meaning concepts for collocational structures and their translational equivalents, and
therefore suggest themselves as some kind of
language-independent semantic primitives from
which translation strategies can be developed.1

1

Description of the Problem

Collocations present specific problems in translation, both in human and automatic contexts. If we
take the construction heavy smoker in English and
attempt to translate it into French and German,
we find that a literal translation of heavy yields the
wrong result, since the concept expressed by the
adjective (something like ‘excessive’) is translated by grand (large) in French and stark (strong)
in German. We observe then that in some sense
the adjectives stark, grand and heavy are equivalent in the collocational context, but that this
is of course not typically the case in other contexts, cf grande boite, starke Schachtel and heavy
box, where the adjectives could hardly be viewed as
equivalent. It seems then that adjectives which are
not literal translations of one another may share
meaning properties specifically in the collocational
context.
How then can we specify this special equivalence
in the machine translation dictionary? The answer
seems to lie in addressing the concept which underlies the union of adjective and noun in these
three cases, i.e., intensification, and hence establish a single meaning representation for the adjectives which can be viewed as an interlingual pivot
for translation.
Collocations
have been studied by computational linguists in
1 The research reported in this paper was undertaken as
the project “Collocations and the Lexicalisation of Semantic Operations” (ET-10/75). Financial contributions were
by the Commission of the European Community, Association Suissetra (Geneva) and Oxford University Press.

different contexts. For instance, there is a substantial body of papers on the extraction of “frequently co-occurring words” from corpora using
statistical methods (e.g., ([Choueka et al., 1983]),
([Church and Hanks, 1989]), ([Smadja, 1993]) to
list only a few). These authors focus on techniques for providing material that can be used
in other processing tasks such as word sense disambiguation, information retrieval, natural language generation and so on.
Also, the use
of collocations in different applications has been
discussed by various authors (([McRoy, 1992]),
([Pustejovsky et al., 1992]),
([Smadja and McKeown, 1990]) etc.). However,
collocations are not only considered useful, but
also a problem both in certain applications
(e.g. generation, ([Nirenburg et al., 1988]), machine translation, ([Heid and Raab, 1989])) and
from
a
more
theoretical
point of view (e.g. ([Abeillé and Schabes, 1989]),
([Krenn and Erbach, To appear])).
We have been concerned with investigating the lexical functions (LFs) of Mel’čuk
([Mel’čuk and Žolkovsky, 1984]) as a candidate interlingual device for the translation of adjectival
and verbal collocates. Our work is related to research by ([Heid and Raab, 1989]). In some respects it is an extension of some of their suggestions. Our work differs from theirs in scope and
also in the exploration of various other directions.

2

Representation

The use we make of lexical functions as interlingual representations, does not respect their original Mel’čukian interpretation. Furthermore, we
have transferred them from their context in the
Meaning-Text Theory to a different theoretical setting. We have embedded the concept in an HPSGlike grammar theory.2 In this section we review
this operation. First we consider the features of
Mel’čuk’s treatment that we have wanted to preserve. Next we show how they have been imported
into the HPSG framework.
2 Head

Driven Phrase Structure grammar, see
([Pollard and Sag, 1987]), ([Pollard and Sag, to appear]).
For another treatment of collocations in HPSG, see
([Krenn and Erbach, To appear]).


```

## Images

![Image from page 1](images/page_1_img_001.pbm)

![Image from page 1](images/page_1_img_002.pbm)

