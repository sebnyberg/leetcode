package p0379designphonedirectory

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPhoneDirectory(t *testing.T) {
	dir := Constructor(3)
	available := make(map[int]bool, 3)
	for i := 0; i < 3; i++ {
		available[i] = true
	}
	first := dir.Get()
	delete(available, first)
	second := dir.Get()
	delete(available, second)
	var lastAvailable int
	for k := range available {
		lastAvailable = k
		break
	}
	checkRes := dir.Check(lastAvailable)
	require.Equal(t, true, checkRes)
	getRes := dir.Get()
	require.Equal(t, lastAvailable, getRes)
	checkRes = dir.Check(lastAvailable)
	require.Equal(t, false, checkRes)
	dir.Release(lastAvailable)
	checkRes = dir.Check(lastAvailable)
	require.Equal(t, true, checkRes)
}

type PhoneDirectory struct {
	maxNumbers int
	available  map[int]struct{}
	taken      map[int]struct{}
}

func Constructor(maxNumbers int) PhoneDirectory {
	d := PhoneDirectory{
		maxNumbers: maxNumbers,
		available:  make(map[int]struct{}, maxNumbers),
		taken:      make(map[int]struct{}, maxNumbers),
	}
	for i := 0; i < maxNumbers; i++ {
		d.available[i] = struct{}{}
	}
	return d
}

// Get returns a phone number which is not yet taken. Returns -1 if no new
// number is available.
func (this *PhoneDirectory) Get() int {
	if len(this.available) == 0 {
		return -1
	}
	var res int
	for k := range this.available {
		res = k
		break
	}
	delete(this.available, res)
	this.taken[res] = struct{}{}
	return res
}

// Check if a number is available or not.
func (this *PhoneDirectory) Check(number int) bool {
	_, exists := this.available[number]
	return exists
}

// Release or recycle a number
func (this *PhoneDirectory) Release(number int) {
	delete(this.taken, number)
	this.available[number] = struct{}{}
}
