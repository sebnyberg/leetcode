package p0946validatestacksequences

func validateStackSequences(pushed []int, popped []int) bool {
	stack := []int{}
	var j int
	for _, x := range pushed {
		stack = append(stack, x)
		for j < len(popped) && len(stack) > 0 && popped[j] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			j++
		}
	}
	return j == len(popped)
}
