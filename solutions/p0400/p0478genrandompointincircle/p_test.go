package p0478genrandompointincircle

import (
	"math"
	"math/rand"
)

type Solution struct {
	radius  float64
	xCenter float64
	yCenter float64
}

func Constructor(radius float64, x_center float64, y_center float64) Solution {
	return Solution{radius: radius, xCenter: x_center, yCenter: y_center}
}

func (this *Solution) RandPoint() []float64 {
	radians := rand.Float64() * 2 * math.Pi
	radius := math.Sqrt(rand.Float64()) * this.radius
	return []float64{this.xCenter + radius*math.Cos(radians), this.yCenter + radius*math.Sin(radians)}
}
