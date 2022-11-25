package p1003checkifawordisvalidaftersubstitutions

func isValid(s string) bool {
	stack := []byte{}
	for _, ch := range s {
		stack = append(stack, byte(ch))
		for len(stack) >= 3 && string(stack[len(stack)-3:]) == "abc" {
			stack = stack[:len(stack)-3]
		}
	}
	return len(stack) == 0
}
