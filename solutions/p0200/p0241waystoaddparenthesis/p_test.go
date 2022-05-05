package p0241waystoaddparenthesis

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_diffWaysToCompute(t *testing.T) {
	for _, tc := range []struct {
		input string
		want  []int
	}{
		// {"2-1-1", []int{0, 2}},
		{"2*3-4*5", []int{-34, -14, -10, -10, 10}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.input), func(t *testing.T) {
			require.ElementsMatch(t, tc.want, diffWaysToCompute(tc.input))
		})
	}
}

func diffWaysToCompute(input string) []int {
	i := strings.IndexAny(input, "+*-")
	if i == -1 {
		return []int{mustParse(input)}
	}

	res := make([]int, 0)
	for {
		res1, res2 := diffWaysToCompute(input[:i]), diffWaysToCompute((input[i+1:]))
		for _, a := range res1 {
			for _, b := range res2 {
				switch input[i] {
				case '+':
					res = append(res, a+b)
				case '-':
					res = append(res, a-b)
				case '*':
					res = append(res, a*b)
				default:
					log.Fatalln("invalid operator", input[i])
				}
			}
		}
		next := strings.IndexAny(input[i+1:], "+*-")
		if next == -1 {
			break
		}
		i = i + 1 + next
	}

	return res
}

func mustParse(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalln(err)
	}
	return n
}
