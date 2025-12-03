package p3625countnumberoftrapezoidsii

import (
	"fmt"
	"math"
	"sort"
	"testing"

	"github.com/sebnyberg/leetcode"
)

func gcd2(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func abs2(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func TestFloatPrecision(t *testing.T) {
	points := leetcode.ParseMatrix("[[209,-385],[-35,319],[379,-93],[452,10],[179,163],[-118,196],[430,-365],[179,-365],[-299,465],[209,-410],[-375,-403],[-163,-227],[77,-365],[268,441],[460,465],[-163,465],[-412,-267],[-412,53],[-46,-280],[61,-209],[-234,32],[-35,296],[-276,-93],[-412,-475],[-470,-181],[-412,-283],[367,175],[-371,218],[209,-79],[-226,-74],[-435,-410],[-80,10],[-433,-365],[-35,-93],[-470,-67],[-378,0],[-82,-331],[144,268],[449,-106],[-470,-28],[452,-370],[449,-204],[-96,-245],[195,465],[-353,422],[-265,-2],[-178,219],[-35,222],[-375,-411],[-118,-93],[-199,71],[49,-209],[-301,-276],[79,219],[-46,32],[-35,181],[435,402],[449,465],[321,-209],[-412,148],[187,465],[367,496],[16,101],[179,244],[-346,151],[-353,319],[-251,-106],[-35,119],[-118,-370],[-102,465],[-35,311],[452,-62],[-118,441],[-412,-259],[375,441],[483,-182],[-35,-471],[462,289],[179,465],[-412,344],[206,302],[449,417],[-25,500],[-118,43],[372,-93],[180,167],[-118,473],[106,144],[151,-370],[-375,-245],[77,-399],[-478,465],[-405,-478],[-25,-291],[359,-347],[-371,-410],[179,250],[-353,206],[414,-494],[-118,-245],[-199,465],[28,-127],[435,175],[187,-436],[209,-250],[-35,-28],[-35,219],[-257,-93],[-25,-24],[-35,-236]]")

	// Group lines by slope and collect their m values
	linesBySlope := make(map[[2]int][]float64)

	for i := range points {
		for j := i + 1; j < len(points); j++ {
			a := points[i]
			b := points[j]

			var k [2]int
			var m float64
			if b[0] == a[0] {
				k[0] = 1
				k[1] = 0
				m = float64(a[0])
			} else if b[1] == a[1] {
				k[0] = 0
				k[1] = 1
				m = float64(a[1])
			} else {
				dx := b[0] - a[0]
				dy := b[1] - a[1]
				g := gcd2(abs2(dy), abs2(dx))
				k[0] = dx / g
				k[1] = dy / g
				if k[0] < 0 || (k[0] == 0 && k[1] < 0) {
					k[0] = -k[0]
					k[1] = -k[1]
				}
				kk := float64(dy) / float64(dx)
				m = float64(a[1]) - kk*float64(a[0])
			}

			linesBySlope[k] = append(linesBySlope[k], m)
		}
	}

	// Check for slopes with multiple close-but-different m values
	fmt.Println("Checking for floating point precision issues:")
	foundIssues := false

	for slope, mValues := range linesBySlope {
		if len(mValues) <= 1 {
			continue
		}

		// Sort m values
		sort.Float64s(mValues)

		// Check for values that are close but not exactly equal
		for i := 0; i < len(mValues)-1; i++ {
			diff := math.Abs(mValues[i+1] - mValues[i])
			if diff > 0 && diff < 1e-9 {
				fmt.Printf("Slope %v has close m values: %.15f and %.15f (diff: %e)\n",
					slope, mValues[i], mValues[i+1], diff)
				foundIssues = true
			}
		}

		// Also check for suspiciously many distinct m values for the same slope
		uniqueM := make(map[float64]int)
		for _, m := range mValues {
			uniqueM[m]++
		}

		if len(uniqueM) > len(mValues)/2 && len(mValues) > 4 {
			fmt.Printf("Slope %v: %d pairs, %d unique m values (suspicious?)\n",
				slope, len(mValues), len(uniqueM))
		}
	}

	if !foundIssues {
		fmt.Println("No obvious floating point precision issues found in m values")
	}
}
