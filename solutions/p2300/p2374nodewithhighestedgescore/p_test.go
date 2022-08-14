package p2374nodewithhighestedgescore

func edgeScore(edges []int) int {
	n := len(edges)
	scores := make([]int, n)
	var maxScore int
	var maxNode int
	for i, e := range edges {
		scores[e] += i
		if scores[e] > maxScore ||
			(scores[e] == maxScore && e < maxNode) {
			maxScore = scores[e]
			maxNode = e
		}
	}
	return maxNode
}
