# Path Sum III

## Intuition


* Input size is 1000, so O(n^2) is likely alright unless there is a tight
  requirement on acceptance time.
* Problem looks like a typical "find subarrays / subsequences with sum equal to
  target sum" but for a graph.

## Naïve solution

Do post-order traversal (left, right, root) of the graph. When visiting
left/right nodes, collect possible sums from those paths, then merge the two
paths together at the current node.

```go
func pathSum(root *TreeNode, targetSum int) int {
    res, _ := visit(root, targetSum)
    return int(res)
}

func visit(node *TreeNode, targetSum int) (uint32, map[int]uint32) {
    if node == nil {
        return 0, map[int]uint32{}
    }

    lcount, lsums := visit(node.Left, targetSum)
    rcount, rsums := visit(node.Right, targetSum)
    res := make(map[int]uint32, len(lsums)+len(rsums))
    for k, v := range lsums {
        res[k+node.Val] += v
    }
    for k, v := range rsums {
        res[k+node.Val] += v
    }
    res[node.Val]++

    return res[targetSum] + lcount + rcount, res
}
```

## Prefix sums

Consider a single path in the graph where we are visiting the last node (4):

```bash
path = [10, 5, 3, 4]
```

The current sum is 22

Consider all possible prefix sums that start with the root node:

```txt
path = [10,    5,    3,    3]
       |-----11-------|
       |---15----|
       |-10-|
```

With this setup, the answer is easy to find:

* If the current sum is equal to target => increment count
* For each prefix sum that can be removed from the current sum to form target =>
  increment count

### Solution

```go
func pathSum(root *TreeNode, targetSum int) int {
    parentSums := make(map[int]uint32, 0)
    res := visit(root, parentSums, 0, targetSum)
    return int(res)
}

func visit(node *TreeNode, parentSums map[int]uint32, curr, target int) uint32 {
    if node == nil {
        return 0
    }
    curr += node.Val
    res := parentSums[curr-target]
    if curr == target {
        res++
    }
    parentSums[curr]++
    left := visit(node.Left, parentSums, curr, target)
    right := visit(node.Right, parentSums, curr, target)
    parentSums[curr]--
    return res + left + right
}
```
