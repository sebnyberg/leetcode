package p1948deleteduplicatefoldersinsystem

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_deleteDuplicateFolder(t *testing.T) {
	for _, tc := range []struct {
		paths [][]string
		want  [][]string
	}{
		{
			[][]string{{"a"}, {"c"}, {"d"}, {"a", "b"}, {"c", "b"}, {"d", "a"}},
			[][]string{{"d"}, {"d", "a"}},
		},
		{
			[][]string{{"a"}, {"c"}, {"a", "b"}, {"c", "b"}, {"a", "b", "x"}, {"a", "b", "x", "y"}, {"w"}, {"w", "y"}},
			[][]string{{"c"}, {"c", "b"}, {"a"}, {"a", "b"}},
		},
		{
			[][]string{{"a", "b"}, {"c", "d"}, {"c"}, {"a"}},
			[][]string{{"c"}, {"c", "d"}, {"a"}, {"a", "b"}},
		},
		{
			[][]string{{"a"}, {"a", "x"}, {"a", "x", "y"}, {"a", "z"}, {"b"}, {"b", "x"}, {"b", "x", "y"}, {"b", "z"}},
			[][]string{},
		},
		{
			[][]string{{"a"}, {"a", "x"}, {"a", "x", "y"}, {"a", "z"}, {"b"}, {"b", "x"}, {"b", "x", "y"}, {"b", "z"}, {"b", "w"}},
			[][]string{{"b"}, {"b", "w"}, {"b", "z"}, {"a"}, {"a", "z"}},
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.paths), func(t *testing.T) {
			require.ElementsMatch(t, tc.want, deleteDuplicateFolder(tc.paths))
		})
	}
}

func deleteDuplicateFolder(paths [][]string) [][]string {
	// Build trie with paths
	root := new(TrieNode)
	root.next = make(map[string]*TrieNode)
	for _, path := range paths {
		cur := root
		for _, folderName := range path {
			cur = cur.AddNext(folderName)
		}
	}

	// Mark nodes which share subfolder structure for deletion
	d := Deduper{seen: make(map[string]*TrieNode)}
	for _, node := range root.next {
		d.dedupe(node)
	}

	// Traverse all paths to form the result
	var c ResultCollector
	for _, node := range root.next {
		c.collect(node)
	}

	return c.result
}

type Deduper struct {
	seen map[string]*TrieNode
}

func (d *Deduper) dedupe(cur *TrieNode) string {
	var subFolder strings.Builder
	for _, next := range cur.next {
		subFolder.WriteString(d.dedupe(next))
	}
	subFolderStr := subFolder.String()
	if len(cur.next) > 0 {
		if node, exists := d.seen[subFolderStr]; exists {
			node.delete = true
			cur.delete = true
		} else {
			d.seen[subFolderStr] = cur
		}
	}
	return "(" + cur.name + subFolder.String() + ")"
}

type ResultCollector struct {
	result [][]string
	path   []string
}

func (c *ResultCollector) collect(cur *TrieNode) {
	if cur.delete {
		return
	}
	c.path = append(c.path, cur.name)
	c.result = append(c.result, make([]string, len(c.path)))
	copy(c.result[len(c.result)-1], c.path)
	for _, next := range cur.next {
		c.collect(next)
	}
	c.path = c.path[:len(c.path)-1]
}

type TrieNode struct {
	name   string
	next   map[string]*TrieNode
	delete bool
}

func (n *TrieNode) AddNext(name string) *TrieNode {
	if _, exists := n.next[name]; !exists {
		n.next[name] = &TrieNode{
			name: name,
			next: make(map[string]*TrieNode),
		}
	}
	return n.next[name]
}
