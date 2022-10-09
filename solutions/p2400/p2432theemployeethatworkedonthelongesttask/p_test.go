package p2432theemployeethatworkedonthelongesttask

import (
	"math"
)

func hardestWorker(n int, logs [][]int) int {
	var t int
	var maxD int
	res := math.MaxInt32
	for _, l := range logs {
		d := l[1] - t
		if d > maxD || (d == maxD && l[0] < res) {
			maxD = d
			res = l[0]
		}
		t = l[1]
	}
	return res
}
