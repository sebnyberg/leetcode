package p0588designaninmemoryfilesystem

import (
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFileSystem(t *testing.T) {
	t.Run("a", func(t *testing.T) {
		fs := Constructor()
		require.Equal(t, []string{}, fs.Ls("/"))
		fs.Mkdir("/a/b/c")
		fs.AddContentToFile("/a/b/c/d", "hello")
		require.Equal(t, []string{"a"}, fs.Ls("/"))
		require.Equal(t, "hello", fs.ReadContentFromFile("/a/b/c/d"))
	})
	t.Run("b", func(t *testing.T) {
		fs := Constructor()
		fs.Mkdir("/goowmfn")
		require.Equal(t, []string{}, fs.Ls("/goowmfn"))
		require.Equal(t, []string{"goowmfn"}, fs.Ls("/"))
		fs.Mkdir("/z")
		require.Equal(t, []string{"goowmfn", "z"}, fs.Ls("/"))
		fs.AddContentToFile("/goowmfn/c", "shetopcy")
		require.Equal(t, []string{}, fs.Ls("/z"))
		require.Equal(t, []string{"c"}, fs.Ls("/goowmfn/c"))
	})
}

type nodeType int8

const (
	nodeTypeDir  = 1
	nodeTypeFile = 2
)

type node struct {
	name       string
	parentPath string
	children   []*node
	contents   string
	typ        nodeType
}

func (n *node) findChild(name string) *node {
	for _, n := range n.children {
		if n.name == name {
			return n
		}
	}
	return nil
}

func (n *node) absPath() string {
	return n.parentPath + n.name
}

type FileSystem struct {
	root *node
}

func Constructor() FileSystem {
	return FileSystem{
		root: &node{
			parentPath: "",
			name:       "",
			typ:        nodeTypeDir,
		},
	}
}

func (this *FileSystem) Ls(path string) []string {
	dirs := strings.Split(path[1:], "/")
	cur := this.root
	if len(dirs) == 1 && dirs[0] == "" {
		dirs = dirs[:0]
	}
	// Traverse parts of the path
	for _, relPath := range dirs {
		cur = cur.findChild(relPath)
		if cur == nil {
			return []string{}
		}
	}
	if cur.typ == nodeTypeFile {
		return []string{cur.name}
	}
	res := make([]string, len(cur.children))
	for i, c := range cur.children {
		res[i] = c.name
	}
	sort.Strings(res)
	return res
}

func (this *FileSystem) Mkdir(path string) {
	dirs := strings.Split(path[1:], "/")
	cur := this.root
	for _, dir := range dirs {
		next := cur.findChild(dir)
		if next == nil { // create directory
			next = &node{
				name:       dir,
				parentPath: cur.absPath() + "/",
				children:   []*node{},
				typ:        nodeTypeDir,
			}
			cur.children = append(cur.children, next)
		}
		cur = next
	}
}

func (this *FileSystem) AddContentToFile(filePath string, content string) {
	pathParts := strings.Split(filePath[1:], "/")
	n := len(pathParts)
	cur := this.root
	for _, dir := range pathParts[:n-1] {
		cur = cur.findChild(dir)
	}
	f := cur.findChild(pathParts[n-1])
	if f == nil {
		f = &node{
			name:       pathParts[n-1],
			parentPath: cur.absPath() + "/",
			typ:        nodeTypeFile,
		}
		cur.children = append(cur.children, f)
	}
	f.contents += content
}

func (this *FileSystem) ReadContentFromFile(filePath string) string {
	pathParts := strings.Split(filePath[1:], "/")
	cur := this.root
	for _, dir := range pathParts {
		cur = cur.findChild(dir)
	}
	return cur.contents
}
