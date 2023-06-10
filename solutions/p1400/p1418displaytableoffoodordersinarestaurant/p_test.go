package p1418displaytableoffoodordersinarestaurant

import (
	"fmt"
	"sort"
)

func displayTable(orders [][]string) [][]string {
	foods := make(map[string]struct{})
	for _, x := range orders {
		foods[x[2]] = struct{}{}
	}
	tables := make(map[string]map[string]int)
	for i := range orders {
		tableNumber := orders[i][1]
		foodItem := orders[i][2]
		if _, exists := tables[tableNumber]; !exists {
			tables[tableNumber] = make(map[string]int)
		}
		tables[tableNumber][foodItem]++
	}
	header := make([]string, 0, len(foods)+1)
	for f := range foods {
		header = append(header, f)
	}
	header = append(header, "")
	sort.Strings(header)
	header[0] = "Table"
	var res [][]string
	res = append(res, header)
	for x := 1; x <= 500; x++ {
		v, exists := tables[fmt.Sprint(x)]
		if !exists {
			continue
		}
		ss := []string{fmt.Sprint(x)}
		for j := 1; j < len(res[0]); j++ {
			food := res[0][j]
			ss = append(ss, fmt.Sprint(v[food]))
		}
		res = append(res, ss)
	}
	return res
}
