package p1796secondlargestdigitinastring

func secondHighest(s string) int {
	var digits [10]bool
	for _, r := range s {
		if r >= '0' && r <= '9' {
			digits[r-'0'] = true
		}
	}
	count := 0
	for i := 9; i >= 0; i-- {
		if digits[i] {
			count++
			if count == 2 {
				return i
			}
		}
	}
	return -1
}
