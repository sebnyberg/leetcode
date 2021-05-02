package p0299bullsandcows

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_bullsAndCows(t *testing.T) {
	for _, tc := range []struct {
		secret string
		guess  string
		want   string
	}{
		{"00112233445566778899", "16872590340158679432", "3A17B"},
		{"1807", "7810", "1A3B"},
		{"1123", "0111", "1A1B"},
		{"1", "0", "0A0B"},
		{"1", "1", "1A0B"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.secret), func(t *testing.T) {
			require.Equal(t, tc.want, getHint(tc.secret, tc.guess))
		})
	}
}

func getHint(secret string, guess string) string {
	var secrets [10]int
	var guessed [10]int
	var matched int
	for i := range guess {
		if guess[i] == secret[i] {
			matched++
			continue
		}
		secrets[secret[i]-'0']++
		guessed[guess[i]-'0']++
	}

	var guessedWrongPos int
	for ch, count := range secrets {
		if count > 0 {
			guessedWrongPos += min(count, guessed[ch])
		}
	}
	return strconv.Itoa(matched) + "A" + strconv.Itoa(guessedWrongPos) + "B"
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
