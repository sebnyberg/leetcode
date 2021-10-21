package p2011finalvalueofvariableafterperformingoperations

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_finalValueAfterOperations(t *testing.T) {
	for _, tc := range []struct {
		operations []string
		want       int
	}{
		{[]string{"--X", "X++", "X++"}, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.operations), func(t *testing.T) {
			require.Equal(t, tc.want, finalValueAfterOperations(tc.operations))
		})
	}
}

func finalValueAfterOperations(operations []string) int {
	var val int
	for _, op := range operations {
		switch op {
		case "++X", "X++":
			val++
		case "--X", "X--":
			val--
		}
	}
	return val
}
