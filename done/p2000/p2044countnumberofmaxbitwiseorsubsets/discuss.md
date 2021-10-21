# Intuition

* Looking at constraints, there are up to 16 numbers => to 2^16 different
  permutations => ~10^5 operations. This is fast enough to try brute-force.
* Brute-force finding all permutations is usually done with DFS. Visit each
  number and pick/don't pick.

# Naïve DFS - O(2^n)

```go
func countMaxOrSubsets(nums []int) int {
    var maxOR int
    for _, n := range nums {
        maxOR |= n
    }
    return dfs(nums, 0, len(nums), 0, maxOR)
}

func dfs(nums []int, i, n, cur, want int) int {
    if i == n {
        if cur == want {
            return 1
        }
        return 0
    }
    return dfs(nums, i+1, n, cur|nums[i], want) +
        dfs(nums, i+1, n, cur, want)
}
```

# Dynamic Programming - O(m*n)

The idea here is to count the number of permutations that lead to a certain OR
value rather than trying all permutations.

For each number in the input and each possible value that could be obtained from
ORing elements in the input.

DP is actually a lot slower than DFS for this problem. If the first element is
large, then it will be very expensive to count permutations for all following
elements. Sorting helps somewhat with execution time.

```go
func countMaxOrSubsets(nums []int) int {
    var dp [1 << 17]int
    var maxOR int
    dp[0] = 1
    for _, num := range nums {
        for a := maxOR; a >= 0; a-- {
            dp[a|num] += dp[a]
        }
        maxOR |= num
    }
    return dp[maxOR]
}
```
