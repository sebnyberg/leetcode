package p0394decodestring

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/stretchr/testify/require"
)

func Test_decodeString(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want string
	}{
		{"3[a]2[bc]", "aaabcbc"},
		{"3[a2[c]]", "accaccacc"},
		{"2[abc]3[cd]ef", "abcabccdcdcdef"},
		{"abc3[cd]xyz", "abccdcdcdxyz"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, decodeString(tc.s))
		})
	}
}

type EncodedString struct {
	s     string
	n     int
	pos   int
	start int
}

var eof = rune(0)

func (s *EncodedString) peek() rune {
	if s.pos == s.n {
		return eof
	}
	ch, _ := utf8.DecodeRuneInString(s.s[s.pos:])
	return ch
}

func (s *EncodedString) next() rune {
	if s.pos == s.n {
		return eof
	}
	ch, width := utf8.DecodeRuneInString(s.s[s.pos:])
	s.pos += width
	return ch
}

func (s *EncodedString) accept(alphabet string) rune {
	ch := s.peek()
	if strings.ContainsRune(alphabet, ch) {
		s.next()
	} else {
		return eof
	}
	s.start = s.pos
	return ch
}

func (s *EncodedString) acceptRun(alphabet string) string {
	for strings.ContainsRune(alphabet, s.peek()) {
		s.next()
	}
	res := s.s[s.start:s.pos]
	s.start = s.pos
	return res
}

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func decodeString(s string) string {
	encoded := &EncodedString{
		s: s,
		n: len(s),
	}
	return decode(encoded)
}

func decode(encoded *EncodedString) string {
	result := ""
	alpha := "abcdefghijklmnopqrstuvwxyz"
	num := "0123456789"
	for {
		switch ch := encoded.peek(); {
		case strings.ContainsRune(alpha, ch):
			result += encoded.acceptRun(alpha)
		case strings.ContainsRune(num, ch):
			numStr := encoded.acceptRun(num)
			n, err := strconv.Atoi(numStr)
			check(err)
			encoded.accept("[")
			inner := decode(encoded)
			encoded.accept("]")
			for i := 0; i < n; i++ {
				result += inner
			}
		default: // eof, "[" or "]"
			return result
		}
	}
}
