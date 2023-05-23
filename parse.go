package leetcode

import (
	"log"
	"strconv"
	"strings"
)

func ParseMatrix(s string) [][]int {
	s = s[2 : len(s)-2]
	if s == "" {
		return nil
	}
	parts := strings.Split(s, "],[")
	res := make([][]int, len(parts))
	for i, part := range parts {
		if part == "" {
			continue
		}
		for _, numStr := range strings.Split(part, ",") {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				log.Fatalf("failed to parse number, %v, %v\n", numStr, err)
			}
			res[i] = append(res[i], num)
		}
	}
	return res
}
