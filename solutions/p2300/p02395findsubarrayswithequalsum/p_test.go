package p02395findsubarrayswithequalsum

func findSubarrays(nums []int) bool {
	seen := make(map[int]struct{})
	for i := 1; i < len(nums)-1; i++ {
		x := nums[i-1] + nums[i]
		if _, exists := seen[x]; exists {
			continue
		}
		seen[x] = struct{}{}
		for j := i; j < len(nums)-1; j++ {
			if nums[j]+nums[j+1] == x {
				return true
			}
		}
	}
	return false
}
