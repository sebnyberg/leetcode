package p1604alertusingsamekeycarthreeormoretimesinaonehourperiod

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_alertNames(t *testing.T) {
	for _, tc := range []struct {
		keyName []string
		keyTime []string
		want    []string
	}{
		{[]string{"daniel", "daniel", "daniel", "luis", "luis", "luis", "luis"}, []string{"10:00", "10:40", "11:00", "09:00", "11:00", "13:00", "15:00"}, []string{"daniel"}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.keyName), func(t *testing.T) {
			require.Equal(t, tc.want, alertNames(tc.keyName, tc.keyTime))
		})
	}
}

func alertNames(keyName []string, keyTime []string) []string {
	times := make(map[string][]int)
	parseTime := func(timeStr string) int {
		var h, m int
		fmt.Sscanf(timeStr, "%02d:%02d", &h, &m)
		return 60*h + m
	}

	for i, name := range keyName {
		t := parseTime(keyTime[i])
		times[name] = append(times[name], t)
	}
	var res []string
	for n, tt := range times {
		sort.Ints(tt)
		for i := 2; i < len(tt); i++ {
			if tt[i]-tt[i-2] <= 60 {
				res = append(res, n)
				break
			}
		}
	}
	sort.Strings(res)
	return res
}
