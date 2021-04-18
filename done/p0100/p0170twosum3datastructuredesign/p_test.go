package p0170twosum3datastructuredesign

type TwoSum struct {
	nums map[int]int
}

/** Initialize your data structure here. */
func Constructor() TwoSum {
	return TwoSum{
		nums: make(map[int]int),
	}
}

/** Add the number to an internal data structure.. */
func (this *TwoSum) Add(number int) {
	this.nums[number]++
}

/** Find if there exists any pair of numbers which sum is equal to the value. */
func (this *TwoSum) Find(value int) bool {
	for n1, count := range this.nums {
		rest := value - n1
		if rest == n1 {
			if count >= 2 {
				return true
			}
		} else if this.nums[rest] >= 1 {
			return true
		}
	}
	return false
}

/**
 * Your TwoSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * param_2 := obj.Find(value);
 */
