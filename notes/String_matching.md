# String Matching

## Rolling Hash

A rolling hash is a function which can compute a hash from a sliding window of symbols in an efficient manner.

### Basic example

A very basic example is that of a sequence of numbers 0-9, like this:

```
[1,5,4,6,2,5,4]
```

If we wish to find all subarrays (contiguous subsequences) of length k that are equal, we may start with the following hash:

```
1*10+5 = 15
```

Then, sliding the window to the right, it is only a matter of re-calculating the hash:

```
15 - arr[i-1]*10 = 5
5 * 10 + 4 = 14
```

### Prime modulo

Before going into this section, I recommend watching this video on prime modulo and it's operations:

https://www.youtube.com/watch?v=-OPohCQqi_E

As the base and window grows, so does the maximum possible hash values. For all letters of the alphabet, the base is 26. This is roughly 2^5, so with a window greater than 12 in size, the maximum possible hash value would exceed a 64-bit integer.

A solution to this problem is to introduce the use of prime modulus.

Given the numbers `a` and `b`:

```
(a * b) % mod = a % mod * b % mod
```

The interesting thing about this is that previously massive numbers can now be modulo'd into much smaller values.

It is also possible to subtract from the hash. To make sure that this operation works for the case where the subtraction yields a negative number, it is necessary to add the modulo and do an extra modulo operation:

```
h1 = (h1 - a*base*mod + mod) % mod
```

### Picking a prime modulo and base

Picking a prime is not an exact art. There may be a larger prime with more collisions, and vice versa. In general, however, a good starting point is to:

1. Use modulo 1e9+7. This fits within a 32-bit integer and is easy to remember.
2. Use the first prime greater than the required base size as the base. 26 characters => base 29.

### Resolving collisions

In some cases collisions will occur. To manage this, it is best to store matching indices in the lists that are compared, then do a good ol' list compare to verify that the hash was correct. The hash algorithm will still be able to rule out a large number of non-matches, so this is still quite efficient.

## LPS / KMP

Building the table:

```go
lps := make([]int, len(s))
lps[0] = 0
i, j := 0, 1
for j < len(s){
    if s[i] == s[j]{
        i++
        lps[j] = i
        j++
    } else if i==0{
        lps[j] = 0
        j++
    } else {
        i = lps[i-1]
    }
}
```
