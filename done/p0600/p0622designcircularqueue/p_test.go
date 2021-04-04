package p0622designcircularqueue

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMyCircularQueue(t *testing.T) {
	type action struct {
		name string
		args []interface{}
		want []interface{}
	}

	for i, tc := range []struct {
		actions []action
	}{
		{
			[]action{
				{"Constructor", []interface{}{3}, nil},
				{"IsEmpty", nil, []interface{}{true}},
				{"enQueue", []interface{}{1}, []interface{}{true}},
				{"enQueue", []interface{}{2}, []interface{}{true}},
				{"enQueue", []interface{}{3}, []interface{}{true}},
				{"enQueue", []interface{}{4}, []interface{}{false}},
				{"Front", nil, []interface{}{1}},
				{"isEmpty", nil, []interface{}{false}},
				{"isFull", nil, []interface{}{true}},
				{"deQueue", nil, []interface{}{true}},
				{"Front", nil, []interface{}{2}},
				{"enQueue", []interface{}{4}, []interface{}{true}},
				{"Rear", nil, []interface{}{4}},
				{"isFull", nil, []interface{}{true}},
				{"deQueue", nil, []interface{}{true}},
				{"deQueue", nil, []interface{}{true}},
				{"deQueue", nil, []interface{}{true}},
				{"deQueue", nil, []interface{}{false}},
				{"isFull", nil, []interface{}{false}},
				{"isEmpty", nil, []interface{}{true}},
			},
		},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			q := Constructor(tc.actions[0].args[0].(int))
			for _, a := range tc.actions[1:] {
				switch a.name {
				case "enQueue":
					require.Equal(t, a.want[0].(bool), q.EnQueue(a.args[0].(int)))
				case "deQueue":
					require.Equal(t, a.want[0].(bool), q.DeQueue())
				case "Rear":
					require.Equal(t, a.want[0].(int), q.Rear())
				case "Front":
					require.Equal(t, a.want[0].(int), q.Front())
				case "isFull":
					require.Equal(t, a.want[0].(bool), q.IsFull())
				case "isEmpty":
					require.Equal(t, a.want[0].(bool), q.IsEmpty())
				}
			}
		})
	}
}

type MyCircularQueue struct {
	queue     []int
	insertPos int
	n         int
	maxlen    int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		queue:  make([]int, k),
		maxlen: k,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.n == this.maxlen {
		return false
	}
	this.queue[this.insertPos%this.maxlen] = value
	this.insertPos++
	this.n++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.n == 0 {
		return false
	}
	this.n--
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.n == 0 {
		return -1
	}
	return this.queue[(this.insertPos-this.n)%this.maxlen]
}

func (this *MyCircularQueue) Rear() int {
	if this.n == 0 {
		return -1
	}
	return this.queue[(this.insertPos-1)%this.maxlen]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.n == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.n == this.maxlen
}
