package p1357applydiscounteverynorders

type Cashier struct {
	price          map[int]int
	discountFactor float64
	k              int
	n              int
}

func Constructor(n int, discount int, products []int, prices []int) Cashier {
	price := make(map[int]int)
	for i := range products {
		price[products[i]] = prices[i]
	}
	return Cashier{
		k:              0,
		n:              n,
		discountFactor: float64(100-discount) / 100,
		price:          price,
	}
}

func (this *Cashier) GetBill(product []int, amount []int) float64 {
	var sum int
	for i := range product {
		sum += this.price[product[i]] * amount[i]
	}
	this.k = (this.k + 1) % this.n
	res := float64(sum)
	if this.k == 0 {
		res *= this.discountFactor
	}
	return res
}
