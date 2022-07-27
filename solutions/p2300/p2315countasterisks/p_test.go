package p2315countasterisks

import "strings"

func countAsterisks(s string) int {
	parts := strings.Split(s, "|")
	var count int
	for i := 0; i < len(parts); i += 2 {
		for _, ch := range parts[i] {
			if ch == '*' {
				count++
			}
		}
	}
	return count
}
