package p2598smallestmissingnonnegativeintegerafteroperations

func findSmallestInteger(nums []int, value int) int {
	// count numbers per modulo
	// note that Go's modulo returns negative values at times
	count := make(map[int]int)
	for _, x := range nums {
		count[mod(x, value)]++
	}
	for i := 0; ; i++ {
		if count[i%value] == 0 {
			return i
		}
		count[i%value]--
	}
}

func mod(a, b int) int {
	// Can't remember why this works but it does
	return (b + (a % b)) % b
}
