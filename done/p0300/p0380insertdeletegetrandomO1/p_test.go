package p0380insertdeletegetrandomo1

import "math/rand"

type RandomizedSet struct {
	itemIndex map[int]int
	itemsList []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		itemIndex: make(map[int]int, 1e4),
		itemsList: make([]int, 0, 1e4),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, exists := this.itemIndex[val]; exists {
		return false
	}
	this.itemIndex[val] = len(this.itemsList)
	this.itemsList = append(this.itemsList, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	itemIdx, exists := this.itemIndex[val]
	if !exists {
		return false
	}
	// swap current and last index
	n := len(this.itemsList)
	this.itemsList[itemIdx], this.itemsList[n-1] = this.itemsList[n-1], this.itemsList[itemIdx]
	this.itemIndex[this.itemsList[itemIdx]] = itemIdx
	delete(this.itemIndex, val)
	this.itemsList = this.itemsList[:n-1]
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.itemsList[rand.Intn(len(this.itemsList))]
}
