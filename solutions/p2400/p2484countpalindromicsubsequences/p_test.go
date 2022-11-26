package p2484countpalindromicsubsequences

func countPalindromes(s string) int {
	if len(s) < 5 {
		return 0
	}
	n := len(s)
	const mod = 1e9 + 7
	left := make([][100]int, n+1)
	left1 := make([][10]int, n+1)
	left1[1][s[0]-'0']++
	for i := 2; i <= len(s); i++ {
		y := int(s[i-1] - '0')
		left[i] = left[i-1]
		for x, c := range left1[i-1] {
			a := x*10 + y
			left[i][a] = (left[i][a] + c) % mod
		}
		left1[i] = left1[i-1]
		left1[i][y]++
	}
	var right1 [10]int
	var right [100]int
	right1[s[n-1]-'0']++
	var res int
	rev := func(x int) int {
		var res int
		res += x % 10
		res *= 10
		x /= 10
		res += x
		return res
	}
	for i := len(s) - 2; i >= 2; i-- {
		for x, c := range right {
			want := rev(x)
			tot := c * left[i][want]
			res = (res + tot) % mod
		}
		y := int(s[i] - '0')
		for x, c := range right1 {
			a := y*10 + x
			right[a] = (right[a] + c) % mod
		}
		right1[s[i]-'0']++
	}
	return res
}
