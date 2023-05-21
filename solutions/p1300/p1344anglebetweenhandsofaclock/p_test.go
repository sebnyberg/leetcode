package p1344anglebetweenhandsofaclock

import "math"

func angleClock(hour int, minutes int) float64 {
	hour %= 12
	hpos := float64(hour*60+minutes) / (60 * 12)
	minpos := float64(minutes) / 60
	delta := math.Abs(hpos - minpos)
	deg := delta * 360
	return math.Min(deg, 360-deg)
}
