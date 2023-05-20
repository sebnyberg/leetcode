package p1295findnumberfwithevennumberofdigits

import "fmt"

func findNumbers(nums []int) int {
	var count int
	for _, x := range nums {
		count += 1 - len(fmt.Sprint(x))%2
	}
	return count
}
