package p2999countthenumberofpowerfulintegers

import (
	"fmt"
	"strconv"
)

func numberOfPowerfulInt(start int64, finish int64, limit int, s string) int64 {
	// start with the limit
	str := func(x int) string {
		return fmt.Sprint(x)
	}

	toInt := func(s string) int {
		x, _ := strconv.Atoi(s)
	}
}
