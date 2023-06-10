package p1436destinationcity

func destCity(paths [][]string) string {
	outdeg := make(map[string]int)
	for _, p := range paths {
		a := p[0]
		b := p[1]
		if _, exists := outdeg[a]; !exists {
			outdeg[a] = 0
		}
		if _, exists := outdeg[b]; !exists {
			outdeg[b] = 0
		}
		outdeg[a]++
	}
	for k, deg := range outdeg {
		if deg == 0 {
			return k
		}
	}
	return "NOT GOOD!!"
}
