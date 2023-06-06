package p1361validatebinarytreenodes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_validateBinaryTreeNodes(t *testing.T) {
	for i, tc := range []struct {
		n          int
		leftChild  []int
		rightChild []int
		want       bool
	}{
		{4, []int{1, -1, 3, -1}, []int{2, -1, -1, -1}, true},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, validateBinaryTreeNodes(tc.n, tc.leftChild, tc.rightChild))
		})
	}
}

func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	// The node that isn't present in left/right children must be the root
	seen := make([]bool, n)
	for i := range leftChild {
		if leftChild[i] != -1 {
			seen[leftChild[i]] = true
		}
		if rightChild[i] != -1 {
			seen[rightChild[i]] = true
		}
	}
	root := -1
	for i := range seen {
		if !seen[i] {
			if root != -1 {
				return false
			}
			root = i
		}
	}
	if root == -1 {
		return false
	}

	k := 1
	for i := range seen {
		seen[i] = false
	}
	seen[root] = true
	q := []int{root}
	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		l := leftChild[x]
		r := rightChild[x]
		if l != -1 {
			if seen[l] {
				return false
			}
			seen[l] = true
			q = append(q, l)
			k++
		}
		if r != -1 {
			if seen[r] {
				return false
			}
			seen[r] = true
			q = append(q, r)
			k++
		}
	}
	return k == n
}
