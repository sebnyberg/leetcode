package p1733minimumnumberofpeopletoteach

import (
	"math"
	"sort"
)

func minimumTeachings(n int, languages [][]int, friendships [][]int) int {
	// I thought we could teach any language to any user.. which turned out to be
	// a very complex problem.
	// If we can only teach one language, we can just brute-force and check the
	// feasibility of each language being taught to everyone.

	for i := range languages {
		sort.Ints(languages[i]) // ensure that languages are sorted
	}
	for i := range friendships {
		for j := range friendships[i] {
			friendships[i][j]--
		}
	}

	m := len(languages) // number of users

	canCommunicate := make([]bool, len(friendships))
outerFor:
	for k, f := range friendships {
		a := f[0]
		b := f[1]
		var i, j int
		for i < len(languages[a]) && j < len(languages[b]) {
			la := languages[a][i]
			lb := languages[b][j]
			switch {
			case la == lb:
				canCommunicate[k] = true
				continue outerFor
			case la < lb:
				i++
			default:
				j++
			}
		}
	}

	try := func(i int) int {
		var res int
		// Try teaching everyone language i
		// Any friendship which already has a language in common should be omitted
		taught := make([]bool, m)
		for j := range languages {
			for _, l := range languages[j] {
				if l == i {
					taught[j] = true
					break
				}
			}
		}

		for j, f := range friendships {
			if canCommunicate[j] {
				continue
			}
			// Need to teach one or both the users this language
			a := f[0]
			b := f[1]
			if !taught[a] {
				taught[a] = true
				res++
			}
			if !taught[b] {
				taught[b] = true
				res++
			}
		}
		return res
	}

	mm := make(map[int]struct{})
	for _, l := range languages {
		for _, ll := range l {
			mm[ll] = struct{}{}
		}
	}
	res := math.MaxInt32
	for lang := range mm {
		res = min(res, try(lang))
	}
	return res
}
