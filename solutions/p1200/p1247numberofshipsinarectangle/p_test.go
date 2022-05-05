package p1247numberofshipsinarectangle

type Sea interface {
	hasShips(topRight, bottomLeft []int) bool
}

func countShips(sea Sea, topRight, bottomLeft []int) int {
	if topRight[0] < bottomLeft[0] || topRight[1] < bottomLeft[1] || !sea.hasShips(topRight, bottomLeft) {
		return 0
	}

	dx := topRight[0] - bottomLeft[0]
	dy := topRight[1] - bottomLeft[1]
	if dx == 0 && dy == 0 {
		return 1
	}

	xmid := bottomLeft[0] + (dx / 2)
	ymid := bottomLeft[1] + (dy / 2)
	// Split all four
	return countShips(sea, topRight, []int{xmid + 1, ymid + 1}) +
		countShips(sea, []int{xmid, topRight[1]}, []int{bottomLeft[0], ymid + 1}) +
		countShips(sea, []int{topRight[0], ymid}, []int{xmid + 1, bottomLeft[1]}) +
		countShips(sea, []int{xmid, ymid}, bottomLeft)
}
