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


