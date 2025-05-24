package p2942findwordscontainingcharacter

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindWords(t *testing.T) {
	tests := []struct {
		words    []string
		char     byte
		expected []int
	}{
		{
			words:    []string{"leet", "code"},
			char:     'e',
			expected: []int{0, 1},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			result := findWordsContaining(tt.words, tt.char)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("findWords() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func findWordsContaining(words []string, x byte) []int {
	var indices []int
	for i, w := range words {
		for j := range w {
			if w[j] == x {
				indices = append(indices, i)
				break
			}
		}
	}
	return indices
}
