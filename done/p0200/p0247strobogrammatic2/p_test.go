package p0247strobogrammatic2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findStrobogrammatic(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want []string
	}{
		{2, []string{"11", "69", "88", "96"}},
		{1, []string{"0", "1", "8"}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, findStrobogrammatic(tc.n))
		})
	}
}

func findStrobogrammatic(n int) []string {
	res := make([]string, 0)
	find(n, 0, "", "", &res)
	return res
}

func find(n int, i int, left string, right string, res *[]string) {
	if i == n {
		*res = append(*res, left+right)
		return
	}
	// middle
	if n-i == 1 {
		find(n, i+1, left+"0", right, res)
		find(n, i+1, left+"1", right, res)
		find(n, i+1, left+"8", right, res)
		return
	}
	if i > 0 {
		find(n, i+2, left+"0", "0"+right, res)
	}
	find(n, i+2, left+"1", "1"+right, res)
	find(n, i+2, left+"6", "9"+right, res)
	find(n, i+2, left+"9", "6"+right, res)
	find(n, i+2, left+"8", "8"+right, res)
}
