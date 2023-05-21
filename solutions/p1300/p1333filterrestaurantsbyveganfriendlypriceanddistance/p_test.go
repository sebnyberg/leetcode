package p1333filterrestaurantsbyveganfriendlypriceanddistance

import "sort"

func filterRestaurants(restaurants [][]int, veganFriendly int, maxPrice int, maxDistance int) []int {
	// Filter
	var j int
	for _, r := range restaurants {
		if r[2] == 0 && veganFriendly == 1 {
			continue
		}
		if r[3] > maxPrice {
			continue
		}
		if r[4] > maxDistance {
			continue
		}
		restaurants[j] = r
		j++
	}
	restaurants = restaurants[:j]

	// Sort
	sort.Slice(restaurants, func(i, j int) bool {
		a := restaurants[i]
		b := restaurants[j]
		if a[1] == b[1] {
			return a[0] > b[0]
		}
		return a[1] > b[1]
	})
	res := make([]int, j)
	for i := range res {
		res[i] = restaurants[i][0]
	}
	return res
}
