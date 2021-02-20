package a

func romanToInt(s string) (res int) {
	var i int
	n := len(s)
	for i < n {
		ch := rune(s[i])
		switch ch {
		case 'M':
			res += 1000
		case 'D':
			res += 500
		case 'C':
			if i < n-1 {
				if s[i+1] == 'D' {
					res += 400
					i++
					break
				} else if s[i+1] == 'M' {
					res += 900
					i++
					break
				}
			}
			res += 100
		case 'L':
			res += 50
		case 'X':
			if i < n-1 {
				if s[i+1] == 'L' {
					res += 40
					i++
					break
				} else if s[i+1] == 'C' {
					res += 90
					i++
					break
				}
			}
			res += 10
		case 'V':
			res += 5
		case 'I':
			if i < n-1 {
				if s[i+1] == 'V' {
					res += 4
					i++
					break
				} else if s[i+1] == 'X' {
					res += 9
					i++
					break
				}
			}
			res += 1
		}
		i++
	}
	return res
}
