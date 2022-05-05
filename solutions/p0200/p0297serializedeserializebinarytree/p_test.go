package p0297serializedeserializebinarytree

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"strings"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestCodec(t *testing.T) {
	ser := Constructor()
	deser := Constructor()
	root := &TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 2},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
	}
	data := ser.serialize(root)
	ans := deser.deserialize(data)
	_ = ans
	fmt.Println(ans)
}

type Codec struct {
	buf bytes.Buffer
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "nil"
	}
	this.buf.Reset()
	enc := gob.NewEncoder(&this.buf)
	err := enc.Encode(root)
	if err != nil {
		log.Fatalln(err)
	}
	return this.buf.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "nil" {
		return nil
	}
	b := strings.NewReader(data)
	dec := gob.NewDecoder(b)
	var res TreeNode
	err := dec.Decode(&res)
	if err != nil {
		log.Fatalln(err)
	}
	return &res
}
