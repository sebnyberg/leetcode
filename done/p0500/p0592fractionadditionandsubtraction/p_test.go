package p0592fractionadditionandsubtraction

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_fractionAddition(t *testing.T) {
	for _, tc := range []struct {
		expression string
		want       string
	}{
		{"5/3+1/3", "2/1"},
		{"-1/2+1/2", "0/1"},
		{"-1/2+1/2+1/3", "1/3"},
		{"1/3-1/2", "-1/6"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.expression), func(t *testing.T) {
			require.Equal(t, tc.want, fractionAddition(tc.expression))
		})
	}
}

func fractionAddition(expression string) string {
	// Split expression into parts
	parts := []string{}
	expression += "-" // sentinel for final pop
	for i, j := 0, 1; j < len(expression); j++ {
		if expression[j] == '-' || expression[j] == '+' {
			parts = append(parts, expression[i:j])
			i = j
		}
	}
	fracs := make([][2]int, len(parts))
	for i, p := range parts {
		var d, n int
		if n, _ := fmt.Sscanf(p, "%d/%d", &n, &d); n != 2 {
			return "err"
		}
		fracs[i][0] = n
		fracs[i][1] = d
	}
	// Find LCM of all expressions
	d := fracs[0][1]
	for i := 1; i < len(fracs); i++ {
		d = lcm(d, fracs[i][1])
	}
	var sum int
	for i := range fracs {
		mul := d / fracs[i][1]
		sum += fracs[i][0] * mul
	}
	if sum%d == 0 {
		return fmt.Sprintf("%v/1", sum/d)
	}
	sign := 1
	if sum < 0 {
		sign = -1
		sum = -sum
	}
	divisor := gcd(d, sum)
	res := fmt.Sprintf("%v/%v", (sum*sign)/divisor, d/divisor)
	return res
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
