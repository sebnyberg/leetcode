package p1985findthekthlargestintegerinthearray

import (
	"fmt"
	"sort"
	"testing"
)

func Test_kthLargestNumber(t *testing.T) {
	type args struct {
		nums []string
		k    int
	}
	tests := []struct {
		args args
		want string
	}{
		{args{[]string{"3", "6", "7", "10"}, 4}, "3"},
		{args{[]string{"2", "21", "12", "1"}, 3}, "2"},
		{args{[]string{"0", "0"}, 2}, "0"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.args), func(t *testing.T) {
			if got := kthLargestNumber(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("kthLargestNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func kthLargestNumber(nums []string, k int) string {
	sort.Slice(nums, func(i, j int) bool {
		return less(nums[i], nums[j])
	})
	res := nums[len(nums)-k]
	return res
}

func less(a, b string) bool {
	if len(a) != len(b) {
		return len(a) < len(b)
	}
	for i := 0; i < len(a); i++ {
		if a[i] < b[i] {
			return true
		} else if a[i] > b[i] {
			return false
		}
	}
	return true
}
