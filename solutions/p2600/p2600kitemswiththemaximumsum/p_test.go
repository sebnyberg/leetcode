package p2600kitemswiththemaximumsum

func kItemsWithMaximumSum(numOnes int, numZeros int, numNegOnes int, k int) int {
	if k > numOnes+numZeros {
		return numOnes - (k - numOnes - numZeros)
	}
	return min(k, numOnes)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
