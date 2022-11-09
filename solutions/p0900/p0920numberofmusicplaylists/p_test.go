package p0920numberofmusicplaylists

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numMusicPlaylists(t *testing.T) {
	for i, tc := range []struct {
		n    int
		goal int
		k    int
		want int
	}{
		{2, 3, 1, 2},
		{3, 3, 1, 6},
		{2, 3, 0, 6},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, numMusicPlaylists(tc.n, tc.goal, tc.k))
		})
	}
}

func numMusicPlaylists(n int, goal int, k int) int {
	// This is quite a difficult problem. If you think long and hard enough, you
	// notice certain patterns.
	//
	// Keeping track of exactly which playlists have been formed in the previous
	// step is too complex (as indicated by the modulo in the problem
	// description).
	//
	// We must somehow find a definition of state such that we can consider only
	// the last song in the playlist. Let's say that it's the song at index "i".
	//
	// Which prior playlists can we add that song to?
	//
	// Here you might think: "well, that depends on which song it is!" That's
	// true, but all songs are equal in the eyes of the playlist maker, so we
	// don't need to know if it's song 1 or 2, we just need to know that the
	// specific song was not put in the most recent k slots.
	//
	// So let's imagine that there was no requirement to have all songs in the
	// playlist. How many ways could we form a playlist with "goal" songs
	// without putting a certain song too often?
	//
	// For the first song, we have n choices.
	// Second song, max(n-1, n-k) choices. Third. max(n-2, n-k).
	//
	// However, we must also keep track of how many songs we have introduced. In
	// the end, we must both follow this "k distance" rule, and keep track of
	// the total number of songs.
	//
	// Let's once again consider the first song. There are n choices and we now
	// have 1 unique song in the list.
	//
	// When we add the second song, we may either add a new song or an old one
	// depending on k. If we add a new song, then we have n-1 choices to choose
	// from. If we add an old one, there is only one choice (and it might not
	// even be valid according to k)
	//
	// This gives us the state for this problem:
	//
	// dp[m][l] = number of ways to have a playlist containing m distinct songs
	// and a total length of m.
	//
	const mod = 1e9 + 7
	dp := make([][]int, n+1)
	for m := range dp {
		dp[m] = make([]int, goal+1)
	}
	dp[0][0] = 1
	for m := 1; m <= n; m++ {
		for l := m; l <= goal; l++ {
			dp[m][l] = (dp[m-1][l-1]*(n-m+1) + dp[m][l-1]*max(0, m-k)) % mod
		}
	}
	return dp[n][goal]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
