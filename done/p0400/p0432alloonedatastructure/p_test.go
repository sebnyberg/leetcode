package p0432alloonedatastructure

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAllOne(t *testing.T) {
	type any = interface{}
	type action struct {
		name string
		args []any
		want []any
	}
	for _, tc := range []struct {
		name    string
		actions []action
	}{
		{
			"complex",
			[]action{
				{"inc", []any{"a"}, []any{}},
				{"inc", []any{"a"}, []any{}},
				{"inc", []any{"a"}, []any{}},
				{"inc", []any{"a"}, []any{}},
				{"inc", []any{"b"}, []any{}},
				{"inc", []any{"b"}, []any{}},
				{"inc", []any{"c"}, []any{}},
				{"dec", []any{"c"}, []any{}},
				{"getMinKey", []any{}, []any{"b"}},
			},
		},
		{
			"simple",
			[]action{
				{"inc", []any{"hello"}, []any{}},
				{"inc", []any{"hello"}, []any{}},
				{"getMaxKey", []any{}, []any{"hello"}},
				{"getMinKey", []any{}, []any{"hello"}},
				{"inc", []any{"leet"}, []any{}},
				{"getMaxKey", []any{}, []any{"hello"}},
				{"getMinKey", []any{}, []any{"leet"}},
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			ao := Constructor()
			for _, act := range tc.actions {
				switch act.name {
				case "inc":
					ao.Inc(act.args[0].(string))
				case "dec":
					ao.Dec(act.args[0].(string))
				case "getMaxKey":
					require.Equal(t, act.want[0].(string), ao.GetMaxKey())
				case "getMinKey":
					require.Equal(t, act.want[0].(string), ao.GetMinKey())
				}
			}
		})
	}
}

type AllOne struct {
	valStrs []map[string]struct{}
	strVal  map[string]int
	minVal  int
	maxVal  int
}

func Constructor() AllOne {
	return AllOne{
		valStrs: make([]map[string]struct{}, 1, 100),
		strVal:  make(map[string]int, 100),
		minVal:  0,
		maxVal:  0,
	}
}

func (this *AllOne) Inc(key string) {
	v := this.strVal[key]
	if v >= 1 {
		// Remove current value
		delete(this.valStrs[v], key)
	}

	// Increase value
	if this.maxVal == v {
		this.maxVal = v + 1
	}
	if this.minVal == v && len(this.valStrs[v]) == 0 || v < this.minVal {
		this.minVal = v + 1
	}
	this.strVal[key] = v + 1
	if len(this.valStrs) <= this.maxVal {
		this.valStrs = append(this.valStrs, make(map[string]struct{}, 10))
	}
	this.valStrs[v+1][key] = struct{}{}
}

func (this *AllOne) Dec(key string) {
	// The key is guaranteed to exist
	v := this.strVal[key]
	delete(this.valStrs[v], key)
	this.strVal[key]--

	// Check if min val must be updated
	if this.minVal == v && len(this.valStrs[v]) == 0 {
		this.minVal = v - 1
		if v == 1 {
			for val, m := range this.valStrs {
				if len(m) > 0 {
					this.minVal = val
					break
				}
			}
		}
	}

	// Check if max val must be updated
	if this.maxVal == v && len(this.valStrs[v]) == 0 {
		this.maxVal = v - 1
	}

	v--
	if v >= 1 {
		this.valStrs[v][key] = struct{}{}
	}
}

func (this *AllOne) GetMaxKey() string {
	for k := range this.valStrs[this.maxVal] {
		return k
	}
	return ""
}

func (this *AllOne) GetMinKey() string {
	for k := range this.valStrs[this.minVal] {
		return k
	}
	return ""
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
