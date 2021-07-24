package p0374guessnumberhigherorlower

import "math"

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guess(num int) int {
	return -1
}

func guessNumber(n int) int {
	l, r := 1, math.MaxInt32
	for l < r {
		mid := (r + l) / 2
		switch guess(mid) {
		case 0:
			return mid
		case -1:
			r = mid - 1
		case 1:
			l = mid + 1
		}
	}
	return l
}
