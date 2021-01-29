package p0049groupanagrams

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
)

func Test_groupAnagrams(t *testing.T) {
	for _, tc := range []struct {
		strs []string
		want [][]string
	}{
		{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.strs), func(t *testing.T) {
			require.Equal(t, tc.want, groupAnagrams(tc.strs))
		})
	}
}

func AsBytes(s string) []byte {
	p := unsafe.Pointer((*reflect.StringHeader)(unsafe.Pointer(&s)).Data)

	var b []byte
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	hdr.Data = uintptr(p)
	hdr.Cap = len(s)
	hdr.Len = len(s)

	// maybeDetectMutations(b)
	return b
}

func AsString(b []byte) string {
	p := unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&b)).Data)

	var s string
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	hdr.Data = uintptr(p)
	hdr.Len = len(b)

	return s
}

func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)
	m := make(map[string][]string)
	for _, s := range strs {
		b := []rune(s)
		sort.Slice(b, func(i, j int) bool { return b[i] > b[j] })
		m[string(b)] = append(m[string(b)], s)
	}

	for _, v := range m {
		res = append(res, v)
	}
	return res
}
