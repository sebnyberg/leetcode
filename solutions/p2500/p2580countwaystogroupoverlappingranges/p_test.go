package p2580countwaystogroupoverlappingranges

import (
	"fmt"
	"sort"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_countWays(t *testing.T) {
	for i, tc := range []struct {
		ranges [][]int
		want   int
	}{
		{
			leetcode.ParseMatrix("[[57,92],[139,210],[306,345],[411,442],[533,589],[672,676],[801,831],[937,940],[996,1052],[1113,1156],[1214,1258],[1440,1441],[1507,1529],[1613,1659],[1773,1814],[1826,1859],[2002,2019],[2117,2173],[2223,2296],[2335,2348],[2429,2532],[2640,2644],[2669,2676],[2786,2885],[2923,2942],[3035,3102],[3177,3249],[3310,3339],[3450,3454],[3587,3620],[3725,3744],[3847,3858],[3901,3993],[4100,4112],[4206,4217],[4250,4289],[4374,4446],[4510,4591],[4675,4706],[4732,4768],[4905,4906],[5005,5073],[5133,5142],[5245,5309],[5352,5377],[5460,5517],[5569,5602],[5740,5791],[5823,5888],[6036,6042],[6096,6114],[6217,6262],[6374,6394],[6420,6511],[6564,6587],[6742,6743],[6797,6877],[6909,6985],[7042,7117],[7141,7144],[7276,7323],[7400,7456],[7505,7557],[7690,7720],[7787,7800],[7870,7880],[8013,8031],[8114,8224],[8272,8328],[8418,8435],[8493,8537],[8600,8704],[8766,8812],[8839,8853],[9032,9036],[9108,9189],[9222,9291],[9344,9361],[9448,9502],[9615,9673],[9690,9800],[9837,9868],[85,96],[145,202],[254,304],[372,411],[534,551],[629,692],[727,787],[861,944],[1041,1084],[1133,1174],[1260,1307],[1339,1358],[1478,1548],[1580,1618],[1694,1814],[1848,1891],[1936,1990],[2058,2130]]"),
			570065479,
		},
		{
			leetcode.ParseMatrix("[[10,11],[25,28],[29,34],[49,51],[58,66],[77,84],[91,91],[100,104],[120,121],[127,128],[151,151],[156,168],[0,8]]"),
			8192,
		},
		{
			leetcode.ParseMatrix("[[6,10],[5,15]]"),
			2,
		},
		{
			leetcode.ParseMatrix("[[1,3],[10,20],[2,5],[4,8]]"),
			4,
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, countWays(tc.ranges))
		})
	}
}

func countWays(ranges [][]int) int {
	// First merge any ranges that overlap, lets call these "components"
	// Then calculate the number of ways that two groups can be selected.
	type delta struct {
		i int
		d int
	}
	deltas := make([]delta, 0)
	for _, r := range ranges {
		deltas = append(deltas, delta{r[0], 1})
		deltas = append(deltas, delta{r[1] + 1, -1})
	}
	sort.Slice(deltas, func(i, j int) bool {
		a := deltas[i]
		b := deltas[j]
		if a.i == b.i {
			return a.d < b.d
		}
		return a.i < b.i
	})
	var n int
	var diff int
	for _, d := range deltas {
		diff += d.d
		if diff == 0 {
			n++
		}
	}
	// However the first group is formed, the second group must be the "rest".
	// Also the second group can be mirrored.
	//
	// So we want to calculate n over k (n choose k) which is given by
	// n!/(k!*(n-k)!) for each k = 0...n
	const mod = 1e9 + 7
	facs := make([]int, n+1)
	facs[0] = 1
	facs[1] = 1
	for x := 2; x <= n; x++ {
		facs[x] = (x * facs[x-1]) % mod
	}
	var res int
	for k := 0; k <= n; k++ {
		nf := facs[n]
		bot := modInverse((facs[k]*facs[n-k])%mod, mod)
		a := nf * bot
		res = (res + a) % mod
	}

	return res
}

func modInverse(a, mod int) int {
	return modPow(a, mod-2, mod)
}

func modPow(a, b, mod int) int {
	if b == 0 {
		return 1
	}
	p := modPow(a, b/2, mod) % mod
	p = p * p % mod
	if b%2 == 0 {
		return p
	}
	return (a * p) % mod
}
