# Daily Coding Problem: Problem #604 [Hard]

This problem was asked by Grammarly.

Soundex is an algorithm used to categorize phonetically, such that two
names that sound alike but are spelled differently have the same
representation.

Soundex maps every name to a string consisting of one letter and three
numbers, like M460.

One version of the algorithm is as follows:

Remove consecutive consonants with the same sound (for example, change ck -> c).

Keep the first letter. The remaining steps only apply to the rest of the string.

Remove all vowels, including y, w, and h.

Replace all consonants with the following digits:

        b, f, p, v &rarr; 1
        c, g, j, k, q, s, x, z &rarr; 2
        d, t &rarr 3
        l &rarr; 4
        m, n &rarr; 5
        r &rarr; 6

If you don't have three numbers yet, append zeros until you do.
Keep the first three numbers.

Using this scheme, Jackson and Jaxen both map to J250.

Implement Soundex.

## Analysis

The algorithm outlined above seems to be based on
[American Soundex](https://en.wikipedia.org/wiki/Soundex).

American Soundex Rule #3:

    If two or more letters with the same number are adjacent in the original
    name (before step 1), only retain the first letter; also two letters
    with the same number separated by 'h' or 'w' are coded as a single
    number, whereas such letters separated by a vowel are coded twice. This
    rule also applies to the first letter.

gets replaced by

	Remove consecutive consonants with the same sound (for example,
    change ck -> c).

Why? Wikipedia's Rule #3 is objectively easier to implement,
"the same sound" doesn't really mean anything.
ck &rarr; c, but does cks &rarr; cs &rarr; c?
What about 'g' and 'j'?

I'm going to guess that Grammarly was interested in finding out
if or how candidates did the "sound alike" part.
The rest of the work is pretty basic character-by-character
examination of an input string.

If you, the interviewer, aren't working for Grammarly,
this is not a great question.
You won't find anything out about the candidate's programming skills.
