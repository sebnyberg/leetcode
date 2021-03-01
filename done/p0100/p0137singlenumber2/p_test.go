package p0136singlenumber2

func singleNumber(nums []int) (res int) {
	var x1, x2 int
	for _, n := range nums {
		x1 = (x1 ^ n) & ^x2
		x2 = (x2 ^ n) & ^x1
	}
	return x1
}
