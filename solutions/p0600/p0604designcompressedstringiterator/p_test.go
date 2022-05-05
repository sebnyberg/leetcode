package p0604designcompressedstringiterator

import "strconv"

type StringIterator struct {
	chars   []byte
	repeats []int
	i       int
}

func Constructor(compressedString string) StringIterator {
	s := compressedString
	var si StringIterator
	var i int
	for i < len(s) {
		// fetch character
		si.chars = append(si.chars, s[i])
		i++
		j := i
		for ; j < len(s) && s[j] >= '0' && s[j] <= '9'; j++ {
		}
		v, _ := strconv.Atoi(s[i:j])
		si.repeats = append(si.repeats, v)
		i = j
	}

	return si
}

func (this *StringIterator) Next() byte {
	for this.i < len(this.repeats) && this.repeats[this.i] == 0 {
		this.i++
	}
	if this.i >= len(this.repeats) {
		return ' '
	}
	this.repeats[this.i]--
	res := this.chars[this.i]
	for this.i < len(this.repeats) && this.repeats[this.i] == 0 {
		this.i++
	}
	return res
}

func (this *StringIterator) HasNext() bool {
	return this.i < len(this.repeats)
}
