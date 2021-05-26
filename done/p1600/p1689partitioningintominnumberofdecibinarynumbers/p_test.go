package p1689partitioningintominnumberofdecibinarynumbers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minPartitions(t *testing.T) {
	for _, tc := range []struct {
		n    string
		want int
	}{
		{"32", 3},
		{"82734", 8},
		{"27346209830709182346", 9},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, minPartitions(tc.n))
		})
	}
}

func minPartitions(n string) int {
	var maxVal byte
	for i := range n {
		if n[i]-'0' > maxVal {
			maxVal = n[i] - '0'
		}
	}
	return int(maxVal)
}
