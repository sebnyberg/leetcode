package p0939minimumarearectangle

import "math"

func minAreaRect(points [][]int) int {
	m := make(map[[2]int]struct{})
	for _, p := range points {
		m[[2]int{p[0], p[1]}] = struct{}{}
	}

	res := math.MaxInt32
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			p := points[i]
			q := points[j]
			dx := abs(p[0] - q[0])
			dy := abs(p[1] - q[1])
			area := dx * dy
			if area > 0 && area < res {
				_, ok1 := m[[2]int{p[0], q[1]}]
				_, ok2 := m[[2]int{q[0], p[1]}]
				if ok1 && ok2 {
					res = area
				}
			}
		}
	}
	if res != math.MaxInt32 {
		return res
	}
	return 0
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
