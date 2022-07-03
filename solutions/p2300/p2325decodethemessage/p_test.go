package p2325decodethemessage

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_decodeMessage(t *testing.T) {
	for _, tc := range []struct {
		key     string
		message string
		want    string
	}{
		{"the quick brown fox jumps over the lazy dog", "vkbs bs t suepuv", "this is a secret"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.key), func(t *testing.T) {
			require.Equal(t, tc.want, decodeMessage(tc.key, tc.message))
		})
	}
}
func decodeMessage(key string, message string) string {
	var charMap [26]byte
	var seen [26]bool
	var j int
	for _, ch := range key {
		if ch == ' ' || seen[ch-'a'] {
			continue
		}
		seen[ch-'a'] = true
		charMap[ch-'a'] = byte(j + 'a')
		j++
	}
	res := make([]byte, len(message))
	for i := range res {
		if message[i] == ' ' {
			res[i] = message[i]
			continue
		}
		res[i] = charMap[message[i]-'a']
	}
	return string(res)
}
