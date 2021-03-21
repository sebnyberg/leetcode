package p0273integertoenglish

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numberToWords(t *testing.T) {
	for _, tc := range []struct {
		num  int
		want string
	}{
		{100, "One Hundred"},
		{1000, "One Thousand"},
		{1000000, "One Million"},
		{123, "One Hundred Twenty Three"},
		{12345, "Twelve Thousand Three Hundred Forty Five"},
		{1234567, "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"},
		{1234567891, "One Billion Two Hundred Thirty Four Million Five Hundred Sixty Seven Thousand Eight Hundred Ninety One"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.num), func(t *testing.T) {
			require.Equal(t, tc.want, numberToWords(tc.num))
		})
	}
}

var single = map[int]string{
	1: "One",
	2: "Two",
	3: "Three",
	4: "Four",
	5: "Five",
	6: "Six",
	7: "Seven",
	8: "Eight",
	9: "Nine",
}

var double = map[int]string{
	10: "Ten",
	11: "Eleven",
	12: "Twelve",
	13: "Thirteen",
	14: "Fourteen",
	15: "Fifteen",
	16: "Sixteen",
	17: "Seventeen",
	18: "Eighteen",
	19: "Nineteen",
	20: "Twenty",
	30: "Thirty",
	40: "Forty",
	50: "Fifty",
	60: "Sixty",
	70: "Seventy",
	80: "Eighty",
	90: "Ninety",
}

func numberToWords(num int) string {
	res := numbersToWordsRev(num)

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return strings.Join(res, " ")
}

func numbersToWordsRev(num int) []string {
	if num == 0 {
		return []string{"Zero"}
	}
	res := make([]string, 0)
	switch n := num % 100; {
	case n >= 1 && n < 10:
		res = append(res, single[n])
	case n >= 10 && n <= 20:
		res = append(res, double[n])
	case n > 20:
		if n%10 != 0 {
			res = append(res, single[n%10])
		}
		res = append(res, double[n-(n%10)])
	}
	if n := (num / 100) % 10; n != 0 {
		res = append(res, "Hundred")
		res = append(res, single[n])
	}
	postfixes := []string{"Thousand", "Million", "Billion"}
	for i := 0; num >= 1000; i++ {
		num /= 1000
		if num%1000 > 0 {
			subres := numbersToWordsRev(num % 1000)
			res = append(res, postfixes[i])
			res = append(res, subres...)
		}
	}

	return res
}
