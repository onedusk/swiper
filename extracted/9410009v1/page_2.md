# Page 2

## Text Content

```
2.1

Collocations and LFs

In Mel’čuk’s Explanatory Combinatory Dictionary (ECD, see ([Mel’čuk et al., 1984])), expressions such as une ferme intention, une résistance
acharnée, un argument de poids, un bruit infernal
and donner une leçon, faire un pas, commetre un
crime are described in the lexical combinatorics
zone. These “expressions plus ou moins figées”
will be called ‘collocations’. They are considered
to consist of two parts — the base and the collocate. In the examples above, the nouns are the
bases and the adjectives and the verbs are the collocates. The idea that all adjective collocates and
all the verb collocates share an important meaning
component — roughly paraphrasable as intense
and do respectively — and the fact that the adjectives and verbs are not interchangeable but are
restricted with this meaning to the accompanying nouns, is coded in the dictionary using lexical
functions (in this case Magn and Oper).
Each article in the ECD describes what is called
a ‘lexeme’: a word in some specific reading. In
the lexical combinatorics zone, we find a list of
the lexical functions that are relevant to this particular lexeme. Each lexical function is followed
by one or more lexemes (the result or value of the
function applied to the head word). The idea is
that each combination of the argument with one
of the values of the function forms a collocation in
our terminology. The argument corresponds to the
base and each value is a collocate. The following
features of this representation are important to us.

• Coding the base-collocate relation in the lexicon.
• Choosing the level at which lexical functions
will be situated.
• Relating the collocate information to the free
variant entry.
We have provided straightforward solutions to
these problems. For the first problem we have
taken over the ECD architecture rather directly,
by creating a dedicated ‘collocates’ field in the entry for bases which contains all the relevant collocates. As far as the second problem is concerned,
the obvious place to put lexical functions is in the
semantic representation provided by HPSG. There
are various reasons for this. One is that LFs are
used in the deep syntax level in Mel’čuk’s model,
a level oriented towards meaning. Another reason is that this level seems most appropriate to be
used in transfer/translation and because we want
to use lexical functions in transfer, this is where
they should be. In contrast to the ECD, the meaning of the collocate is represented by the lexical
function only.
The following is an example of the entry for criticism with the encoding of strong as a collocate.3
We use sem ind as an abbreviation for the feature
path sem.cont.ind.
 PHON

criticism


VAR
1
 SEM IND



REST
{criticism(
1
)}


 


$strong 



 COLLS {
} 
VAR
1
SEM IND

• Lexical functions are used to represent an important syntactico-semantic relation between
the base and the collocate.
• The restricted combinatorial potential of the
collocate lexeme is accounted for by listing it
at each base with which it can occur.
The second of these characteristics points out
that the collocational restriction is seen as a purely
lexical, idiosyncratic one: all collocations are explicitly listed.
One other aspect of collocations which we have
to deal with is the relation between the collocate
lexeme and its freely occurring counterpart. Collocate lexemes often differ in some respects from
their literal variants while sharing other properties. Mel’čuk deals with this by including in the
ECD an entry for the free variant and putting the
collocate-specific information in the entry for the
base (with the result of the lexical functions). The
full entry of the collocate is the result of taking the
entry for the free variant and overwriting it with
the information provided at the base.

2.2

Collocations in HPSG

The three aspects of Mel’čuk’s analysis we wanted
to encode in HPSG were the following.

REST {Magn( 1 )}

Just as in the ECD the base contains a specific
zone in which the collocates are listed. In our case,
the feature ‘colls’ has a set of lexical entries as
its value.
Each collocate subentry bears the value of the
lexical function in its semantics field. In this representation the lexical function is chosen as the
real semantic value of the collocate. One should
read the feature structure as specifying that the semantics of strong (as a collocate) is the predicate
Magn( 1 ).
The collocate subentry only provides partial information. In fact, it provides only the information
that is specific to the occurrence of strong in its
combination with criticism. In this case only the
semantics is given. We further assume that the lexicon also contains a ‘super-entry’ which provides
all the information that is shared by all the different occurrences of strong. This entry is where the
variable $strong points to. Of course, other architectures that try to avoid redundant specification
of information are equally possible. For instance
if one assumes a mechanism of default unification,
3 Notice that here we use a simple version of HPSG based
on ([Pollard and Sag, 1987]) whereas the actual implementation was based on ([Pollard and Sag, to appear]).


```

## Images

![Image from page 2](images/page_2_img_001.pbm)

![Image from page 2](images/page_2_img_002.pbm)

![Image from page 2](images/page_2_img_003.pbm)

![Image from page 2](images/page_2_img_004.pbm)

![Image from page 2](images/page_2_img_005.pbm)

![Image from page 2](images/page_2_img_006.pbm)

![Image from page 2](images/page_2_img_007.pbm)

![Image from page 2](images/page_2_img_008.pbm)

![Image from page 2](images/page_2_img_009.pbm)

![Image from page 2](images/page_2_img_010.pbm)

![Image from page 2](images/page_2_img_011.pbm)

![Image from page 2](images/page_2_img_012.pbm)

![Image from page 2](images/page_2_img_013.pbm)

![Image from page 2](images/page_2_img_014.pbm)

![Image from page 2](images/page_2_img_015.pbm)

![Image from page 2](images/page_2_img_016.pbm)

![Image from page 2](images/page_2_img_017.pbm)

![Image from page 2](images/page_2_img_018.pbm)

![Image from page 2](images/page_2_img_019.pbm)

![Image from page 2](images/page_2_img_020.pbm)

![Image from page 2](images/page_2_img_021.pbm)

![Image from page 2](images/page_2_img_022.pbm)

![Image from page 2](images/page_2_img_023.pbm)

![Image from page 2](images/page_2_img_024.pbm)

![Image from page 2](images/page_2_img_025.pbm)

![Image from page 2](images/page_2_img_026.pbm)

![Image from page 2](images/page_2_img_027.pbm)

![Image from page 2](images/page_2_img_028.pbm)

![Image from page 2](images/page_2_img_029.pbm)

![Image from page 2](images/page_2_img_030.pbm)

![Image from page 2](images/page_2_img_031.pbm)

