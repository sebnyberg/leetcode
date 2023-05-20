package p1287elementappearingmorethan25percentinsortedarray

func findSpecialInteger(arr []int) int {
	freq := make(map[int]int)
	n := len(arr)
	for _, x := range arr {
		freq[x]++
		if freq[x] > n/4 {
			return x
		}
	}
	return -1
}
