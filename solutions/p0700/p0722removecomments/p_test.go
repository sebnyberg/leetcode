package p0722removecomments

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_removeComments(t *testing.T) {
	for _, tc := range []struct {
		source []string
		want   []string
	}{
		{
			[]string{"/*Test program */", "int main()", "{ ", "  // variable declaration ", "int a, b, c;", "/* This is a test", "   multiline  ", "   comment for ", "   testing */", "a = b + c;", "}"},
			[]string{"int main()", "{ ", "  ", "int a, b, c;", "a = b + c;", "}"},
		},
		{
			[]string{"a/*comment", "line", "more_comment*/b"}, []string{"ab"},
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.source), func(t *testing.T) {
			require.Equal(t, tc.want, removeComments(tc.source))
		})
	}
}

type stateFn func(i *input) stateFn

type input struct {
	s   string
	pos int
	buf []byte
	res []string
}

func removeComments(source []string) []string {
	i := &input{
		s:   strings.Join(source, "\n") + "\n",
		pos: 0,
		buf: make([]byte, 0, len(source)),
		res: make([]string, 0, len(source)),
	}
	for s := parseText(i); s != nil; {
		s = s(i)
	}
	return i.res
}

func parseText(i *input) stateFn {
	for {
		if i.pos == len(i.s) {
			return nil
		}
		if i.pos < len(i.s)-1 {
			a := i.s[i.pos : i.pos+2]
			if a == "/*" {
				i.pos += 2
				return parseMulti(i)
			} else if a == "//" {
				i.pos += 2
				return parseSingle(i)
			}
		}
		if i.s[i.pos] == '\n' {
			if len(i.buf) > 0 {
				i.res = append(i.res, string(i.buf))
			}
			i.buf = i.buf[:0]
		} else {
			i.buf = append(i.buf, i.s[i.pos])
		}
		i.pos++
	}
}

func parseMulti(i *input) stateFn {
	for i.pos < len(i.s)-1 {
		if i.s[i.pos:i.pos+2] == "*/" {
			i.pos += 2
			return parseText(i)
		}
		i.pos++
	}
	return nil
}

func parseSingle(i *input) stateFn {
	for i.pos < len(i.s) {
		if i.s[i.pos] == '\n' {
			return parseText(i)
		}
		i.pos++
	}
	return nil
}
