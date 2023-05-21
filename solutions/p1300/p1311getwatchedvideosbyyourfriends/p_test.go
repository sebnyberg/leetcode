package p1311getwatchedvideosbyyourfriends

import (
	"fmt"
	"sort"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_watchedVideosByFriends(t *testing.T) {
	for i, tc := range []struct {
		watchedVideos [][]string
		friends       [][]int
		i             int
		level         int
		want          []string
	}{
		{
			[][]string{{"A", "B"}, {"C"}, {"B", "C"}, {"D"}},
			leetcode.ParseMatrix("[[1,2],[0,3],[0,3],[1,2]]"),
			0, 2,
			[]string{"D"},
		},
		{
			[][]string{{"A", "B"}, {"C"}, {"B", "C"}, {"D"}},
			leetcode.ParseMatrix("[[1,2],[0,3],[0,3],[1,2]]"),
			0, 1,
			[]string{"B", "C"},
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, watchedVideosByFriends(tc.watchedVideos, tc.friends, tc.i, tc.level))
		})
	}
}

func watchedVideosByFriends(watchedVideos [][]string, friends [][]int, id int, level int) []string {
	curr := []int{id}
	next := []int{}
	freq := make(map[string]int)
	addVideos := func(vids []string) {
		for _, v := range vids {
			freq[v]++
		}
	}
	n := len(friends)
	seen := make([]bool, n)
	seen[id] = true
	addVideos(watchedVideos[id])
	for k := 0; k < level; k++ {
		for k := range freq {
			delete(freq, k)
		}
		if len(curr) == 0 {
			break
		}
		for _, x := range curr {
			for _, y := range friends[x] {
				if seen[y] {
					continue
				}
				seen[y] = true
				addVideos(watchedVideos[y])
				next = append(next, y)
			}
		}
		curr, next = next, curr
	}
	type s struct {
		name  string
		count int
	}
	ss := make([]s, 0, len(freq))
	for name, count := range freq {
		ss = append(ss, s{name, count})
	}
	sort.Slice(ss, func(i, j int) bool {
		a := ss[i]
		b := ss[j]
		if a.count == b.count {
			return a.name < b.name
		}
		return a.count < b.count
	})
	res := make([]string, len(freq))
	for i := range res {
		res[i] = ss[i].name
	}
	return res
}
