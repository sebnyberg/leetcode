package dsx

func GenPermsDFS(vals []int) chan []int {
	n := len(vals)
	perm := make([]int, n)
	res := make(chan []int)

	go func() {
		var dfs func(int, int)
		dfs = func(idx, seen int) {
			if idx == n {
				cpy := make([]int, n)
				copy(cpy, perm)
				res <- cpy
			}
			for i := 0; i < n; i++ {
				if seen&(1<<i) == 0 {
					perm[idx] = vals[i]
					dfs(idx+1, seen|(1<<i))
				}
			}
		}
		dfs(0, 0)
		close(res)
	}()

	return res
}

const maxFac = 15

// See https://en.wikipedia.org/wiki/Factorial_number_system
// and https://en.wikipedia.org/wiki/Lehmer_code
//
// Note: this is not more efficient than using a simple
func GenPermsLehmer(vals []int) chan []int {
	n := len(vals)
	if n > maxFac {
		panic("list of values is too long")
	}
	// Imagine a set of available values {1,2,3,...,14}
	// For any permutation, the first value will change fourteen times in total.
	// The total number of permutations is 14!, and so the number will change
	// every 13! permutations.
	//
	// Given some number between 1 and 14!, the number can be divided by 13! to
	// figure out which alternative should be in the first position. For example,
	// if i / 13! = 2, then the permutation starts with {3, x, y, ...} and the
	// list of remaining alternatives is now {1,2,4,5,...,14}.
	//
	// Then we must find the permutation within the remaining 13 numbers. By
	// taking i % 13!, the remainder is the permutation # within the remaining
	// numbers.
	//
	// Then the process continues until there are no more alternatives.
	//

	res := make(chan []int)

	go func() {
		// Calculate factorial bases
		fac := make([]int, n+1)
		fac[0] = 0
		fac[1] = 1
		for i := 2; i <= n; i++ {
			fac[i] = fac[i-1] * i
		}

		perm := make([]int, n)

		// For each possible permutation
		for i := 0; i < fac[n]; i++ {
			var seen [maxFac]bool // keep track of 'seen' numbers
			permNumber := i
			for permPos := n - 1; permPos >= 0; permPos-- {
				permNumber %= fac[permPos+1]
				var altPos int
				if permPos != 0 {
					altPos = permNumber / fac[permPos]
				}
				var j int
				for seen[j] || altPos > 0 {
					if !seen[j] {
						altPos--
					}
					j++
				}
				perm[n-1-permPos] = vals[j]
				seen[j] = true
			}
			cpy := make([]int, n)
			copy(cpy, perm)
			res <- cpy
		}
		close(res)
	}()

	return res
}
