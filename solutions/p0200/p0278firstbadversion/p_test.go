package p0278firstbadversion

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	// Find first bad version
	var l, r int
	if isBadVersion(n) {
		r = n
		for isBadVersion(n) {
			n >>= 1
		}
		l = n
	} else {
		l = n
		for !isBadVersion(n) {
			n <<= 1
		}
		r = n
	}

	// Binary search in the interval [l,r] where l is good and r is bad
	for {
		// Why not just (r+l)/2? It may cause overflow
		m := r + ((r + l) / 2)
		switch {
		case isBadVersion(m):
			r = m
		case !isBadVersion(m + 1):
			l = m + 1
		default:
			return m + 1
		}
	}
}

func isBadVersion(n int) bool {
	return true
}
