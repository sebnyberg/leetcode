package p0127wordladder

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findLadders(t *testing.T) {
	for _, tc := range []struct {
		beginWord string
		endWord   string
		wordList  []string
		want      [][]string
	}{
		{"a", "c", []string{"a", "b", "c"}, [][]string{{"a", "c"}}},
		{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}, [][]string{
			{"hit", "hot", "dot", "dog", "cog"}, {"hit", "hot", "lot", "log", "cog"},
		}},
		{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}, [][]string{}},
	} {
		t.Run(fmt.Sprintf("%v/%v/%+v", tc.beginWord, tc.endWord, tc.wordList), func(t *testing.T) {
			got := findLadders(tc.beginWord, tc.endWord, tc.wordList)
			require.Equal(t, tc.want, got)
		})
	}
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	n := len(wordList)
	m := len(wordList[0])

	// Join words together using wildcards
	// E.g. abcd will be put into the mask-lists *bcd, a*cd, ab*d and abc*
	encode := func(w []byte) uint32 {
		var res uint32
		for _, ch := range w {
			res = (res * 27) + uint32(ch-'a')
		}
		return res
	}
	const wildcard = 'a' + 26
	maskList := make(map[uint32][]uint16)
	wordList = append(wordList, beginWord)
	adj := make([][]uint16, n+1)
	addWord := func(w string, j uint16) {
		s := []byte(w)
		adj[j] = make([]uint16, 0, 10)
		for i := 0; i < m; i++ {
			s[i] = wildcard
			e := encode(s)
			for _, nei := range maskList[e] {
				adj[nei] = append(adj[nei], j)
				adj[j] = append(adj[j], nei)
			}
			maskList[e] = append(maskList[e], j)
			s[i] = w[i]
		}
	}
	beginIdx := uint16(n)
	var endIdx uint16 = math.MaxUint16
	for i, w := range wordList {
		if i < n && w == beginWord {
			continue
		}
		if w == endWord {
			endIdx = uint16(i)
		}
		addWord(w, uint16(i))
	}
	if endIdx == math.MaxUint16 {
		return [][]string{}
	}
	n++

	// Next, perform a BFS-based version of Dijkstra
	// Note! In order to quickly assemble paths from start->end, we search from
	// end->start during the path search. This way, appending during Dijkstra's
	// will lead to the correct order.
	predecessors := make([][]uint16, n)
	curr := []uint16{endIdx}
	next := []uint16{}
	dist := make([]uint16, n)
	for i := range dist {
		dist[i] = math.MaxUint16
	}
	dist[endIdx] = 0
	for k := uint16(1); len(curr) > 0; k++ {
		next = next[:0]
		for _, x := range curr {
			if x == beginIdx {
				goto done
			}
			for _, nei := range adj[x] {
				if dist[nei] < k {
					continue
				}
				if dist[nei] > k {
					next = append(next, nei)
				}
				dist[nei] = k
				predecessors[nei] = append(predecessors[nei], x)
			}
		}
		curr, next = next, curr
	}
done:

	// Collect results (if any)
	res := [][]string{}
	if dist[beginIdx] == math.MaxUint16 {
		return [][]string{}
	}
	pathLen := dist[beginIdx] + 1
	path := make([]string, pathLen)
	path[0] = beginWord
	var j int
	var buildResult func(curr uint16, i int)
	buildResult = func(curr uint16, i int) {
		if curr == endIdx {
			res = append(res, make([]string, pathLen))
			copy(res[j], path)
			j++
			return
		}
		for _, p := range predecessors[curr] {
			path[i] = wordList[p]
			buildResult(p, i+1)
		}
	}
	buildResult(beginIdx, 1)
	return res
}
