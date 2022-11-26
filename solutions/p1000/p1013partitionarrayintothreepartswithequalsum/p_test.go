package p1013partitionarrayintothreepartswithequalsum

func canThreePartsEqualSum(arr []int) bool {
	n := len(arr)
	pre := make([]int, n+1)
	for i := range arr {
		pre[i+1] = pre[i] + arr[i]
	}
	for i := 1; i <= n-2; i++ {
		s1 := pre[i]
		for j := i + 1; j <= n-1; j++ {
			s2 := pre[j] - s1
			s3 := pre[n] - s1 - s2
			if s1 == s2 && s2 == s3 {
				return true
			}
		}
	}
	return false
}
