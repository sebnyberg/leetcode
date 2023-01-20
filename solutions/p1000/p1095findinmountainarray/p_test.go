package p1095findinmountainarray

import "fmt"

/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */

func findInMountainArray(target int, mountainArr *MountainArray) int {
	var lo int
	hi := mountainArr.length() - 1
	var peak int
	for k := 0; k < 100; k++ {
		m := lo + (hi-lo)/2
		pre := mountainArr.get(m - 1)
		mid := mountainArr.get(m)
		next := mountainArr.get(m + 1)
		if pre < mid && mid > next {
			fmt.Println("found peak")
			peak = m
			break
		}
		if pre < mid {
			lo = m
		} else if mid > next {
			hi = m
		} else {
			panic("wut")
		}
	}
	lo = 0
	hi = peak
	for lo < hi {
		mid := lo + (hi-lo)/2
		v := mountainArr.get(mid)
		if v == target {
			return mid
		}
		if target > v {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	lo = peak
	hi = mountainArr.length()
	for lo < hi {
		mid := lo + (hi-lo)/2
		v := mountainArr.get(mid)
		if v == target {
			return mid
		}
		if target < v {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return -1
}
