package p3025findthenumberofwaystoplacepeoplei

import "sort"

func numberOfPairs(points [][]int) int {
	// There are only 50 points in total..
	//
	// So let's do a clever brute-force.
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] > points[j][1]
		}
		return points[i][0] < points[j][0]
	})

	seen := make([]int, 0, len(points))
	var res int
	for i := range points {
		seen = seen[:0]
		y := points[i][1]
	outer:
		for j := i + 1; j < len(points); j++ {
			yy := points[j][1]
			seen = append(seen, yy)
			if yy > y {
				continue
			}
			for _, yyy := range seen[:len(seen)-1] {
				if yyy <= y && yyy >= yy {
					continue outer
				}
			}
			res++
		}
	}
	return res
}
