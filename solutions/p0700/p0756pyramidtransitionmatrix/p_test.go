package p0756pyramidtransitionmatrix

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_pyramidTransition(t *testing.T) {
	for _, tc := range []struct {
		bottom  string
		allowed []string
		want    bool
	}{

		{"AFFFFA", []string{"ADA", "ADC", "ADB", "AEA", "AEC", "AEB", "AFA", "AFC", "AFB", "CDA", "CDC", "CDB", "CEA", "CEC", "CEB", "CFA", "CFC", "CFB", "BDA", "BDC", "BDB", "BEA", "BEC", "BEB", "BFA", "BFC", "BFB", "DAA", "DAC", "DAB", "DCA", "DCC", "DCB", "DBA", "DBC", "DBB", "EAA", "EAC", "EAB", "ECA", "ECC", "ECB", "EBA", "EBC", "EBB", "FAA", "FAC", "FAB", "FCA", "FCC", "FCB", "FBA", "FBC", "FBB", "DDA", "DDC", "DDB", "DEA", "DEC", "DEB", "DFA", "DFC", "DFB", "EDA", "EDC", "EDB", "EEA", "EEC", "EEB", "EFA", "EFC", "EFB", "FDA", "FDC", "FDB", "FEA", "FEC", "FEB", "FFA", "FFC", "FFB", "DDD", "DDE", "DDF", "DED", "DEE", "DEF", "DFD", "DFE", "DFF", "EDD", "EDE", "EDF", "EED", "EEE", "EEF", "EFD", "EFE", "EFF", "FDD", "FDE", "FDF", "FED", "FEE", "FEF", "FFD", "FFE", "FFF"},
			false},
		{"BCD", []string{"BCC", "CDE", "CEA", "FFF"}, true},
		{"AAAA", []string{"AAB", "AAC", "BCD", "BBE", "DEF"}, false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.bottom), func(t *testing.T) {
			require.Equal(t, tc.want, pyramidTransition(tc.bottom, tc.allowed))
		})
	}
}

func pyramidTransition(bottom string, allowed []string) bool {
	seen := make(map[string]struct{})
	var ok [26][26][26]bool
	for _, a := range allowed {
		ok[a[0]-'A'][a[1]-'A'][a[2]-'A'] = true
	}
	res := dfs(seen, &ok, bottom, "", 0)
	return res
}

func dfs(
	seen map[string]struct{},
	ok *[26][26][26]bool,
	curr string,
	next string,
	i int,
) bool {
	if _, exists := seen[curr]; exists {
		return false
	}
	if len(curr) == 1 {
		return true
	}
	if i == len(curr)-1 {
		res := dfs(seen, ok, next, "", 0)
		seen[next] = struct{}{}
		return res
	}
	for ch := 0; ch < 26; ch++ {
		a, b := curr[i]-'A', curr[i+1]-'A'
		if ok[a][b][ch] {
			if dfs(seen, ok, curr, next+string(byte(ch+'A')), i+1) {
				return true
			}
		}
	}
	return false
}
