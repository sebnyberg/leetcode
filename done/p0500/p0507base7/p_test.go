package p0507base7

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	res := make([]byte, 0)
	sign := 1
	if num < 0 {
		sign = -1
		num = -num
	}
	for num > 0 {
		res = append(res, byte(num%7)+'0')
		num /= 7
	}
	if sign == -1 {
		res = append(res, '-')
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}
