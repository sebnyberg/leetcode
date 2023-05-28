package p1467probabilityofatwoboxechavingthesamenumberofdistinctballs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getProbability(t *testing.T) {
	for i, tc := range []struct {
		balls []int
		want  float64
	}{
		{[]int{2, 1, 1}, 0.66667},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.InEpsilon(t, tc.want, getProbability(tc.balls), eps)
		})
	}
}

const eps = 1e-5

func getProbability(balls []int) float64 {
	m := len(balls)
	var n int
	for _, b := range balls {
		n += b
	}
	n /= 2

	var mem [25]float64
	fac := func(x int) float64 {
		if mem[x] != 0 {
			return mem[x]
		}
		res := 1.0
		for k := float64(x); k > 1; k-- {
			res *= k
		}
		mem[x] = res
		return res
	}
	picked := make([]int, m)

	balanced, all := dfs(fac, balls, picked, 0, m, n, 0, 0)
	return balanced / all
}

func dfs(fac func(x int) float64, balls, picked []int, i, m, n, npicked, balance int) (float64, float64) {
	if i == m {
		if npicked != n {
			// Invalid
			return 0, 0
		}
		leftDups := 1.0
		rightDups := 1.0
		for i := range picked {
			leftDups *= fac(picked[i])
			rightDups *= fac(balls[i] - picked[i])
		}
		fn := fac(n)
		tot := (fn / leftDups) * (fn / rightDups)
		if balance == 0 {
			return tot, tot
		}
		return 0, tot
	}
	if npicked > n {
		return 0, 0
	}

	var balanced, all float64
	for left := 0; left <= balls[i]; left++ {
		b := balance
		if left == 0 {
			b++
		}
		if left == balls[i] {
			b--
		}
		picked[i] = left
		newBalanced, newAll := dfs(fac, balls, picked, i+1, m, n, npicked+left, b)
		balanced += newBalanced
		all += newAll
	}
	return balanced, all
}
