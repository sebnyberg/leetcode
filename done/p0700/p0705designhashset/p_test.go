package p0705designhashmap

type MyHashSet struct {
	sets [15626]uint64
}

const keyModulo = 15625

/** Initialize your data structure here. */
func Constructor() (m MyHashSet) {
	return
}

func (this *MyHashSet) Add(key int) {
	this.sets[key/64] |= (1 << (key % 64))
}

func (this *MyHashSet) Remove(key int) {
	this.sets[key/64] &^= (1 << (key % 64))
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	return this.sets[key/64]&(1<<(key%64)) > 0
}

/**
* Your MyHashSet object will be instantiated and called as such:
* obj := Constructor();
* obj.Add(key);
* obj.Remove(key);
* param_3 := obj.Contains(key);
 */
