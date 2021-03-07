package p0706designhashmap

import "math"

type item struct {
	key   int
	value int
}

type MyHashMap struct {
	items [][]*item
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{
		items: make([][]*item, 1000),
	}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	for _, it := range this.items[key%1000] {
		if it.key == key {
			it.value = value
			return
		}
	}
	this.items[key%1000] = append(this.items[key%1000], &item{key, value})
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	for _, it := range this.items[key%1000] {
		if it.key == key {
			return it.value
		}
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	for _, it := range this.items[key%1000] {
		if it.key == key {
			it.key = math.MinInt64
			return
		}
	}
}
