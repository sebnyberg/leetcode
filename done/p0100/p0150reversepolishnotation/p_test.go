package p0150reversepolishnotation

import (
	"fmt"
	"log"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_evalRPN(t *testing.T) {
	for _, tc := range []struct {
		tokens []string
		want   int
	}{
		{[]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, 22},
		{[]string{"2", "1", "+", "3", "*"}, 9},
		{[]string{"4", "13", "5", "/", "+"}, 6},
	} {
		t.Run(fmt.Sprintf("%+v", tc.tokens), func(t *testing.T) {
			require.Equal(t, tc.want, evalRPN(tc.tokens))
		})
	}
}

func evalRPN(tokens []string) int {
	// Push tokens into the stack
	// When encountering an operand, pop the top 2 values off the stack,
	// calculate the value, and push it back into the stack
	// At the end, return the first and only value in the stack
	stack := make([]int, 0)
	n := 0
	for _, token := range tokens {
		switch token {
		case "+":
			stack[n-2] = stack[n-2] + stack[n-1]
			stack = stack[:n-1]
			n--
		case "-":
			stack[n-2] = stack[n-2] - stack[n-1]
			stack = stack[:n-1]
			n--
		case "*":
			stack[n-2] = stack[n-2] * stack[n-1]
			stack = stack[:n-1]
			n--
		case "/":
			stack[n-2] = stack[n-2] / stack[n-1]
			stack = stack[:n-1]
			n--
		default:
			val, err := strconv.Atoi(token)
			if err != nil {
				log.Fatalln(err)
			}
			stack = append(stack, val)
			n++
		}
	}
	return stack[0]
}
