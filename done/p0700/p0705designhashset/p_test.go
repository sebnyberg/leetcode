package p0705designhashset

const nBuckets = 1024
const nBucketElems = 16
const bucketSize = nBucketElems * 64

type MyHashSet struct {
	items [nBuckets]*[nBucketElems]uint64
}

func Constructor() MyHashSet {
	return MyHashSet{}
}

func (this *MyHashSet) Add(key int) {
	bucket := key / bucketSize
	k := (key % bucketSize) / 64
	if this.items[bucket] == nil {
		this.items[bucket] = &[nBucketElems]uint64{}
	}
	this.items[bucket][k] |= 1 << (key % 64)
}

func (this *MyHashSet) Remove(key int) {
	bucket := key / bucketSize
	k := (key % bucketSize) / 64
	if this.items[bucket] == nil {
		return
	}
	this.items[bucket][k] &^= 1 << (key % 64)
}

func (this *MyHashSet) Contains(key int) bool {
	bucket := key / bucketSize
	k := (key % bucketSize) / 64
	return this.items[bucket] != nil && this.items[bucket][k]&(1<<(key%64)) > 0
}
