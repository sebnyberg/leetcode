package msft2_test

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
		{"0", "0", "0"},
		{"2", "3", "6"},
		{"123", "456", "56088"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.num1), func(t *testing.T) {
			res := multiply(tc.num1, tc.num2)
			require.Equal(t, tc.want, res)
		})
	}
}

func multiply(num1 string, num2 string) string {
	n1 := len(num1)
	n2 := len(num2)
	// Create a very large string to hold values
	res := make([]byte, 500)
	var offset int
	// For each digit in num1
	for i := n1 - 1; i >= 0; i-- {
		a := int(num1[i] - '0')
		// Multiply the digit with each digit in num2
		var carry int
		for j := n2 - 1; j >= 0; j-- {
			b := int(num2[j] - '0')
			// Multiply a, b
			ab := a * b

			// Add first digit of result to current position + c
			c := int(res[offset+n2-1-j])
			tot := ab + carry + c
			res[offset+n2-1-j] = byte(tot % 10)
			carry = tot / 10
		}
		j := 1
		for carry > 0 {
			c := int(res[offset+n2-1+j])
			tot := c + carry
			res[offset+n2-1+j] = byte(tot % 10)
			carry /= 10
			j++
		}
		offset++
		// Add carry
	}

	j := len(res) - 1
	for j > 0 && res[j] == 0 {
		j--
	}
	res = res[:j+1]
	// reverse
	for l, r := 0, len(res)-1; l < r; l, r = l+1, r-1 {
		res[l], res[r] = res[r], res[l]
	}
	// Convert to numbers
	for i := range res {
		res[i] += '0'
	}
	resStr := string(res)

	return resStr
}
