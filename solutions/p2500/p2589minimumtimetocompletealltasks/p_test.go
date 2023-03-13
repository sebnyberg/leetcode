package p2589minimumtimetocompletealltasks

import "sort"

func findMinimumTime(tasks [][]int) int {
	var busy [2001]bool
	sort.Slice(tasks, func(i, j int) bool {
		a := tasks[i]
		b := tasks[j]
		return a[1] < b[1]
	})
	var count int
	for _, t := range tasks {
		start := t[0]
		end := t[1]
		d := t[2]
		for i := start; i <= end; i++ {
			if busy[i] {
				d--
			}
		}
		if d <= 0 {
			continue
		}
		for i := end; i >= start; i-- {
			if !busy[i] {
				d--
				busy[i] = true
				count++
			}
			if d <= 0 {
				break
			}
		}
	}
	return count
}
