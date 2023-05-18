package p1233removesubfoldersfromthefilesystem

import (
	"sort"
	"strings"
)

type trieNode struct {
	next     map[string]*trieNode
	isParent bool
}

func removeSubfolders(folder []string) []string {
	// Use a trie to keep track of "seen" folders.
	// Each edge in the trie is a directory name
	// Each node in the trie marks whether it is a parent folder or not.
	// If a parent node is visited when adding a new folder, then skip it.
	root := &trieNode{
		next:     map[string]*trieNode{},
		isParent: false,
	}

	sort.Strings(folder)
	var res []string
	for _, f := range folder {
		curr := root
		for _, part := range strings.Split(f, "/")[1:] {
			var next *trieNode
			var exists bool
			if next, exists = curr.next[part]; exists {
				if next.isParent {
					goto hasParent
				}
			} else {
				curr.next[part] = &trieNode{
					next: map[string]*trieNode{},
				}
				next = curr.next[part]
			}
			curr = next
		}
		res = append(res, f)
		curr.isParent = true
	hasParent:
	}
	return res
}
