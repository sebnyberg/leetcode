package p1475finalpriceswithaspecialdiscountinashop

func finalPrices(prices []int) []int {
	var res []int
	for i := range prices {
		var discount int
		for j := i + 1; j < len(prices); j++ {
			if prices[j] <= prices[i] {
				discount = prices[j]
				break
			}
		}
		res = append(res, prices[i]-discount)
	}
	return res
}
