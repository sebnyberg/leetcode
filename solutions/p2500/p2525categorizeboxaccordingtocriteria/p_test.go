package p2525categorizeboxaccordingtocriteria

func categorizeBox(length int, width int, height int, mass int) string {
	bulky := length >= 10000 || height >= 10000 || width >= 10000 ||
		length*width*height >= 1000000000
	heavy := mass >= 100
	if bulky && heavy {
		return "Both"
	}
	if !bulky && !heavy {
		return "Neither"
	}
	if bulky {
		return "Bulky"
	}
	return "Heavy"
}
