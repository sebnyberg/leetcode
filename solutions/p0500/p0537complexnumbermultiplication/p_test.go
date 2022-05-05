package p0537complexnumbermultiplication

import (
	"fmt"
	"log"
	"strconv"
	"strings"
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
		t.Run(fmt.Sprintf("%v + %v", tc.num1, tc.num2), func(t *testing.T) {
			require.Equal(t, tc.want, complexNumberMultiply(tc.num1, tc.num2))
		})
	}
}

func complexNumberMultiply(num1 string, num2 string) string {
	a, b := parseComplex(num1), parseComplex(num2)
	real := a.real*b.real - (a.im * b.im)
	im := a.real*b.im + b.real*a.im
	res := complexNum{real, im}
	return res.String()
}

type complexNum struct {
	real int
	im   int
}

func (n complexNum) String() string {
	return fmt.Sprintf("%v+%vi", n.real, n.im)
}

func parseComplex(num string) complexNum {
	parts := strings.Split(num, "+")
	return complexNum{
		real: mustParseNum(parts[0]),
		im:   mustParseNum(parts[1][:len(parts[1])-1]),
	}
}

func mustParseNum(num string) int {
	n, err := strconv.Atoi(num)
	if err != nil {
		log.Fatalln(err)
	}
	return n
}
