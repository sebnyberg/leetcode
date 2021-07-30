package p0838pushdominoes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_pushDominoes(t *testing.T) {
	for _, tc := range []struct {
		dominoes string
		want     string
	}{
		{"RR.L", "RR.L"},
		{".L.R...LR..L..", "LL.RR.LLRRLL.."},
	} {
		t.Run(fmt.Sprintf("%+v", tc.dominoes), func(t *testing.T) {
			require.Equal(t, tc.want, pushDominoes(tc.dominoes))
		})
	}
}

const (
	unset = 0
	left  = 1
	right = 2
	both  = 3
)

func pushDominoes(dominoes string) string {
	n := len(dominoes)
	visited := make([]bool, n)
	state := make([]byte, n)
	todo := []int{}
	for i, d := range dominoes {
		switch d {
		case 'L':
			state[i] = left
			todo = append(todo, i)
			visited[i] = true
		case 'R':
			state[i] = right
			todo = append(todo, i)
			visited[i] = true
		default:
			state[i] = unset
		}
	}
	next := []int{}
	for len(todo) > 0 {
		next = next[:0]
		for _, idx := range todo {
			if state[idx] == left && idx >= 1 && !visited[idx-1] {
				state[idx-1] += left
				next = append(next, idx-1)
				continue
			}
			if state[idx] == right && idx < n-1 && !visited[idx+1] {
				state[idx+1] += right
				next = append(next, idx+1)
			}
		}
		for _, idx := range next {
			visited[idx] = true
		}
		todo, next = next, todo
	}
	res := make([]byte, n)
	for i, d := range state {
		switch d {
		case unset:
			res[i] = '.'
		case left:
			res[i] = 'L'
		case right:
			res[i] = 'R'
		case both:
			res[i] = '.'
		}
	}
	ss := string(res)
	return ss
}
