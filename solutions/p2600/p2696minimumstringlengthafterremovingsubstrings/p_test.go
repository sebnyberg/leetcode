package p2696minimumstringlengthafterremovingsubstrings

func minLength(s string) int {
	var stack []byte
	for _, ch := range s {
		stack = append(stack, byte(ch))
		for len(stack) >= 2 {
			end := string(stack[len(stack)-2:])
			if end == "AB" || end == "CD" {
				stack = stack[:len(stack)-2]
			} else {
				break
			}
		}
	}
	return len(stack)
}
