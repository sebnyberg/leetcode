package p3461checkifdigitsareequalinstringafteroperationsi

func hasSameDigits(s string) bool {
	digits := make([]int, len(s))
	for i := range s {
		digits[i] = int(s[i] - '0')
	}
	for len(digits) > 2 {
		for i := 0; i < len(digits)-1; i++ {
			digits[i] = (digits[i] + digits[i+1]) % 10
		}
		digits = digits[:len(digits)-1]
	}
	return digits[0] == digits[1]
}
