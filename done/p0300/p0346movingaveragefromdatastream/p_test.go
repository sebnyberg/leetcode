package p0346movingaveragefromdatastream

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMovingAverage(t *testing.T) {
	ma := Constructor(3)
	res := ma.Next(1)
	require.InEpsilon(t, 1.0, res, 0.01)
	res = ma.Next(10)
	require.InEpsilon(t, 5.5, res, 0.01)
	res = ma.Next(3)
	require.InEpsilon(t, 4.66667, res, 0.01)
	res = ma.Next(5)
	require.InEpsilon(t, 6.0, res, 0.01)
}

type MovingAverage struct {
	buf            []int
	pos, sum, size int
}

func Constructor(size int) MovingAverage {
	return MovingAverage{
		buf:  make([]int, size),
		size: size,
	}
}

func (this *MovingAverage) Next(val int) float64 {
	this.sum -= this.buf[this.pos%this.size]
	this.sum += val
	this.buf[this.pos%this.size] = val
	this.pos++
	return float64(this.sum) / float64(min(this.pos, this.size))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
