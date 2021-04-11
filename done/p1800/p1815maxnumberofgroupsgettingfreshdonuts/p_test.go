package p1815maxnumberofgroupsgettingfreshdonuts

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxHappyGroups(t *testing.T) {
	for _, tc := range []struct {
		batchSize int
		groups    []int
		want      int
	}{
		{7, []int{2, 7, 5, 2, 3, 2, 6, 5, 3, 6, 2, 3, 7, 2, 2, 5, 4, 6, 6, 4, 7, 5, 6, 1, 6, 2, 6, 6, 2, 5}, 15},
		{9, []int{3, 1, 3, 3, 5, 6, 1, 1, 9, 10, 3, 3, 3, 1, 1, 3, 3, 3, 19, 20, 1, 3, 3, 3, 3, 1, 1, 3, 3, 30}, 9},
		{2, []int{369205928, 981877451, 947462486, 899465743, 737778942, 573732515, 520226542, 824581298, 571789442, 251943251, 70139785, 778962318, 43379662, 90924712, 142825931, 182207697, 178834435, 978165687}, 14},
		{2, []int{652231582, 818492002, 823729239, 2261354, 747144855, 478230860, 285970257, 774747712, 860954510, 245631565, 634746161, 109765576, 967900367, 340837477, 32845752, 23968185}, 12},
		{4, []int{1, 3, 2, 5, 2, 2, 1, 6}, 4},
		{3, []int{1, 2, 3, 4, 5, 6}, 4},
	} {
		t.Run(fmt.Sprintf("%v/%+v", tc.batchSize, tc.groups), func(t *testing.T) {
			require.Equal(t, tc.want, maxHappyGroups(tc.batchSize, tc.groups))
		})
	}
}

func maxHappyGroups(batchSize int, groups []int) int {
	groupSize := make([]int, batchSize)
	for _, g := range groups {
		groupSize[g%batchSize]++
	}
	happy := groupSize[0]
	groupSize[0] = 0

	// Match two-groups
	even := batchSize%2 == 0
	for k := 1; k <= batchSize/2; k++ {
		var nhappy int
		if k == batchSize/2 && even {
			nhappy = groupSize[k] / 2
			groupSize[k] -= 2 * (groupSize[k] / 2)
		} else {
			nhappy = min(groupSize[k], groupSize[batchSize-k])
			groupSize[k] -= nhappy
			groupSize[batchSize-k] -= nhappy
		}
		happy += nhappy
	}

	happy += maxMatch(groupSize, batchSize, 0, map[string]int{})

	// for _, g := range groupSize {
	// 	if g > 0 {
	// 		return happy + 1
	// 	}
	// }

	return happy
}

func maxMatch(groupSize []int, batchSize int, cur int, mem map[string]int) int {
	k := fmt.Sprintf("%v", groupSize)
	if v, exists := mem[k]; exists {
		return v
	}
	res := 0
	// if cur > 0 && groupSize[batchSize-cur] > 0 {
	// 	groupSize[batchSize-cur]--
	// 	res = maxMatch(groupSize, batchSize, 0, mem)
	// 	groupSize[batchSize-cur]++
	// } else {
	for k := 1; k < batchSize; k++ {
		if groupSize[k] == 0 {
			continue
		}
		groupSize[k]--
		subRes := maxMatch(groupSize, batchSize, (cur+k)%batchSize, mem)
		if cur == 0 {
			subRes++
		}
		res = max(res, subRes)
		groupSize[k]++
	}
	mem[k] = res
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
