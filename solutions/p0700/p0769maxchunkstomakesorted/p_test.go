package p0769maxchunkstomakesorted

func maxChunksToSorted(arr []int) int {
	// A valid chunk at an index position contains all values up to and
	// and including that index (as a value).
	// I.e., to chunk at 0, the subarray up to and including 0 must include
	// the zero. To chunk at 2, we need [0,1,2] within the subarray.
	// Chunking as soon as possible is optimal.

	n := len(arr)
	seen := make([]bool, n)
	var j int
	var res int
	for i, v := range arr {
		seen[v] = true
		for j < n && seen[j] {
			j++
		}
		if j == i+1 {
			res++
		}
	}
	return res
}
