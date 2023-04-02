package p2611miceandcheese

import "sort"

func miceAndCheese(reward1 []int, reward2 []int, k int) int {
	// Imagine that the second mouse eats all cheese
	var sum int
	for _, x := range reward2 {
		sum += x
	}

	// Imagine the delta if the second mouse hands over some of the cheese to
	// the first mouse
	n := len(reward1)
	delta := make([]int, n)
	for i := range reward1 {
		delta[i] = reward1[i] - reward2[i]
	}

	// Maximize the delta
	sort.Ints(delta)
	for i := 0; i < k; i++ {
		sum += delta[len(delta)-i-1]
	}
	return sum
}
