# DSU & Union Find

## DSU - Disjoint Set Union

[Video](https://www.youtube.com/watch?v=0jNmHPfA_yE)

The basic idea is to form a root node for each component group found in the input.

For example, if somehow nodes 1, 6, 8 belong together, the end goal is to have a group such that all three elements point to the same index.

This is achieved by creating a `find` function which recursively calls itself until the parent of the current index is itself.

For example, if we arbitrarily choose to set `6` as the group index, then parent[6] = 6. When assigning 8 to the group, parent[8] will point to parent[6] which is itself (6).

When assigning 1 to the group, it does not matter if the union happens from 1 to 6 or 1 to 8, the chain will be either

```text
parent[1] = find(6) = find(parent[6]) = 6
# or
parent[1] = find(8) = find(parent[8]) = find(find(parent[6])) = 6
```

When we are done, iterating over the list of parents will yield a list of disjoint union sets.
