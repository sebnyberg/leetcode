package p3625countnumberoftrapezoidsii

import (
	"fmt"
	"maps"
	"math"
	"slices"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_countTrapezoids(t *testing.T) {
	for _, tc := range []struct {
		points [][]int
		want   int
	}{
		{leetcode.ParseMatrix("[[125,-174],[-196,-434],[-149,-231],[102,-222],[-192,-473],[-458,-120],[-397,262],[243,46],[84,-222],[-261,-226],[199,-104],[-192,-444],[-144,499],[-171,262],[-425,-457],[-458,-271],[164,-496],[-449,-46],[-459,-102],[259,-222],[326,-260],[416,-495],[-404,98],[342,-469],[-165,499],[-317,112],[-459,446],[271,273],[460,426],[-395,-473],[-227,-50],[-358,284],[386,284],[-449,-133],[-495,361],[301,408],[-202,-222],[-122,-367],[199,-223],[301,499],[46,-243],[-251,196],[-449,406],[445,-174],[-458,-164],[-458,446],[-423,-222],[332,498],[24,446],[-409,262],[-309,-371],[-497,213],[445,475],[-75,262],[-488,499],[421,337],[69,-112],[120,-457],[45,-222],[221,196],[84,337],[386,-457],[-41,446],[14,-457],[410,446],[443,361],[-43,-317],[164,134],[-43,-308],[-290,-457],[-103,46],[-497,446],[423,-137],[155,445],[271,-18],[-175,152],[-458,-112],[34,45],[-122,499],[-497,-119],[-48,-119],[24,-222],[189,-457],[14,-7],[-289,-371],[-349,-50],[-335,-222],[-149,499],[-397,-112],[61,-457],[89,407],[199,352],[-189,319],[416,-135],[199,-222],[-132,499],[-380,-371],[-115,-260],[-337,475],[81,-285],[388,-222],[-37,-457],[34,-168],[98,-495],[84,-393],[148,-174],[-449,446],[-425,419],[61,-270],[-458,-469],[-413,-98],[326,49],[-158,-318],[46,446],[-380,152],[3,475],[-458,-442],[433,-469],[-192,89],[-272,227],[470,446]]"), 12521},
		{leetcode.ParseMatrix("[[209,-385],[-35,319],[379,-93],[452,10],[179,163],[-118,196],[430,-365],[179,-365],[-299,465],[209,-410],[-375,-403],[-163,-227],[77,-365],[268,441],[460,465],[-163,465],[-412,-267],[-412,53],[-46,-280],[61,-209],[-234,32],[-35,296],[-276,-93],[-412,-475],[-470,-181],[-412,-283],[367,175],[-371,218],[209,-79],[-226,-74],[-435,-410],[-80,10],[-433,-365],[-35,-93],[-470,-67],[-378,0],[-82,-331],[144,268],[449,-106],[-470,-28],[452,-370],[449,-204],[-96,-245],[195,465],[-353,422],[-265,-2],[-178,219],[-35,222],[-375,-411],[-118,-93],[-199,71],[49,-209],[-301,-276],[79,219],[-46,32],[-35,181],[435,402],[449,465],[321,-209],[-412,148],[187,465],[367,496],[16,101],[179,244],[-346,151],[-353,319],[-251,-106],[-35,119],[-118,-370],[-102,465],[-35,311],[452,-62],[-118,441],[-412,-259],[375,441],[483,-182],[-35,-471],[462,289],[179,465],[-412,344],[206,302],[449,417],[-25,500],[-118,43],[372,-93],[180,167],[-118,473],[106,144],[151,-370],[-375,-245],[77,-399],[-478,465],[-405,-478],[-25,-291],[359,-347],[-371,-410],[179,250],[-353,206],[414,-494],[-118,-245],[-199,465],[28,-127],[435,175],[187,-436],[209,-250],[-35,-28],[-35,219],[-257,-93],[-25,-24],[-35,-236]]"), 11025},
		{leetcode.ParseMatrix("[[92,100],[-4,55],[92,-87],[92,-91],[92,-30],[27,45],[66,82],[92,79],[92,-89],[-4,95],[92,-70],[-10,-18]]"), 21},
		{leetcode.ParseMatrix("[[-57,-76],[-94,-82],[-78,-18],[-64,-36],[-71,-82],[-47,-69],[-64,-82],[-64,-76]]"), 3},
		{leetcode.ParseMatrix("[[-3,2],[3,0],[2,3],[3,2],[2,-3]]"), 2},
		{leetcode.ParseMatrix("[[0,0],[1,0],[0,1],[2,1]]"), 1},
		{leetcode.ParseMatrix("[[32,-8],[-78,-25],[72,-8],[-77,-25]]"), 1},
		{leetcode.ParseMatrix("[[71,-89],[-75,-89],[-9,11],[-24,-89],[-51,-89],[-77,-89],[42,11]]"), 10},
	} {
		t.Run(fmt.Sprintf("%+v", tc.points), func(t *testing.T) {
			require.Equal(t, tc.want, countTrapezoids(tc.points))
		})
	}
}

type pair struct {
	a [2]int
	b [2]int

	k [2]int // dy/dx
	m float64
}

type lineKey struct {
	k [2]int
	m float64
}

func (p pair) points() [2][2]int {
	return [2][2]int{p.a, p.b}
}

func (p pair) slope() [2]int {
	return p.k
}

func (p pair) line() lineKey {
	return lineKey{p.k, p.m}
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func countTrapezoids(points [][]int) int {
	// Just as in the previous trapezoid problem, we want to match pairs of points
	//
	// We can do that by partitioning pairs of points by their y-offset (k=0) and slope
	//

	// The difficult part with this problem is realising that parallelograms are
	// double-counted. To detect a parallelogram, we can add the two points together
	// to form a sum. Two pairs of points that share the same mid-way sum are either
	// part of a parallelogram and should be deducted from the final result, or
	// are on the same slope. This means that we need to count parallelograms by
	// having a map of counts of segments sharing the same midpoint and deducting
	// the amount sharing the same midpoint and the same slope

	midPointCount := make(map[[2]int]int)
	midPointSlopeCount := make(map[[2][2]int]int)
	var parallelograms int

	// There are 500 points => 250000 pairs, which is OK to iterate over
	pairs := make(map[[2][2]int]*pair)
	lineCounts := make(map[lineKey]int)
	slopeCounts := make(map[[2]int]int)

	for i := range points {
		for j := i + 1; j < len(points); j++ {
			a := points[i]
			b := points[j]

			var k [2]int
			var m float64
			if b[0] == a[0] {
				// infinite slope.
				k[0] = 1 // arbitrary number, just needs to be non-zero
				k[1] = 0
				m = float64(a[0])
				// fmt.Println(a, b, k, m)
			} else if b[1] == a[1] {
				// flatline
				k[0] = 0
				k[1] = 1 // again, arbitrary number.
				m = float64(a[1])
			} else {
				dx := b[0] - a[0]
				dy := b[1] - a[1]
				g := gcd(abs(dy), abs(dx))
				k[0] = dx / g
				k[1] = dy / g
				// Normalize slope direction: ensure first non-zero component is positive
				if k[0] < 0 || (k[0] == 0 && k[1] < 0) {
					k[0] = -k[0]
					k[1] = -k[1]
				}
				kk := float64(dy) / float64(dx)
				m = float64(a[1]) - kk*float64(a[0])

				// round to something
				ratio := math.Pow(10, 11)
				m = math.Round(m*ratio) / ratio
			}

			p := &pair{
				a: [2]int{a[0], a[1]},
				b: [2]int{b[0], b[1]},
				k: k,
				m: m,
			}
			pairs[p.points()] = p
			slopeCounts[p.slope()]++
			lineCounts[p.line()]++

			// Note: this is not exactly the midpoint, that would need division by 2
			// but dividing would potentially result in a non-integer value
			midPointish := [2]int{a[0] + b[0], a[1] + b[1]}
			newParallellograms := midPointCount[midPointish] - midPointSlopeCount[[2][2]int{midPointish, k}]
			parallelograms += newParallellograms
			midPointCount[midPointish]++
			midPointSlopeCount[[2][2]int{midPointish, k}]++
		}
	}

	// For each unique line (which also has a slope), count combinations of pairs in that line
	// and pairs from any other line. When done, this line is no longer eligible for creating
	// a trapezoid.
	lines := slices.Collect(maps.Keys(lineCounts))
	var res int
	for i := range lines {
		l := lines[i]
		k := l.k

		// We can't pair points on this line with points from other pairs on this line,
		// so we remove them from the slope counts.
		slopeCounts[k] -= lineCounts[l] // remove pairs from this line from the total slope count

		// We can pair any pair of points from this line with any other pair of points sharing the
		// same slope.
		res += lineCounts[l] * slopeCounts[k]
	}

	return res - parallelograms
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
