package p2069walkingrobotsimulation2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type any interface{}

func TestRobot(t *testing.T) {
	const (
		actionMove = iota
		actionGetPos
		actionGetDir
	)

	type action struct {
		name  int
		input []any
		want  []any
	}

	type testCase struct {
		width, height int
		actions       []action
	}

	testCases := []testCase{
		{
			6, 3, []action{
				{actionMove, []any{2}, nil},
				{actionMove, []any{2}, nil},
				{actionGetPos, []any{}, []any{[]int{4, 0}}},
				{actionGetDir, []any{}, []any{"East"}},
				{actionMove, []any{2}, nil},
				{actionMove, []any{1}, nil},
				{actionMove, []any{4}, nil},
				{actionGetPos, []any{}, []any{[]int{1, 2}}},
				{actionGetDir, []any{}, []any{"West"}},
			},
		},
		{
			6, 3, []action{
				{actionMove, []any{5}, nil},
				{actionGetDir, []any{}, []any{"East"}},
				{actionGetPos, []any{}, []any{[]int{5, 0}}},
				{actionMove, []any{1}, nil},
				{actionGetDir, []any{}, []any{"North"}},
				{actionGetPos, []any{}, []any{[]int{5, 1}}},
				{actionMove, []any{1}, nil},
				{actionGetDir, []any{}, []any{"North"}},
				{actionGetPos, []any{}, []any{[]int{5, 2}}},
				{actionMove, []any{1}, nil},
				{actionGetDir, []any{}, []any{"West"}},
				{actionGetPos, []any{}, []any{[]int{4, 2}}},
				{actionMove, []any{4}, nil},
				{actionGetDir, []any{}, []any{"West"}},
				{actionGetPos, []any{}, []any{[]int{0, 2}}},
				{actionMove, []any{1}, nil},
				{actionGetDir, []any{}, []any{"South"}},
				{actionGetPos, []any{}, []any{[]int{0, 1}}},
				{actionMove, []any{1}, nil},
				{actionGetDir, []any{}, []any{"South"}},
				{actionGetPos, []any{}, []any{[]int{0, 0}}},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			r := Constructor(tc.width, tc.height)
			for _, act := range tc.actions {
				switch act.name {
				case actionMove:
					r.Move(act.input[0].(int))
				case actionGetPos:
					res := r.GetPos()
					require.Equal(t, act.want[0].([]int), res)
				case actionGetDir:
					res := r.GetDir()
					require.Equal(t, act.want[0].(string), res)
				}
			}
		})
	}
}

type Robot struct {
	width, height int
	pos           int
	loopLength    int
	didMove       bool
}

func Constructor(width int, height int) Robot {
	return Robot{
		width:      width,
		height:     height,
		pos:        0,
		loopLength: width*2 + height*2 - 4,
		didMove:    false,
	}
}

func (this *Robot) Move(num int) {
	this.didMove = true
	this.pos += num
	this.pos %= this.width*2 + this.height*2 - 4
}

func (this *Robot) GetPos() []int {
	if !this.didMove || this.pos == 0 {
		return []int{0, 0}
	}

	switch this.GetDir() {
	case "East":
		return []int{this.pos, 0}
	case "North":
		return []int{this.width - 1, this.pos - this.width + 1}
	case "West":
		x := this.width - (this.pos - (this.width + this.height - 2)) - 1
		return []int{x, this.height - 1}
	case "South":
		y := this.height - (this.pos - (2*this.width + this.height - 3)) - 1
		return []int{0, y}
	default:
		panic("weird direction")
	}
}

func (this *Robot) GetDir() string {
	switch {
	case !this.didMove || this.pos > 0 && this.pos < this.width:
		return "East"
	case this.pos >= this.width && this.pos < this.width+this.height-1:
		return "North"
	case this.pos >= this.width+this.height-1 && this.pos < 2*this.width+this.height-2:
		return "West"
	default:
		return "South"
	}
}
