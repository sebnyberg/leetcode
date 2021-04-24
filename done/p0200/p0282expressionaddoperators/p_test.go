package p0282expressionaddoperators

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_addOperators(t *testing.T) {
	for _, tc := range []struct {
		num    string
		target int
		want   []string
	}{
		{"123", 6, []string{"1*2*3", "1+2+3"}},
		{"232", 8, []string{"2*3+2", "2+3*2"}},
		{"105", 5, []string{"1*0+5", "10-5"}},
		{"00", 0, []string{"0*0", "0+0", "0-0"}},
		{"3456237490", 9191, []string{}},
		{"2147483647", 2147483647, []string{"2147483647"}},
	} {
		t.Run(fmt.Sprintf("%v/%v", tc.num, tc.target), func(t *testing.T) {
			require.ElementsMatch(t, tc.want, addOperators(tc.num, tc.target))
		})
	}
}

func addOperators(num string, target int) []string {
	f := exprFinder{
		nums: make([]int, 10),
		ops:  make([]op, 10),
	}
	f.findNums(num, 0, target)
	return f.matchingExprs
}

type op int

const (
	add op = 0
	sub op = 1
	mul op = 2
)

var opStr = map[op]rune{
	add: '+',
	sub: '-',
	mul: '*',
}

type exprFinder struct {
	matchingExprs []string
	nums          []int
	numPos        int
	ops           []op
	opPos         int
}

func (f *exprFinder) findNums(num string, start int, target int) {
	if start == len(num) {
		f.findExprs(target)
		return
	}
	if num[start] == '0' {
		f.nums[f.numPos] = 0
		f.numPos++
		f.findNums(num, start+1, target)
		f.numPos--
	} else {
		var n int
		for i := start; i < len(num); i++ {
			n *= 10
			n += int(num[i] - '0')
			f.nums[f.numPos] = n
			f.numPos++
			f.findNums(num, i+1, target)
			f.numPos--
		}
	}
}

func (f *exprFinder) findExprs(target int) {
	// Base case
	if f.opPos == f.numPos-1 {
		res := evalExpr(f.nums[:f.numPos], f.ops[:f.opPos])
		if res == target {
			var sb strings.Builder
			sb.WriteString(strconv.Itoa(f.nums[0]))
			for i := 1; i < f.numPos; i++ {
				sb.WriteRune(opStr[f.ops[i-1]])
				sb.WriteString(strconv.Itoa(f.nums[i]))
			}
			f.matchingExprs = append(f.matchingExprs, sb.String())
		}
		return
	}

	f.ops[f.opPos] = add
	f.opPos++
	f.findExprs(target)
	f.opPos--

	f.ops[f.opPos] = sub
	f.opPos++
	f.findExprs(target)
	f.opPos--

	f.ops[f.opPos] = mul
	f.opPos++
	f.findExprs(target)
	f.opPos--
}

func evalExpr(vals []int, ops []op) int {
	if len(ops) != len(vals)-1 {
		panic("operators should weave values")
	}
	mulVals := []int{vals[0]}
	mulOps := []op{}
	for i := 0; i < len(vals)-1; i++ {
		if ops[i] == mul {
			mulVals[len(mulVals)-1] = mulVals[len(mulVals)-1] * vals[i+1]
		} else {
			mulVals = append(mulVals, vals[i+1])
			mulOps = append(mulOps, ops[i])
		}
	}
	res := mulVals[0]
	for i := 0; i < len(mulVals)-1; i++ {
		switch mulOps[i] {
		case add:
			res += mulVals[i+1]
		case sub:
			res -= mulVals[i+1]
		}
	}
	return res
}
