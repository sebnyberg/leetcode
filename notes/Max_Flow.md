# Max Flow

## Ford-Fulkerson

Given a DAG where each edge has a capacity (c), find the maximmum capacity that could flow through the graph.

Of course this is bounded by the output of the source output, and the sink input, but as a whole there could be other complications along the mid-section of the graph.

There may be multiple optimal solutions, but Ford-Fulkerson is guaranteed to find one of them.

An augmenting path is a path from the source to the sink that changes reduces capacity of edges along that path, i.e. it updates the overall flow in the graph.

Augmenting the flow is the process of finding a path (e.g. by DFS), then increasing the flow by the bottleneck value along the path.

An augmenting path may not be part of the optimal solution, thus it is important to be able to un-do the path as needed.

For this reason, when increasing the flow along an augmenting path, it is necessary to create residual (backward) edges with the negative amount of flow to allow for un-doing bad choices.

The remaining capacity of an edge is "capacity - flow". For residual edges, the capacity is zero, and if the flow is negative, then the capacity is positive.

The graph containing residual edges is called the "residual graph".

For DFS: O(f*E), f is the maximum flow, E is the number of edges

Argument: it is guaranteed that an augmenting path increases the flow, but depending on how the flow is picked, the flow may increase with as little as 1. If that path contains all edges, then the worst case becomes O(fE).

Edmonds-Karp: BFS - shortest path from source / sink O(E^2*V).
Capacity scaling: Pick larger paths first O(E^2log(U))

1. Do a search, only pick edges with capacity > X 
2. If no match, X /= 2