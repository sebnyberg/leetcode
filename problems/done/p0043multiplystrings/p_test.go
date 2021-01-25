package p0043multiplystrings

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_multiply(t *testing.T) {
	for _, tc := range []struct {
		num1 string
		num2 string
		want string
	}{
		{"12", "308", "3696"},
		{"2", "3", "6"},
		{"9133", "0", "0"},
		{"9133", "10", "91330"},
		{"123", "456", "56088"},
		{"1234", "1234", "1522756"},
		{"1234567891", "1234567891", "1524157877488187881"},
	} {
		t.Run(fmt.Sprintf("%v+%v", tc.num1, tc.num2), func(t *testing.T) {
			require.Equal(t, tc.want, multiply(tc.num1, tc.num2))
		})
	}
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	ret := make([]byte, 0, len(num1)+len(num2)-1)
	len1 := len(num1)
	len2 := len(num2)
	var m, carry byte
	for i := range num1 {
		for j := range num2 {
			if len(ret) < i+j+1 {
				ret = append(ret, 0)
			}
			n1 := num1[len1-1-i] - '0'
			n2 := num2[len2-1-j] - '0'
			m = ret[i+j] + carry + n1*n2
			ret[i+j] = m % 10
			carry = m / 10
		}
		for j := len2; carry > 0; j++ {
			if len(ret) < i+j+1 {
				ret = append(ret, 0)
			}
			m = ret[i+j] + carry
			ret[i+j] = m % 10
			carry = m / 10
		}
	}
	for i := range ret {
		ret[i] += '0'
	}
	for i, j := 0, len(ret)-1; i < j; i, j = i+1, j-1 {
		ret[i], ret[j] = ret[j], ret[i]
	}
	return string(ret)
}
