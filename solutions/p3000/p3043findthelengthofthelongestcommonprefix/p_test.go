package p3043findthelengthofthelongestcommonprefix

import "fmt"

type trieNode struct {
	next [10]*trieNode
}

func longestCommonPrefix(arr1 []int, arr2 []int) int {
	root := &trieNode{}
	for _, x := range arr1 {
		curr := root
		for _, ch := range fmt.Sprint(x) {
			if curr.next[ch-'0'] == nil {
				curr.next[ch-'0'] = &trieNode{}
			}
			curr = curr.next[ch-'0']
		}
	}
	var res int
	for _, x := range arr2 {
		curr := root
		var count int
		for _, ch := range fmt.Sprint(x) {
			if curr.next[ch-'0'] == nil {
				break
			}
			curr = curr.next[ch-'0']
			count++
		}
		res = max(res, count)
	}
	return res
}
