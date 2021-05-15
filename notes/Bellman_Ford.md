# Bellman Ford

Bellman-Ford is used to find the shortest path from a->b in a weighted graph.

The algorithm is slower than Dijkstra (which is O(E+VlogV)), but can find paths in graphs where edges are negative.

## Finding shortest path from s -> t 

* Initialize a distance vector where d[t] = 0 (distance from t to t is zero)
* For each edge, check if the edge weight + distance from target to t is smaller than the current value, if it is, update dist
* Iterate V-1 times
* When done, iterate one more time - if a value decreases, there is a negative cycle, set the value to negative infinity.