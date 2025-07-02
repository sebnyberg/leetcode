package p3330findtheoriginaltypedstringi

func possibleStringCount(word string) int {
	res := 1
	count := 1
	for i := 1; i < len(word); i++ {
		if word[i] != word[i-1] {
			res += count - 1
			count = 0
		}
		count++
	}
	res += count - 1
	return res
}
