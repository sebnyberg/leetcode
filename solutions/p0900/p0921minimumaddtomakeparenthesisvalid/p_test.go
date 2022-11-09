package p0921minimumaddtomakeparenthesisvalid

func minAddToMakeValid(s string) int {
	var count int
	var open int
	for _, ch := range s {
		if ch == ')' {
			if open == 0 {
				count++
			} else {
				open--
			}
		} else {
			open++
		}
	}
	return count + open
}
