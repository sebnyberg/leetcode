package p0929uniqueemail

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numUniqueEmails(t *testing.T) {
	for _, tc := range []struct {
		emails []string
		want   int
	}{
		{[]string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.emails), func(t *testing.T) {
			require.Equal(t, tc.want, numUniqueEmails(tc.emails))
		})
	}
}

func numUniqueEmails(emails []string) int {
	uniqueEmails := make(map[string]struct{})
	for _, email := range emails {
		parts := strings.Split(email, "@")
		name, domain := parts[0], parts[1]
		// Strip '+'
		if idx := strings.IndexRune(name, '+'); idx != -1 {
			name = name[:idx]
		}
		// Remove instances of '.'
		name = strings.ReplaceAll(name, ".", "")
		uniqueEmails[strings.Join([]string{name, domain}, "@")] = struct{}{}
	}
	var count int
	for range uniqueEmails {
		count++
	}
	return count
}
