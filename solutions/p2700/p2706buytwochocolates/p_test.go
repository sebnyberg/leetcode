package p2706buytwochocolates

func buyChoco(prices []int, money int) int {
    sort.Ints(prices)
    if prices[0] + prices[1] > money {
        return money
    }
    return money - prices[0] - prices[1]
}
