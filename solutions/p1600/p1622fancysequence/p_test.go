package p1622fancysequence

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFancySequence(t *testing.T) {
	fancy := Constructor()
	fancy.Append(2)
	fancy.AddAll(3)
	fancy.Append(7)
	fancy.MultAll(2)
	res := fancy.GetIndex(0)
	require.Equal(t, 10, res)
	fancy.AddAll(3)
	fancy.Append(10)
	fancy.MultAll(2)
	res = fancy.GetIndex(0)
	require.Equal(t, 26, res)
	res = fancy.GetIndex(1)
	require.Equal(t, 34, res)
	res = fancy.GetIndex(2)
	require.Equal(t, 20, res)
}

func TestFancySequence2(t *testing.T) {
	fancy := Constructor()
	fancy.Append(5)
	res := fancy.GetIndex(0)
	require.Equal(t, 5, res)
	fancy.AddAll(4)
	fancy.Append(3)
	fancy.Append(4)
	fancy.Append(2)
	fancy.AddAll(7)
	fancy.GetIndex(0)
	res = fancy.GetIndex(0)
	require.Equal(t, 16, res)
	res = fancy.GetIndex(1)
	require.Equal(t, 10, res)
}

type Fancy struct {
	vals          []int
	insertIdx     []int
	insertIndices map[int]int
	ops           []op
}

func Constructor() Fancy {
	return Fancy{
		vals:          make([]int, 0),
		ops:           make([]op, 0),
		insertIndices: make(map[int]int),
	}
}

func (this *Fancy) Append(val int) {
	this.vals = append(this.vals, val)
	this.insertIdx = append(this.insertIdx, len(this.ops))
	this.insertIndices[len(this.ops)]++
}

func (this *Fancy) AddAll(inc int) {
	// Check if compression is possible
	n := len(this.ops)
	if n > 0 && this.insertIndices[n] == 0 && this.ops[n-1].typ == opTypeAdd {
		this.ops[n-1].val += inc
		return
	}
	this.ops = append(this.ops, op{
		typ: opTypeAdd,
		val: inc,
	})
}

func (this *Fancy) MultAll(m int) {
	// Check if compression is possible
	n := len(this.ops)
	if n > 0 && this.insertIndices[n] == 0 && this.ops[n-1].typ == opTypeMul {
		this.ops[n-1].val *= m
		this.ops[n-1].val %= 1000000007
		return
	}
	this.ops = append(this.ops, op{
		typ: opTypeMul,
		val: m,
	})
}

func (this *Fancy) GetIndex(idx int) int {
	if idx >= len(this.vals) {
		return -1
	}
	n := len(this.ops)
	res := this.vals[idx]
	for i := this.insertIdx[idx]; i < n; i++ {
		switch this.ops[i].typ {
		case opTypeAdd:
			res += this.ops[i].val
		case opTypeMul:
			res *= this.ops[i].val
		}
		res %= 1000000007
	}
	this.insertIndices[this.insertIdx[idx]]--
	this.insertIndices[n]++
	this.vals[idx] = res
	this.insertIdx[idx] = n
	return res
}

type op struct {
	typ int
	val int
}

const (
	opTypeMul = 0
	opTypeAdd = 1
)
