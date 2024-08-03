package p1460maketwoarraysequalbyreversingsubarrays

func canBeEqual(target []int, arr []int) bool {
	// Count the frequency of each number in both arrays then return true if they
	// are equal.
	var count [2][1001]int
	for i := range target {
		count[0][target[i]]++
		count[1][arr[i]]++
	}
	return count[0] == count[1]
}
