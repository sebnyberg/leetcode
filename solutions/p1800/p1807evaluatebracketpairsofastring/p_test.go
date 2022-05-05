package p1807evaluatebracketpairsofastring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_evaluate(t *testing.T) {
	for _, tc := range []struct {
		s         string
		knowledge [][]string
		want      string
	}{
		{"(name)is(age)yearsold", [][]string{{"name", "bob"}, {"age", "two"}}, "bobistwoyearsold"},
		{"hi(name)", [][]string{{"a", "b"}}, "hi?"},
		{"(a)(a)(a)aaa", [][]string{{"a", "yes"}}, "yesyesyesaaa"},
		{"(a)(b)", [][]string{{"a", "b"}, {"b", "a"}}, "ba"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, evaluate(tc.s, tc.knowledge))
		})
	}
}

func evaluate(s string, knowledge [][]string) string {
	res := make([]byte, 0)
	n := len(s)
	kv := make(map[string]string)
	for _, know := range knowledge {
		kv[know[0]] = know[1]
	}
	for i := 0; i < n; {
		if s[i] != '(' && s[i] != ')' {
			res = append(res, s[i])
			i++
			continue
		}
		l := i + 1
		for ; s[i] != ')'; i++ {
		}
		k := s[l:i]
		if v, exists := kv[k]; exists {
			res = append(res, []byte(v)...)
		} else {
			res = append(res, '?')
		}
		i++
	}

	return string(res)
}
