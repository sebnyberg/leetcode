package p1524numberofsubarrayswithoddsum

const mod = 1e9 + 7

func numOfSubarrays(arr []int) int {
	var oddCount int
	var evenCount int
	var res int
	for _, x := range arr {
		if x&1 == 1 {
			oddCount, evenCount = evenCount+1, oddCount
		} else {
			oddCount, evenCount = oddCount, evenCount+1
		}
		res = (res + oddCount) % mod
	}
	return res
}
