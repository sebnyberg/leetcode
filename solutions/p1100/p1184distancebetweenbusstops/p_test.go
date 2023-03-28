package p1184distancebetweenbusstops

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	n := len(distance)
	if start > destination {
		start, destination = destination, start
	}
	var fwd int
	for i := start; i < destination; i++ {
		fwd += distance[i]
	}
	var back int
	for i := destination; i != start; i = (i + 1) % n {
		back += distance[i]
	}
	return min(fwd, back)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
