package p0925longpressedname

func isLongPressedName(name string, typed string) bool {
	var j int
	for i := range name {
		if j >= len(typed) || name[i] != typed[j] {
			return false
		}
		if i == len(name)-1 || name[i+1] != name[i] {
			// read extra characters
			for j < len(typed)-1 && typed[j] == typed[j+1] {
				j++
			}
		}
		j++
	}
	return j == len(typed)
}
