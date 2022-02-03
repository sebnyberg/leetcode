# Leetcode

Collection of leetcode solutions in Go.

## Algorithm dictionary

### Palindromes

* Manacher's algorithm, works similar to Knuth-Morris-Pratt LPS tables in that
  it re-uses a partial existing match when searching for new palindromes.

### String matching

* Rabin-Karp a.k.a. Rolling Hash. Allows for continuously hashing a sliding 
  window in an efficient way.
* Knuth-Morris-Pratt. Used to search for a word in a string. Usually,
  Rabin-Karp is easier so it's the preferredchoice.
* Aho-Corasick can be used to search for multiple different patterns at the same
  time.

### k-coloring

NP-hard, no good solutions. For small graphs, it's possible to speed up with DP.

### Vertex cover

The goal of vertex covering is to find the minimum set of vertices which 'see'
all other vertices.

König's theorem states that the minimal vertex cover problem for bipartite
graphs is equivalent to finding the maximum cardinality matching, which can be
done with Hopcroft-Karp-Karzanov.

### Bipartite matching

* Hopcroft-Karp-Karzanov
* Hungarian Algorithm

### Min-cut, max-flow

* Edmonds-Karp
* Ford-Fulkerson

### Eulerian path / circuit

Remember to check if a path exists.

* Hierholzer's algorithm

### Cycle detection

* Floyd's Tortoise and Hare - find cycle using fast and slow cursor, useful for
  finding cycles in e.g. linked lists
* DFS with stack of nodes. When detecting a back-edge, there is a cycle.

### Shortest path

* Floyd-Warshall - finds shortest paths between all pairs of vertices in a
  weighted graph (positive and negative weights) O(V^3). Good for dense graphs.
* Dijkstra's shortest path algorithm - finds shortest path between two nodes in
  a weighted graph (positive-only edges).
* Johnson's algorithm - shortest path for sparse, weighted graphs between nodes
  with positive and negative weights
* Bellman-Ford - easy BFS-based searching for weighted graphs with positive and
  negative weights

## Minimal Enclosing Circle

* Welzl's algorithm
