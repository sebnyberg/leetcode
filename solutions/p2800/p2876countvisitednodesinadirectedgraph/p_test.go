package p2876countvisitednodesinadirectedgraph

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countVisitedNodes(t *testing.T) {
	for i, tc := range []struct {
		edges []int
		want  []int
	}{
		{[]int{1, 2, 0, 0}, []int{3, 3, 3, 4}},
		{[]int{1, 2, 3, 4, 0}, []int{5, 5, 5, 5, 5}},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, countVisitedNodes(tc.edges))
		})
	}
}

func countVisitedNodes(edges []int) []int {
	// It is a pretty strange question, because the graph could just as well
	// have more than one outgoing edge for a given node, but the way the
	// question is phrased, I assume that is not the case.
	//
	// In any case, I believe that we could keep a list of number of reachable
	// nodes and a counter of the current node during the current search. If we
	// find a node that we've seen before, then we know that the difference in
	// the counter is the size of that strongly connected component, and we may
	// pop elements from the stack, recording the result until the stack is
	// empty. In other words, we start with a leaf node, DFS until reaching the
	// end or finding a new
	n := len(edges)
	sizes := make([]int, n)
	for i := range sizes {
		sizes[i] = 1
	}
	// For each leaf node
	indeg := make([]int, n)
	for i := range edges {
		indeg[i]++
	}
	curr := []int{}
	for i := range indeg {
		if indeg[i] == 0 {
			curr = append(curr, i)
		}
	}

	res := make([]int, n)
	time := make([]int, n)
	for i := range res {
		if res[i] != 0 {
			continue
		}

		// If we reach a node with a non-zero result, then we pop the stack of
		// nodes and add 1 to each result on each round.
		// If we reach a node with a zero result but a non-zero time, then we
		// have found a cycle in the current search. This means we pop all
		// elements until we find the cycle starter and set the result to the
		// cycle length.
		stack := []int{}
		for t := 1; ; t++ {
			stack = append(stack, i)
			time[i] = t
			next := edges[i]
			if res[next] == 0 && time[next] == 0 {
				i = next
				continue
			}
			if time[next] != 0 {
				cycleLength := t - time[next] + 1
				for stack[len(stack)-1] != next {
					res[stack[len(stack)-1]] = cycleLength
					time[stack[len(stack)-1]] = 0
					stack = stack[:len(stack)-1]
				}
				res[stack[len(stack)-1]] = cycleLength
				time[stack[len(stack)-1]] = 0
				next = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
			// Pop stack and add 1 to the result for each pop.
			for k := res[next] + 1; len(stack) > 0; k++ {
				res[stack[len(stack)-1]] = k
				time[stack[len(stack)-1]] = 0
				stack = stack[:len(stack)-1]
			}
			break
		}
	}
	return res
}

type Tarjan struct {
	time       int
	timestamps []int
	low        []int
}

func (t *Tarjan) dfs(adj [][]int, cur int) {
	t.timestamps[cur] = t.time
	t.low[cur] = t.time
	t.time++

	for _, nei := range adj[cur] {
		// If the neighbour is unseen, continue SCC search and set
		// the current low-link value
		if t.timestamps[nei] == 0 {
			t.dfs(adj, nei)
			t.low[cur] = min(t.low[cur], t.low[nei])
		} else {
			// Seen before and part of this scc, use timestamp as lowlink
			t.low[cur] = min(t.low[cur], t.timestamps[nei])
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
