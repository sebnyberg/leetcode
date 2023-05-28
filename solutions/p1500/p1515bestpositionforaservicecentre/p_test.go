package p1515bestpositionforaservicecentre

import (
	"math"
)

const eps = 1e-7

func getMinDistSum(positions [][]int) float64 {
	// Not really a programming problem but, hey.
	//
	// Let's use Weiszfeld's algorithm.
	//
	n := len(positions)
	xs := make([]float64, n)
	ys := make([]float64, n)
	var xsum float64
	var ysum float64
	for i, p := range positions {
		xs[i] = float64(p[0])
		xsum += xs[i]
		ys[i] = float64(p[1])
		ysum += ys[i]
	}
	// The starting point can be anything but a good guess is the euclidian
	// mean.
	x := xsum / float64(len(positions))
	y := ysum / float64(len(positions))

	prev := math.MaxFloat32
	for {
		var distsum float64
		var recipsum float64
		var a float64
		var b float64
		for i := range positions {
			dx := x - xs[i]
			dy := y - ys[i]
			d := math.Sqrt(dx*dx + dy*dy)
			distsum += d
			if d == 0 {
				// avoid division by zero
				d = 1
			}
			recipsum += 1 / d
			a += xs[i] / d
			b += ys[i] / d
		}
		nextX := a / recipsum
		nextY := b / recipsum
		if prev-distsum < eps {
			return distsum
		}
		prev = distsum
		x = nextX
		y = nextY
	}
}
