# Explanation

Use Dijkstra's, but find the two min times rather than one. When visiting a node with a total time that is the same as before, or greater than the second smallest time, then skip the visit.

# Perf notes

* The maximum possible time for a second visit is (10^4+2) * 10^3 ~= 10^7, so it's OK to use an uint32 instead of int. Recall that ints are a "machine word" in size, which on Leetcode machines is 64 bits. That's 32 bits of wasted space
* Indices should be uint16 (1-10^4). For the weightedEdge struct it will not help because of struct padding / memory alignment, but it reduces the size of the adjacency list.

# Solution

```go
func secondMinimum(n int, edges [][]int, time int, change int) int {
	// Perform Dijkstra's with the modification that each node can be visited
	// twice, and there are two (unique) min times.
	adj := make([][]uint16, n+1)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		adj[a] = append(adj[a], uint16(b))
		adj[b] = append(adj[b], uint16(a))
	}

	minTimes := make([][2]uint32, n+1)
	for i := range minTimes {
		minTimes[i] = [2]uint32{math.MaxInt32, math.MaxInt32}
	}

	h := DistanceHeap{weightedEdge{0, 1}}
	for len(h) > 0 {
		e := heap.Pop(&h).(weightedEdge)
		u, t := e.node, e.arrivalTime
		if t == minTimes[u][0] || t >= minTimes[u][1] {
			continue
		}

		if t < minTimes[u][0] {
			minTimes[u][0] = t
		} else {
			minTimes[u][1] = t
		}

		// If light is currently red, wait for a green signal.
		d := t / uint32(change)
		if isGreen := d%2 == 0; !isGreen {
			nextGreen := (d + 1) * uint32(change)
			t = max(t, nextGreen)
		}
		for _, nei := range adj[u] {
			if t+uint32(time) >= minTimes[nei][1] {
				continue
			}
			heap.Push(&h, weightedEdge{
				arrivalTime: t + uint32(time),
				node:        nei,
			})
		}
	}
	return int(minTimes[n][1])
}

func max(a, b uint32) uint32 {
	if a > b {
		return a
	}
	return b
}

type weightedEdge struct {
	arrivalTime uint32
	node        uint16
}

type DistanceHeap []weightedEdge

func (h DistanceHeap) Len() int { return len(h) }
func (h DistanceHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h DistanceHeap) Less(i, j int) bool {
	return h[i].arrivalTime < h[j].arrivalTime
}
func (h *DistanceHeap) Push(x interface{}) {
	*h = append(*h, x.(weightedEdge))
}
func (h *DistanceHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}
```
