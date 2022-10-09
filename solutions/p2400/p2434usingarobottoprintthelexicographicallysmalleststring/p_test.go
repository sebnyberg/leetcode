package p2434usingarobottoprintthelexicographicallysmalleststring

func robotWithString(s string) string {
	// Let's say that there is a single 'a' at the end of s.
	// Then we should remove all letters until that 'a', then ask the robot to
	// print it.
	//
	// What this means is that IF there exists a character in 's' which is
	// smaller than the final letter in 't', then we should keep adding new
	// letters in 't' until this is no longer the case.
	//
	// So the problem is mainly about keeping track of whether a certain letter
	// exists at or after a certain point in 's'.
	//
	// We can do this using a bitwise or. There is probably some other clever
	// way that is O(1) but I think this will be fast enough anyway.
	//
	n := len(s)
	hasLetter := make([]int, n)
	stack := make([]rune, 0, n)
	hasLetter[n-1] = (1 << int(s[n-1]-'a'))
	for i := n - 2; i >= 0; i-- {
		hasLetter[i] = hasLetter[i+1] | (1 << int(s[i]-'a'))
	}
	res := make([]rune, 0, n)
	for i := 0; i < n; i++ {
		for len(stack) > 0 {
			for ch := 'a'; ch < stack[len(stack)-1]; ch++ {
				if hasLetter[i]&(1<<(ch-'a')) > 0 {
					goto c
				}
			}
			res = append(res, stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
	c:
		stack = append(stack, rune(s[i]))
	}
	for len(stack) > 0 {
		res = append(res, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	return string(res)
}
