package p0706designhashmap

const empty = -1
const tombstone = -2

type MyHashMap struct {
	entries []int
	keys    []int
	n       int
	nmax    int
}

func Constructor() MyHashMap {
	var m MyHashMap
	m.nmax = 70
	m.grow()
	return m
}

func findSlot(keys []int, key int) int {
	n := len(keys)
	j := key % n
	for keys[j] != empty && keys[j] != key {
		j = (j + 1) % n
	}
	return j
}

func (curr *MyHashMap) grow() {
	nmax := curr.nmax * 2
	entries := make([]int, nmax)
	keys := make([]int, nmax)
	for i := range keys {
		keys[i] = empty
	}
	var n int
	for i, k := range curr.keys {
		if k < 0 {
			continue
		}
		n++
		j := findSlot(keys, k)
		entries[j] = curr.entries[i]
		keys[j] = k
	}
	curr.nmax = nmax
	curr.keys = keys
	curr.entries = entries
	curr.n = n
}

func (this *MyHashMap) Put(key int, value int) {
	if this.n*100/this.nmax >= 70 {
		this.grow()
	}
	j := findSlot(this.keys, key)
	if this.keys[j] == empty {
		this.n++
	}
	this.keys[j] = key
	this.entries[j] = value
}

func (this *MyHashMap) Get(key int) int {
	if this.n*100/this.nmax >= 70 {
		this.grow()
	}
	j := findSlot(this.keys, key)
	if this.keys[j] >= 0 {
		return this.entries[j]
	}
	return -1

}

func (this *MyHashMap) Remove(key int) {
	j := findSlot(this.keys, key)
	if this.keys[j] != empty {
		this.keys[j] = tombstone
	}
}
