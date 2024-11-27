package p3243shortestdistanceafterroadadditionqueriesi

func shortestDistanceAfterQueries(n int, queries [][]int) []int {
	// O(n^2) is fine
	dist := make([]int, n)
	for i := range dist {
		dist[i] = i
	}
	m := len(queries)
	res := make([]int, m)
	adj := make([][]int, n)
	for i := range adj {
		adj[i] = []int{i + 1}
	}
	adj[n-1] = []int{}
	for i, q := range queries {
		k, l := q[0], q[1]
		adj[k] = append(adj[k], l)
		for j := range dist {
			for _, b := range adj[j] {
				dist[b] = min(dist[b], dist[j]+1)
			}
		}
		res[i] = dist[n-1]
	}
	return res
}
