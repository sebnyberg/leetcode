package p1093statisticsfromalargesample

import "math"

func sampleStats(count []int) []float64 {
	var sum float64
	var totalCount int
	var maxCount int
	var mode float64
	for i, c := range count {
		totalCount += c
		sum += float64(c) * float64(i)
		if c > maxCount {
			maxCount = c
			mode = float64(i)
		}
	}
	mean := sum / float64(totalCount)
	min := math.MaxFloat64
	var max float64
	var midLo, midHi int
	if totalCount%2 == 0 {
		midLo = totalCount / 2
		midHi = totalCount/2 + 1
	} else {
		midLo = totalCount/2 + 1
		midHi = totalCount/2 + 1
	}
	var medianSum float64
	var countSoFar int
	for i, c := range count {
		if c == 0 {
			continue
		}
		min = math.Min(min, float64(i))
		max = math.Max(max, float64(i))
		next := countSoFar + c
		if countSoFar < midLo && midLo <= next {
			medianSum += float64(i)
		}
		if countSoFar < midHi && midHi <= next {
			medianSum += float64(i)
		}
		countSoFar = next
	}
	median := medianSum / 2

	return []float64{min, max, mean, median, mode}
}
