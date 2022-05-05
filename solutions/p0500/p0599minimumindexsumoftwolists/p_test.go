package p0599minimumindexsumoftwolists

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findRestaurant(t *testing.T) {
	for _, tc := range []struct {
		list1 []string
		list2 []string
		want  []string
	}{
		{
			[]string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
			[]string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"},
			[]string{"Shogun"},
		},
		{
			[]string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
			[]string{"KFC", "Shogun", "Burger King"},
			[]string{"Shogun"},
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.list1), func(t *testing.T) {
			require.Equal(t, tc.want, findRestaurant(tc.list1, tc.list2))
		})
	}
}

func findRestaurant(list1 []string, list2 []string) []string {
	m := make(map[string]int, len(list1))
	for i, r := range list1 {
		m[r] = i
	}
	minSum := math.MaxInt32
	var res []string
	for i, r := range list2 {
		if v, exists := m[r]; exists {
			if v+i < minSum {
				res = res[:0]
				res = append(res, r)
				minSum = v + i
			} else if v+i == minSum {
				res = append(res, r)
			}
		}
	}
	return res
}
