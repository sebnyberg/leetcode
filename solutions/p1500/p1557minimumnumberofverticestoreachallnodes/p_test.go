package p1557minimumnumberofverticestoreachallnodes

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	// Consider first any node that does not have any inbound edges. Such a node
	// must be part of the result.
	// Also any node that has an inbound edge must by definition be reachable
	// from another node.
	// For strongly connected components, it does not matter which edge is
	// chosen.
	//
	// This leads to the algorithm: starting with nodes without inbound edges,
	// perform BFS until no more nodes can be reached. Add these starting nodes
	// to the result.
	// Then iterate over unvisited nodes, adding each node to the result and
	// mark all reachable nodes as visited.
	//
	adj := make([][]int, n)
	indeg := make([]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		indeg[e[1]]++
	}

	visited := make([]bool, n)
	curr := []int{}
	next := []int{}
	bfs := func(startNode int) {
		visited[startNode] = true
		curr = append(curr[:0], startNode)
		for len(curr) > 0 {
			next = next[:0]

			for _, x := range curr {
				for _, y := range adj[x] {
					if visited[y] {
						continue
					}
					visited[y] = true
					next = append(next, y)
				}
			}

			curr, next = next, curr
		}
	}
	var res []int
	for i, deg := range indeg {
		if deg == 0 {
			res = append(res, i)
			bfs(i)
		}
	}
	for i := range visited {
		if !visited[i] {
			res = append(res, i)
			bfs(i)
		}
	}
	return res
}
