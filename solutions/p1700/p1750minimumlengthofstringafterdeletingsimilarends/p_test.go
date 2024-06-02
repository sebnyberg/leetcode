package p1750minimumlengthofstringafterdeletingsimilarends

func minimumLength(s string) int {
	// I guess this is just a greedy problem?
	// Let's start with RLE
	type charCount struct {
			char byte
			count int
	}
	rle := []charCount{charCount{s[0], 1}}
	var j int
	for i := 1; i < len(s); i++ {
			if rle[j].char != s[i] {
					rle = append(rle, charCount{s[i], 1})
					j++
			} else {
					rle[j].count++
			}
	}
	for l, r := 0, j; l <= r; l, r = l+1, r-1 {
			if rle[l].char == rle[r].char {
					if l != r || rle[l].count > 1 {
							rle[l].count = 0
							rle[r].count = 0
					}
					continue
			}
			break
	}
	var res int
	for _, x := range rle {
			res += x.count
	}
	return res
}