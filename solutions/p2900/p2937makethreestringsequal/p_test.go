package p2937makethreestringsequal

import "sort"

func findMinimumOperations(s1 string, s2 string, s3 string) int {
	ss := []string{s1, s2, s3}
	sort.Slice(ss, func(i, j int) bool {
		return len(ss[i]) < len(ss[j])
	})
	s1 = ss[0]
	s2 = ss[1]
	s3 = ss[2]
	var res int
	for len(s2) != len(s1) {
		res++
		s2 = s2[:len(s2)-1]
	}
	for len(s3) != len(s1) {
		res++
		s3 = s3[:len(s3)-1]
	}
	for len(s1) > 1 && (s1 != s2 || s2 != s3) {
		res += 3
		s1 = s1[:len(s1)-1]
		s2 = s2[:len(s2)-1]
		s3 = s3[:len(s3)-1]
	}
	if len(s1) <= 1 {
		return -1
	}
	return res
}
