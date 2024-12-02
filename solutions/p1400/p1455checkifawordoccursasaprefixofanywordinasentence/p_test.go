package p1455checkifawordoccursasaprefixofanywordinasentence

import "strings"

func isPrefixOfWord(sentence string, searchWord string) int {
	for i, w := range strings.Fields(sentence) {
		if len(w) >= len(searchWord) && w[:len(searchWord)] == searchWord {
			return i + 1
		}
	}
	return -1
}
