package p1776carfleet2

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func ParseMatrix(s string) [][]int {
	s = s[2 : len(s)-2]
	if s == "" {
		return nil
	}
	parts := strings.Split(s, "],[")
	res := make([][]int, len(parts))
	for i, part := range parts {
		if part == "" {
			continue
		}
		for _, numStr := range strings.Split(part, ",") {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				log.Fatalf("failed to parse number, %v, %v\n", numStr, err)
			}
			res[i] = append(res[i], num)
		}
	}
	return res
}

func Test_getCollisionTimes(t *testing.T) {
	for _, tc := range []struct {
		cars [][]int
		want []float64
	}{
		{ParseMatrix("[[1,2],[2,1],[4,3],[7,2]]"), []float64{1, -1, 3, -1}},
		{ParseMatrix("[[3,4],[5,4],[6,3],[9,1]]"), []float64{2.00000, 1.00000, 1.50000, -1.00000}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.cars), func(t *testing.T) {
			require.Equal(t, tc.want, getCollisionTimes(tc.cars))
		})
	}
}

func getCollisionTimes(cars [][]int) []float64 {
	// Whenever a set of cars are ascending in speed, there will be no collision
	// Whenever a set of cars are descending in speed, there will be a collision
	//
	// Consider three cars c1, c2, c3, descending in speed.
	//
	// There are two cases:
	// 1. c1 collides into c2, then c12 collides into c3
	// 2. c2 collides into c3, then c1 collides into c23
	//
	// Since c2 is slower than c1, the combined speed of c12 is the same as c2.
	// This means that c12 is guaranteed to collide into c3 at the same time as
	// c2 would crash into c3. This means that the collision time of c3 can be
	// calculated by comparing only c2 and c3.
	//
	// Continuing to c1 and c2. c2 is guaranteed to collide with c1, but it is
	// not certain that it happens before it collides into c3. To check this
	// condition, we check whether c1 will crash into c2 before c2's crash time.
	// If the answer is false, then c2 no longer needs to be considered for
	// anything.
	//
	// For any case where a pair of cars (c1, c2) are ascending in speed, c1
	// will not catch c2, so c2 can be disregarded.
	//
	// To implement this logic, use a monotinic stack, where all cars are
	// descending in speed.
	stack := []int{}
	const pos = 0
	const speed = 1
	res := make([]float64, len(cars))
	for i := range res {
		res[i] = -1
	}

	collisionTime := func(c1, c2 []int) float64 {
		v := c1[speed] - c2[speed]
		if v < 0 {
			return math.Inf(1)
		}
		return float64(c2[pos]-c1[pos]) / float64(v)
	}

	for i := len(cars) - 1; i >= 0; i-- {
		c1 := cars[i]
		for len(stack) > 0 {
			j := stack[len(stack)-1]
			c2 := cars[j]
			t := collisionTime(c1, c2)
			// If c1 is slower than c2, or c2 will collide before c1 collides with c2,
			// then remove c2 from consideration. This pair will not crash.
			if c1[speed] <= c2[speed] || res[j] != -1 && t >= res[j] {
				stack = stack[:len(stack)-1]
				continue
			}
			// c1 will collide with c2
			res[i] = t
			break
		}
		stack = append(stack, i)
	}
	return res
}
