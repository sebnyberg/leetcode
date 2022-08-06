package p0771jewelsandstones

func numJewelsInStones(jewels string, stones string) int {
	var isJewel [256]bool
	for _, ch := range jewels {
		isJewel[ch] = true
	}
	var count int
	for _, s := range stones {
		if isJewel[s] {
			count++
		}
	}
	return count
}
