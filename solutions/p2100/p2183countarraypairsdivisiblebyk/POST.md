For this problem, we wish to find all pairs of numbers `(a,b)` such that `a*b %
k == 0`. It's clear that doing this for all pairs is too expensive, so we need
to somehow group numbers together.

First, some background.

All numbers have a prime factorization. For example, `192 = 2*2*2*7*7`. For
prime numbers, the factor is simply itself times 1.

The greatest common denominator of two numbers, i.e. `gcd(a,b)` is the common
prime factorizations of a and b. For example, given `a = 192 = 2*2*2*7*7` and `b
= 21 = 3*7`, the GCD is `7`.

When we take the GCD (`a,k`) of each element `a` in `nums`, we remove all
factors from each element, keeping only those prime factors that are in common
with `k`. We know that there aren't that many of these factors - in fact, we can
find all factors that evenly divide k in `O(sqrt(N))` time, so there are *at
most* sqrt(N) factors.

Let's take an example from one of the test cases:

```
nums = [8, 10, 2, 5, 9, 6, 3, 8, 2]
k = 6
```

Now, `k = 6 = 3*2`. The only possible factors that divide k are `6, 3, 2, 1`, so
there can be at most 4 different numbers after doing `gcd(a,k)`:

```
nums = [8, 10, 2, 5, 9, 6, 3, 8, 2]
factors = [2*2*2, 2*5, 2, 5, 3*3, 3*2, 3, 2*2*2, 2]
k = 2*3
common_factors = [2, 2, 2, 1, 3, 6, 3, 2, 2]
```

Notice how few factors there are.

Now comes the interesting part: since `gcd(a,k)` captures common prime factors
between `a` and `k`, then any `gcd(a,k) * gcd(b,k)` where `a` and `b` contain
all factors in `k` must be evenly divisible by `k`. If a factor exists e.g.
twice, it will just create a larger number still divisible by `k`.

Therefore, `gcd(a,k)*gcd(b,k) % k = 0` if `(a * b) % k = 0`.

The final tricky part is that if we count each gcd value, then we must be
careful not to double-count pairs, and when the number itself is divisible by k,
we need to pair it with itself.

# Solution

```go
func countPairs(nums []int, k int) int64 {
	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}
    
	n := int(math.Sqrt(float64(k)))
	gcds := make(map[int]int, n)
	for _, num := range nums {
		gcds[gcd(num, k)]++
	}

	var res int
	for a, n1 := range gcds {
		for b, n2 := range gcds {
			if a > b || (a*b)%k != 0 {
				continue
			}
			if a != b {
				res += n1 * n2
			} else { // a == b
				res += n1 * (n1 - 1) / 2
			}
		}
	}
	return int64(res)
}
```