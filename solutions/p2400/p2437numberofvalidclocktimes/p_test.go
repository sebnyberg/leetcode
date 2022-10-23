package p2437numberofvalidclocktimes

import "fmt"

func countTime(time string) int {
	buf := make([]byte, 5)
	copy(buf, []byte(time))
	res := dfs(time, buf, 0)
	return res
}

func dfs(t string, buf []byte, i int) int {
	if i == 5 {
		if valid(string(buf)) {
			return 1
		}
		return 0
	}
	if i == 2 {
		i++
	}
	var res int
	if t[i] == '?' {
		for x := byte('0'); x <= '9'; x++ {
			buf[i] = x
			res += dfs(t, buf, i+1)
		}
	} else {
		res += dfs(t, buf, i+1)
	}
	return res
}

func valid(s string) bool {
	var hh, mm int
	fmt.Sscanf(s, "%d:%d", &hh, &mm)
	return hh <= 23 && mm <= 59
}
