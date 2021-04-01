package apple1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isStrobogrammatic(t *testing.T) {
	for _, tc := range []struct {
		num  string
		want bool
	}{
		{"69", true},
		{"88", true},
		{"2", false},
		{"962", false},
		{"1", true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.num), func(t *testing.T) {
			require.Equal(t, tc.want, isStrobogrammatic(tc.num))
		})
	}
}

func isStrobogrammatic(num string) bool {
	// 0, 1, 6, 9 looks the same up and down
	n := len(num)
	stack := make([]byte, 0)
	var i int
	for i = 0; i < n/2; i++ {
		switch num[i] {
		case '0', '1', '8':
			stack = append(stack, num[i])
		case '6':
			stack = append(stack, '9')
		case '9':
			stack = append(stack, '6')
		default:
			return false
		}
	}
	if n%2 == 1 {
		switch num[i] {
		case '0', '1', '8':
		default:
			return false
		}
		i++
	}
	for ; i < n; i++ {
		if len(stack) == 0 || stack[len(stack)-1] != num[i] {
			return false
		}
		stack = stack[:len(stack)-1]
	}
	return true
}
