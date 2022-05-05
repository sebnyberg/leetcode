package p0251flatten2dvector

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVector2D(t *testing.T) {
	vc := Constructor([][]int{{1, 2}, {3}, {4}})
	res := vc.Next()
	require.Equal(t, 1, res)
	res = vc.Next()
	require.Equal(t, 2, res)
	res = vc.Next()
	require.Equal(t, 3, res)
	boolRes := vc.HasNext()
	require.Equal(t, true, boolRes)
	boolRes = vc.HasNext()
	require.Equal(t, true, boolRes)
	res = vc.Next()
	require.Equal(t, 4, res)
	boolRes = vc.HasNext()
	require.Equal(t, false, boolRes)

	vc = Constructor([][]int{{}})
	boolRes = vc.HasNext()
	require.Equal(t, false, boolRes)

	vc = Constructor([][]int{{}, {3}})
	boolRes = vc.HasNext()
	require.Equal(t, true, boolRes)
	res = vc.Next()
	require.Equal(t, 3, res)
	boolRes = vc.HasNext()
	require.Equal(t, false, boolRes)
}

type Vector2D struct {
	i, j int
	n    int
	vec  [][]int
}

func Constructor(vec [][]int) Vector2D {
	return Vector2D{
		i:   0,
		j:   0,
		n:   len(vec),
		vec: vec,
	}
}

func (this *Vector2D) fwd() {
	if this.i == this.n {
		return
	}
	for this.i != this.n && len(this.vec[this.i]) == this.j {
		this.j = 0
		this.i++
	}
}

func (this *Vector2D) Next() int {
	el := this.vec[this.i][this.j]
	this.j++
	this.fwd()
	return el
}

func (this *Vector2D) HasNext() bool {
	this.fwd()
	return this.n > 0 && this.i != this.n
}
