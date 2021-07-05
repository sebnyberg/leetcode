package leetcode

func findLength(nums1 []int, nums2 []int) int {
	// Values can be from 0-100 => pick 113 as the base prime
	if len(nums2) < len(nums1) {
		nums1, nums2 = nums2, nums1
	}
	l, r := 0, len(nums1)+1
	base := 113
	mod := 1_000_000_007
	// modInv := modInverse(mod, base)
	for l < r {
		mid := l + (r-l)/2
		if check(nums1, nums2, base, mod, mid) {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l - 1
}

// func Test_check(t *testing.T) {
// 	check([]int{1, 2, 1, 2, 1, 2}, []int{2, 1}, 113, 1_000_000_007, 2)
// }

func check(nums1, nums2 []int, base, mod, n int) bool {
	// Find all possible hashes of length n within nums1
	possibleHashes := make(map[int]struct{})
	var h int
	power := 1
	for i, num := range nums1 {
		h = (h*base + num) % mod // add number
		if i < n {
			power = power * base % mod
		}
		if i >= n {
			// remove old element
			h = (h - nums1[i-n]*power%mod + mod) % mod
		}
		if i >= n-1 {
			possibleHashes[h] = struct{}{}
		}
	}

	// Now find all possible hashes in nums2
	// If there is a match with existing hashes, exit with true
	h = 0
	power = 1
	for i, num := range nums2 {
		h = (h*base + num) % mod // add number
		if i < n {
			power = power * base % mod
		}
		if i >= n {
			// remove old element
			h = (h - nums2[i-n]*power%mod + mod) % mod
		}
		if i >= n-1 {
			if _, exists := possibleHashes[h]; exists {
				return true
			}
		}
	}
	return false
}

func modInverse(a, mod int) int {
	return modPow(a, mod-2, mod)
}

func modPow(a, b, mod int) int {
	if b == 0 {
		return 1
	}
	p := modPow(a, b/2, mod) % mod
	p = p * p % mod
	if b%2 == 0 {
		return p
	}
	return (a * p) % mod
}
