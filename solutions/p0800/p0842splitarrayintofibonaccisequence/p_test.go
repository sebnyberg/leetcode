package p0842splitarrayintofibonaccisequence

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_splitIntoFibonacci(t *testing.T) {
	for _, tc := range []struct {
		num  string
		want []int
	}{
		{"1101111", []int{11, 0, 11, 11}},
		{"112358130", []int{}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.num), func(t *testing.T) {
			require.Equal(t, tc.want, splitIntoFibonacci(tc.num))
		})
	}
}

func splitIntoFibonacci(num string) []int {
	parse := func(a string) (int, bool) {
		if len(a) == 0 || (len(a) > 0 && a[0] == '0') {
			return 0, false
		}
		x, err := strconv.ParseInt(a, 10, 32)
		if err != nil {
			return 0, false
		}
		return int(x), true
	}

	res := make([]int, 0)
	for j1 := 1; j1 <= 9; j1++ {
		for k1 := j1 + 1; k1 <= j1+9; k1++ {
			res = res[:0]
			j, k := j1, k1
			if k > len(num) {
				continue
			}
			a, _ := parse(num[:j])
			b, _ := parse(num[j:k])
			res = append(res, a, b)
			for {
				c := a + b
				res = append(res, c)
				s := fmt.Sprint(c)
				if c >= 1<<31 || k+len(s) > len(num) || num[k:k+len(s)] != s {
					break
				}
				a, b = b, c
				k += len(s)
				if k == len(num) {
					return res
				}
			}
		}
	}
	return []int{}
}
