package p0385miniparser

import (
	"testing"
)

type NestedInteger struct {
	val      int
	children []*NestedInteger
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (i *NestedInteger) IsInteger() bool {
	return len(i.children) == 0
}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (i *NestedInteger) GetInteger() int {
	return i.val
}

// Set this NestedInteger to hold a single integer.
func (i *NestedInteger) SetInteger(value int) {
	i.val = value
}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (i *NestedInteger) Add(elem NestedInteger) {
	// if len(i.children) == 0 {
	// 	i.children = append(i.children, &NestedInteger{val: i.val})
	// }
	i.children = append(i.children, &elem)
}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (i *NestedInteger) GetList() []*NestedInteger {
	return i.children
}

func Test_deserialize(t *testing.T) {
	s := "324"
	res := deserialize(s)
	s = "[123,[456,[789]]]"
	res = deserialize(s)
	s = "[100,300,[400,[500]]]"
	res = deserialize(s)
	_ = res
}

const eof = rune(0)

func deserialize(s string) *NestedInteger {
	s += string(eof)            // add EOF to ensure that last number is popped
	dummy := new(NestedInteger) // first element of dummy contains the result
	lists := intStack{dummy}    // stack of lists (.peek() returns current list)

	num := make([]rune, 0, 10)

	// If num is non-empty, create a new integer and add it to the current list
	maybeAddNum := func() {
		if len(num) == 0 {
			return
		}
		newInt := new(NestedInteger)
		newInt.SetInteger(parseInt(num))
		lists.peek().Add(*newInt)
		num = num[:0]
	}

	for _, ch := range s {
		switch ch {
		case '[':
			lists.push(new(NestedInteger))
		case ']':
			maybeAddNum()
			// Add current list to parent list
			list := lists.pop()
			lists.peek().Add(*list)
		case ',', eof:
			maybeAddNum()
		default:
			num = append(num, ch)
		}
	}
	return dummy.GetList()[0]
}

func parseInt(num []rune) int {
	sign := 1
	if num[0] == '-' {
		sign = -1
		num = num[1:]
	}
	var res int
	for _, ch := range num {
		res *= 10
		res += int(ch - '0')
	}
	return sign * res
}

type intStack []*NestedInteger

func (s *intStack) peek() *NestedInteger  { return (*s)[len(*s)-1] }
func (s *intStack) push(x *NestedInteger) { *s = append(*s, x) }
func (s *intStack) pop() *NestedInteger {
	n := len(*s)
	el := (*s)[n-1]
	*s = (*s)[:n-1]
	return el
}
