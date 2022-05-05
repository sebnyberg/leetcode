package p0636exclusivetimeoffunctions

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_exclusiveTime(t *testing.T) {
	for _, tc := range []struct {
		n    int
		logs []string
		want []int
	}{
		{2, []string{"0:start:0", "1:start:2", "1:end:5", "0:end:6"}, []int{3, 4}},
		{1, []string{"0:start:0", "0:start:2", "0:end:5", "0:start:6", "0:end:6", "0:end:7"}, []int{8}},
		{2, []string{"0:start:0", "0:start:2", "0:end:5", "1:start:6", "1:end:6", "0:end:7"}, []int{7, 1}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, exclusiveTime(tc.n, tc.logs))
		})
	}
}

func exclusiveTime(n int, logs []string) []int {
	stack := []process{}
	parse := func(s string) process {
		parts := strings.Split(s, ":")
		typ := parts[1]
		id, _ := strconv.Atoi(parts[0])
		t, _ := strconv.Atoi(parts[2])
		return process{
			typ: typ,
			t:   t,
			id:  id,
		}
	}

	res := make([]int, n)
	for _, l := range logs {
		p := parse(l)
		if p.typ == "start" {
			if len(stack) > 0 {
				x := stack[len(stack)-1]
				res[x.id] += p.t - x.t
			}
			stack = append(stack, p)
		} else {
			x := stack[len(stack)-1]
			res[p.id] += p.t - x.t + 1
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				stack[len(stack)-1].t = p.t + 1
			}
		}
	}
	return res
}

type process struct {
	typ string
	t   int
	id  int
}
