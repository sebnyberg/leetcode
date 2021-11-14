package p2075decodetheslantedciphertext

import (
	"fmt"
	"strings"
	"testing"
	"unicode"

	"github.com/stretchr/testify/require"
)

func Test_decodeCiphertext(t *testing.T) {
	for _, tc := range []struct {
		encodedText string
		rows        int
		want        string
	}{
		{"ch   ie   pr", 3, "cipher"},
		{"iveo    eed   l te   olc", 4, "i love leetcode"},
		{"coding", 1, "coding"},
		{" b  ac", 2, " abc"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.encodedText), func(t *testing.T) {
			require.Equal(t, tc.want, decodeCiphertext(tc.encodedText, tc.rows))
		})
	}
}

func decodeCiphertext(encodedText string, rows int) string {
	encodedRows := make([]string, rows)
	rowLen := len(encodedText) / rows
	for i := 0; i < rows; i++ {
		encodedRows[i] = encodedText[i*rowLen : (i+1)*rowLen]
	}
	result := make([]byte, 0)
	for startJ := 0; ; startJ++ {
		j := startJ
		for i := 0; i < rows; i, j = i+1, j+1 {
			if j >= rowLen {
				return strings.TrimRightFunc(string(result), unicode.IsSpace)
			}
			result = append(result, encodedRows[i][j])
		}
	}
}
