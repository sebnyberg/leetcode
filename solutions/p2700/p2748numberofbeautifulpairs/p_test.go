package p2748numberofbeautifulpairs

import "fmt"

func countBeautifulPairs(nums []int) int {
	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}

	var res int
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			a := fmt.Sprint(nums[i])
			b := fmt.Sprint(nums[j])
			if gcd(int(a[0]-'0'), int(b[len(b)-1]-'0')) == 1 {
				res++
			}
		}
	}
	return res
}
