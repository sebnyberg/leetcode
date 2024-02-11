package p2933highaccessemployees

import "sort"

func findHighAccessEmployees(access_times [][]string) []string {
	minutes := func(t string) int {
		h := int(t[0]-'0')*10 + int(t[1]-'0')
		m := int(t[2]-'0')*10 + int(t[3]-'0')
		return h*60 + m
	}

	sort.Slice(access_times, func(i, j int) bool {
		return access_times[i][1] < access_times[j][1]
	})
	m := make(map[string][]int)
	for i := range access_times {
		name := access_times[i][0]
		t := access_times[i][1]
		m[name] = append(m[name], minutes(t))
	}
	var res []string
	for k, ts := range m {
		for i := 2; i < len(ts); i++ {
			if ts[i]-ts[i-2] < 60 {
				res = append(res, k)
				break
			}
		}
	}
	return res
}
