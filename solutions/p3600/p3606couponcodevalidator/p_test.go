package p3606couponcodevalidator

import (
	"fmt"
	"regexp"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_validateCoupons(t *testing.T) {
	for _, tc := range []struct {
		code         []string
		businessLine []string
		isActive     []bool
		want         []string
	}{
		{
			[]string{"SAVE20", "", "PHARMA5", "SAVE@20"},
			[]string{"restaurant", "grocery", "pharmacy", "restaurant"},
			[]bool{true, true, true, true},
			[]string{"PHARMA5", "SAVE20"},
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.code), func(t *testing.T) {
			require.Equal(t, tc.want, validateCoupons(tc.code, tc.businessLine, tc.isActive))
		})
	}
}

func validateCoupons(code []string, businessLine []string, isActive []bool) []string {
	var res []string
	idx := make([]int, len(code))
	for i := range idx {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool {
		ii := idx[i]
		jj := idx[j]
		if businessLine[ii] == businessLine[jj] {
			return code[ii] < code[jj]
		}
		return businessLine[ii] < businessLine[jj]
	})

	var codePat = regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	validCode := func(code string) bool {
		return codePat.MatchString(code)
	}
	var businessPat = regexp.MustCompile(`^(electronics|grocery|pharmacy|restaurant)$`)
	validBusinessLine := func(s string) bool {
		return businessPat.MatchString(s)
	}

	for _, i := range idx {
		if validCode(code[i]) && validBusinessLine(businessLine[i]) && isActive[i] {
			res = append(res, code[i])
		}
	}

	return res
}
