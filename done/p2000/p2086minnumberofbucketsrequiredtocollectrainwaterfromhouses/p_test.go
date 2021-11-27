package p2086minnumberofbucketstocollectrainwaterfromhouses

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimumBuckets(t *testing.T) {
	for _, tc := range []struct {
		street string
		want   int
	}{
		{".H.H.", 1},
		{"H..H", 2},
		{".HHH.", -1},
		{"H", -1},
		{".", 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.street), func(t *testing.T) {
			require.Equal(t, tc.want, minimumBuckets(tc.street))
		})
	}
}

func minimumBuckets(street string) int {
	bucketPos := -2
	var buckets int
	n := len(street)
	for i, ch := range street {
		if ch != 'H' || bucketPos == i-1 {
			continue
		}
		// Try place after
		if i != n-1 && street[i+1] != 'H' {
			bucketPos = i + 1
		} else if i != 0 && street[i-1] != 'H' {
			bucketPos = i - 1
		} else {
			return -1
		}
		buckets++
	}
	return buckets
}
