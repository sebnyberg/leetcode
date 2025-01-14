package p1769minimunnumberofoperationstomoveballstoeachbox

func minOperations(boxes string) []int {
	n := len(boxes)
	right := make([]int, n+1)
	var count int
	for i := n - 1; i >= 0; i-- {
		right[i] = right[i+1] + count
		count += int(boxes[i] - '0')
	}
	count = 0
	var left int
	res := make([]int, n)
	for i, x := range boxes {
		res[i] = left + right[i]
		count += int(x - '0')
		left += count
	}
	return res
}
