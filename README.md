# Leetcode

Collection of leetcode solutions in Go.

## Permutations

### Lexicographically sorted permutations

[Source](https://stackoverflow.com/questions/31425531/heaps-algorithm-for-permutations)

1. Find the rightmost element which is smaller than the element to its right
2. Swap that element with the smallest element to its right which is larger than it
3. Reveerse the part of the permutation to the right of where that element was