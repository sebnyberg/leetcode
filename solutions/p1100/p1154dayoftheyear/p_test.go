package p1154dayoftheyear

import (
	"fmt"
	"time"
)

func dayOfYear(date string) int {
	var d [3]int
	fmt.Sscanf(date, "%04d-%02d-%02d", &d[0], &d[1], &d[2])
	t := time.Date(d[0], time.Month(d[1]), d[2], 0, 0, 0, 0, time.UTC)
	return t.YearDay()
}
