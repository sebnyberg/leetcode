package p2166designbitset

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBitSet(t *testing.T) {
	type any interface{}
	type action struct {
		name string
		args []any
		want []any
	}
	for _, tc := range []struct {
		name    string
		size    int
		actions []action
	}{
		{
			"advanced", 2,
			[]action{
				{"flip", []any{}, []any{}},
				{"unfix", []any{1}, []any{}},
				{"all", []any{}, []any{false}},
				{"fix", []any{1}, []any{}},
				{"fix", []any{1}, []any{}},
				{"unfix", []any{1}, []any{}},
				{"all", []any{}, []any{false}},
				{"count", []any{}, []any{1}},
				{"toString", []any{}, []any{"10"}},
				{"toString", []any{}, []any{"10"}},
				{"toString", []any{}, []any{"10"}},
				{"unfix", []any{0}, []any{}},
				{"flip", []any{}, []any{}},
				{"all", []any{}, []any{true}},
				{"unfix", []any{0}, []any{}},
				{"one", []any{}, []any{true}},
				{"one", []any{}, []any{true}},
				{"all", []any{}, []any{false}},
				{"fix", []any{0}, []any{}},
				{"unfix", []any{0}, []any{}},
			},
		},
		{
			"simple", 5,
			[]action{
				{"fix", []any{3}, []any{}},
				{"fix", []any{1}, []any{}},
				{"flip", []any{}, []any{}},
				{"all", []any{}, []any{false}},
				{"unfix", []any{0}, []any{}},
				{"flip", []any{}, []any{}},
				{"one", []any{}, []any{true}},
				{"unfix", []any{0}, []any{}},
				{"count", []any{}, []any{2}},
				{"toString", []any{}, []any{"01010"}},
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			b := Constructor(tc.size)
			for _, act := range tc.actions {
				switch act.name {
				case "fix":
					b.Fix(act.args[0].(int))
				case "unfix":
					b.Unfix(act.args[0].(int))
				case "flip":
					b.Flip()
				case "all":
					res := b.All()
					require.Equal(t, act.want[0].(bool), res)
				case "one":
					res := b.One()
					require.Equal(t, act.want[0].(bool), res)
				case "count":
					res := b.Count()
					require.Equal(t, act.want[0].(int), res)
				case "toString":
					res := b.ToString()
					require.Equal(t, act.want[0].(string), res)
				}
			}
		})
	}
}

type Bitset struct {
	vals    []bool
	count   [2]int
	flipped bool
}

func Constructor(size int) Bitset {
	return Bitset{
		vals:  make([]bool, size),
		count: [2]int{size, 0},
	}
}

func (this *Bitset) Fix(idx int) {
	if this.flipped {
		if !this.vals[idx] { // noop
			return
		}
		this.vals[idx] = false
		this.count[1]++
		this.count[0]--
	} else {
		if this.vals[idx] { // noop
			return
		}
		this.vals[idx] = true
		this.count[1]++
		this.count[0]--
	}
}

func (this *Bitset) Unfix(idx int) {
	if this.flipped {
		if this.vals[idx] {
			return
		}
		this.count[0]++
		this.count[1]--
		this.vals[idx] = true
	} else {
		if !this.vals[idx] {
			return
		}
		this.count[0]++
		this.count[1]--
		this.vals[idx] = false
	}
}

func (this *Bitset) Flip() {
	this.flipped = !this.flipped
	this.count[0], this.count[1] = this.count[1], this.count[0]
}

func (this *Bitset) All() bool {
	return this.count[0] == 0
}

func (this *Bitset) One() bool {
	return this.count[1] > 0
}

func (this *Bitset) Count() int {
	return this.count[1]
}

func (this *Bitset) ToString() string {
	res := make([]byte, 0, len(this.vals))
	for _, b := range this.vals {
		if this.flipped {
			if b {
				res = append(res, '0')
			} else {
				res = append(res, '1')
			}
		} else {
			if b {
				res = append(res, '1')
			} else {
				res = append(res, '0')
			}
		}
	}
	return string(res)
}
