package p1169invalidtransactions

import (
	"sort"
	"strconv"
	"strings"
)

func invalidTransactions(transactions []string) []string {
	num := func(s string) int {
		a, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		return a
	}

	type transaction struct {
		name string
		time int
		amt  int
		city string
		s    string
	}

	parse := func(s string) transaction {
		parts := strings.Split(s, ",")
		name := parts[0]
		ts := num(parts[1])
		amount := num(parts[2])
		city := parts[3]
		return transaction{name, ts, amount, city, s}
	}

	// Partition by name
	byName := make(map[string][]transaction)
	for _, s := range transactions {
		t := parse(s)
		byName[t.name] = append(byName[t.name], t)
	}

	// Sort by time
	for _, v := range byName {
		sort.Slice(v, func(i, j int) bool {
			return v[i].time < v[j].time
		})
	}

	// For each list of transactions (ordered by time, ascending)
	var res []string
	var invalid []bool
	for _, v := range byName {
		// Reset. Not sure if this is the most efficient way.
		invalid = invalid[:0]
		invalid = append(invalid, make([]bool, len(v))...)

		// This is sad. But I can't find another simple, reliable way to
		// validate each item than to exhaustively try other items and there are
		// only 1000 elements, so let's do it.
		//
		// There are ways to make this faster, e.g. sliding window, but it makes
		// code more complicated and time complexity allows for this dumb
		// solution.
		for i := range v {
			if v[i].amt > 1000 {
				invalid[i] = true
				continue
			}
			for j := 0; j < len(v); j++ {
				if abs(v[j].time-v[i].time) <= 60 && v[j].city != v[i].city {
					invalid[i] = true
					break
				}
			}
		}

		for i := range invalid {
			if invalid[i] {
				res = append(res, v[i].s)
			}
		}
	}

	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
