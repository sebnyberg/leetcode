package p0591tagvalidator

import (
	"errors"
	"fmt"
	"strings"
	"testing"
	"unicode"

	"github.com/stretchr/testify/require"
)

func Test_isValid(t *testing.T) {
	for _, tc := range []struct {
		code string
		want bool
	}{
		{"<DIV><></></DIV>", false},
		{"<DIV>This is the first line <![CDATA[<div>]]></DIV>", true},
		{"<DIV>>>  ![cdata[]] <![CDATA[<div>]>]]>]]>>]</DIV>", true},
		{"<A>  <B> </A>   </B>", false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.code), func(t *testing.T) {
			require.Equal(t, tc.want, isValid(tc.code))
		})
	}
}

func isValid(code string) bool {
	stack := []tag{}
	var i int
	t, b, e := parseTag(code)
	if e != nil || !b || t.typ != tagTypeOpen {
		return false
	}
	stack = append(stack, t)
	i += t.width
	for i < len(code) && len(stack) > 0 {
		t, b, e = parseTag(code[i:])
		if e != nil {
			return false
		}
		if !b {
			i++
			continue
		}
		// We have a tag
		switch t.typ {
		case tagTypeCData:
		case tagTypeClose:
			if t.name != stack[len(stack)-1].name {
				return false
			}
			stack = stack[:len(stack)-1]
		case tagTypeOpen:
			stack = append(stack, t)
		}
		i += t.width
	}

	return len(stack) == 0 && i == len(code)
}

const (
	tagTypeOpen  = 1
	tagTypeClose = 2
	tagTypeCData = 3
)

type tag struct {
	name  string
	width int
	typ   int
}

const cdataPat = "<![CDATA["
const cdataPatEnd = "]]>"

func parseTag(s string) (tag, bool, error) {
	if len(s) <= 2 || s[0] != '<' {
		return tag{}, false, nil
	}
	var pos int
	if len(s) >= len(cdataPat) && s[:len(cdataPat)] == cdataPat {
		// Parse CData contents. Must end with "]]>"
		pos += len(cdataPat)
		idx := strings.Index(s[pos:], cdataPatEnd)
		if idx == -1 {
			return tag{}, false, errors.New("invalid CDATA tag")
		}
		pos += idx + len(cdataPatEnd)
		return tag{
			name:  "CDATA",
			width: pos,
			typ:   tagTypeCData,
		}, true, nil
	}
	// Starts with '<'
	typ := tagTypeOpen
	pos++
	if s[1] == '/' {
		pos++
		typ = tagTypeClose
	}
	var name []byte
	for pos < len(s) && s[pos] != '>' {
		if !unicode.IsUpper(rune(s[pos])) {
			return tag{}, false, errors.New("invalid tag")
		}
		name = append(name, s[pos])
		if len(name) > 9 {
			return tag{}, false, errors.New("invalid tag")
		}
		pos++
	}
	if pos == len(s) || len(name) == 0 {
		return tag{}, false, errors.New("invalid tag")
	}
	pos++

	return tag{
		width: pos,
		name:  string(name),
		typ:   typ,
	}, true, nil
}
