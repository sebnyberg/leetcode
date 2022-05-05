package p2065maximumpathqualityofagraph

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func ParseTest(fname string) ([]int, [][]int, int) {
	f, err := os.OpenFile(fname, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	contents, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln(err)
	}
	parts := bytes.Split(contents, []byte{'\n'})
	valueStrings := strings.Split(string(parts[0]), ",")
	values := make([]int, len(valueStrings))
	for i := range valueStrings {
		res, err := strconv.ParseInt(valueStrings[i], 10, 64)
		if err != nil {
			log.Fatalln(err)
		}
		values[i] = int(res)
	}
	edgeStrs := bytes.Split(parts[1], []byte("],["))
	n := len(edgeStrs)
	edgeStrs[0] = edgeStrs[0][1:]
	edgeStrs[n-1] = edgeStrs[n-1][:len(edgeStrs[n-1])-1]
	edges := make([][]int, len(edgeStrs))
	for i := range edgeStrs {
		edgeParts := strings.Split(string(edgeStrs[i]), ",")
		u, err := strconv.Atoi(edgeParts[0])
		if err != nil {
			log.Fatalln(err)
		}
		v, err := strconv.Atoi(edgeParts[1])
		if err != nil {
			log.Fatalln(err)
		}
		time, err := strconv.Atoi(edgeParts[2])
		if err != nil {
			log.Fatalln(err)
		}
		edges[i] = []int{u, v, time}
	}
	maxTime, err := strconv.Atoi(string(parts[2]))
	if err != nil {
		log.Fatalln(err)
	}
	return values, edges, maxTime
}

func Test_maximalPathQuality(t *testing.T) {
	aVal, aEdge, aMaxTime := ParseTest("testdata/test1")

	for _, tc := range []struct {
		values  []int
		edges   [][]int
		maxTime int
		want    int
	}{
		{aVal, aEdge, aMaxTime, 0},
		{[]int{0, 32, 10, 43}, [][]int{{0, 1, 10}, {1, 2, 15}, {0, 3, 10}}, 49, 75},
		{[]int{5, 10, 15, 20}, [][]int{{0, 1, 10}, {1, 2, 10}, {0, 3, 10}}, 30, 25},
		{[]int{1, 2, 3, 4}, [][]int{{0, 1, 10}, {1, 2, 11}, {2, 3, 12}, {1, 3, 13}}, 50, 7},
		{[]int{0, 1, 2}, [][]int{{1, 2, 10}}, 10, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.values), func(t *testing.T) {
			require.Equal(t, tc.want, maximalPathQuality(tc.values, tc.edges, tc.maxTime))
		})
	}
}

func maximalPathQuality(values []int, edges [][]int, maxTime int) int {
	// Since time is low, and maxtime is low, we can try brute-force.
	// It's 3^10
	n := len(values)
	adj := make([][]int, n)
	time := make([][]int, n)
	for _, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		adj[edge[1]] = append(adj[edge[1]], edge[0])
		time[edge[0]] = append(time[edge[0]], edge[2])
		time[edge[1]] = append(time[edge[1]], edge[2])
	}

	seen := make([]bool, n)
	var r maxValRecorder
	seen[0] = true
	r.dfs(values, adj, time, seen, 0, values[0], maxTime)
	return r.maxVal
}

type maxValRecorder struct {
	maxVal int
}

func (r *maxValRecorder) dfs(values []int, adj, time [][]int, seen []bool, i, curVal, curTime int) {
	if i == 0 {
		r.maxVal = max(r.maxVal, curVal)
	}

	// Try all possible paths from the current node
	for j, nei := range adj[i] {
		if curTime < time[i][j] {
			continue
		}
		if !seen[nei] {
			seen[nei] = true
			r.dfs(values, adj, time, seen, nei, curVal+values[nei], curTime-time[i][j])
			seen[nei] = false
		} else {
			r.dfs(values, adj, time, seen, nei, curVal, curTime-time[i][j])
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
