package p0537complexnumbermultiplication

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_complexNumberMultiply(t *testing.T) {
	for _, tc := range []struct {
		num1 string
		num2 string
		want string
	}{
		{"1+1i", "1+1i", "0+2i"},
		{"1+-1i", "1+-1i", "0+-2i"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.num1), func(t *testing.T) {
			require.Equal(t, tc.want, complexNumberMultiply(tc.num1, tc.num2))
		})
	}
}

func complexNumberMultiply(num1 string, num2 string) string {
	var r1, im1 float64
	fmt.Sscanf(num1, "%f+%fi", &r1, &im1)
	var r2, im2 float64
	fmt.Sscanf(num2, "%f+%fi", &r2, &im2)
	c1 := complex(r1, im1)
	c2 := complex(r2, im2)
	res := c1 * c2
	return fmt.Sprintf("%.0f+%.0fi", real(res), imag(res))
}
