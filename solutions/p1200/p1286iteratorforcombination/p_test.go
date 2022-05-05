package p1286iteratorforcombination

import (
	"fmt"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/require"
)

type any interface{}

func Test_CombinationIterator(t *testing.T) {
	const (
		actionNext    = 0
		actionHasNext = 1
	)

	type action struct {
		name   int
		inputs []any
		want   []any
	}

	type testCase struct {
		characters        string
		combinationLength int
		actions           []action
	}
	testCases := []testCase{
		{
			"abc", 2, []action{
				{actionNext, nil, []any{"ab"}},
				{actionHasNext, nil, []any{true}},
				{actionNext, nil, []any{"ac"}},
				{actionHasNext, nil, []any{true}},
				{actionNext, nil, []any{"bc"}},
				{actionHasNext, nil, []any{false}},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%+v", tc.actions), func(t *testing.T) {
			combIt := Constructor(tc.characters, tc.combinationLength)
			for _, act := range tc.actions {
				switch act.name {
				case actionNext:
					require.Equal(t, act.want[0].(string), combIt.Next())
				case actionHasNext:
					require.Equal(t, act.want[0].(bool), combIt.HasNext())
				}
			}
		})
	}
}

type CombinationIterator struct {
	perms       chan string
	permRemains int64
}

func Constructor(characters string, combinationLength int) CombinationIterator {
	n := len(characters)
	k := combinationLength
	nperms := fac(n) / (fac(k) * fac(n-k))
	ci := CombinationIterator{
		perms:       make(chan string, 2),
		permRemains: int64(nperms),
	}
	perm := make([]byte, combinationLength)
	go func() {
		ci.findPermutations([]byte(characters), perm, 0, 0, combinationLength)
		close(ci.perms)
	}()
	return ci
}

func (this *CombinationIterator) findPermutations(
	characters []byte,
	perm []byte,
	idx, npicked, toPick int,
) {
	if npicked == toPick {
		this.perms <- string(perm)
		return
	}
	if idx >= len(characters) {
		return
	}
	perm[npicked] = characters[idx]
	this.findPermutations(characters, perm, idx+1, npicked+1, toPick)
	this.findPermutations(characters, perm, idx+1, npicked, toPick)
}

// Visit this position, picking then not picking the current character.
func (this *CombinationIterator) HasNext() bool {
	return atomic.LoadInt64(&this.permRemains) > 0
}

func (this *CombinationIterator) Next() string {
	atomic.AddInt64(&this.permRemains, -1)
	return <-this.perms
}

func fac(n int) int {
	res := 1
	for n >= 2 {
		res *= n
		n--
	}
	return res
}
