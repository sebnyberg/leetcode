package p1404numberofsteptstoreduceanumberinbinaryrepresentatinontoone

func numSteps(s string) int {
	var carry int
	var ops int
	for i := len(s) - 1; i >= 0; i-- {
		if i == 0 && carry == 0 {
			break
		}
		d := int(s[i]-'0') + carry
		if d&1 == 1 {
			ops++
			d++
		}
		carry = d / 2
		ops++
	}
	return ops
}
