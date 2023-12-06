package p1716calculatemoneyinleetcodebank

func totalMoney(n int) int {
	var res int
	var off int
	for i := 0; i < n; i++ {
		res += off + ((i % 7) + 1)
		if i%7 == 6 {
			off++
		}
	}
	return res
}
