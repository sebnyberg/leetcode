package abc123

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_accountsMerge(t *testing.T) {
	type testCase struct {
		accounts [][]string
		want     [][]string
	}

	testCases := []testCase{
		{
			[][]string{{"John", "johnsmith@mail.com", "john_newyork@mail.com"}, {"John", "johnsmith@mail.com", "john00@mail.com"}, {"Mary", "mary@mail.com"}, {"John", "johnnybravo@mail.com"}},
			[][]string{{"John", "john00@mail.com", "john_newyork@mail.com", "johnsmith@mail.com"}, {"Mary", "mary@mail.com"}, {"John", "johnnybravo@mail.com"}},
		},
		{
			[][]string{{"Gabe", "Gabe0@m.co", "Gabe3@m.co", "Gabe1@m.co"}, {"Kevin", "Kevin3@m.co", "Kevin5@m.co", "Kevin0@m.co"}, {"Ethan", "Ethan5@m.co", "Ethan4@m.co", "Ethan0@m.co"}, {"Hanzo", "Hanzo3@m.co", "Hanzo1@m.co", "Hanzo0@m.co"}, {"Fern", "Fern5@m.co", "Fern1@m.co", "Fern0@m.co"}},
			[][]string{{"Ethan", "Ethan0@m.co", "Ethan4@m.co", "Ethan5@m.co"}, {"Gabe", "Gabe0@m.co", "Gabe1@m.co", "Gabe3@m.co"}, {"Hanzo", "Hanzo0@m.co", "Hanzo1@m.co", "Hanzo3@m.co"}, {"Kevin", "Kevin0@m.co", "Kevin3@m.co", "Kevin5@m.co"}, {"Fern", "Fern0@m.co", "Fern1@m.co", "Fern5@m.co"}},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("TestCase %v", i), func(t *testing.T) {
			require.ElementsMatch(t, tc.want, accountsMerge(tc.accounts))
		})
	}
}

func accountsMerge(accounts [][]string) [][]string {

	n := len(accounts)
	parent := make([]int, n+1)
	for i := range parent {
		parent[i] = i
	}
	var find func(a int) int
	find = func(a int) int {
		if parent[a] == a {
			return a
		}
		root := find(parent[a])
		parent[a] = root
		return root
	}
	union := func(a, b int) {
		ra, rb := find(a), find(b)
		if ra != rb {
			parent[rb] = ra
		}
	}

	// nameEmails := make(map[string][]string)
	emailIndices := make(map[string][]int)

	for i, acct := range accounts {
		for _, email := range acct[1:] {
			// nameEmails[acct[0]] = append(nameEmails[acct[0]], email)
			emailIndices[email] = append(emailIndices[email], i)
		}
	}

	// For each email, join together their accounts
	for _, indices := range emailIndices {
		for i := 0; i < len(indices)-1; i++ {
			union(indices[i], indices[i+1])
		}
	}

	accountEmails := make([]map[string]struct{}, n)
	for i := range accountEmails {
		accountEmails[i] = make(map[string]struct{})
	}
	for i, acct := range accounts {
		parent := find(i)
		for _, email := range acct[1:] {
			accountEmails[parent][email] = struct{}{}
		}
	}
	accountEmailsList := make([][]string, 0, n)
	for account, emails := range accountEmails {
		if len(emails) == 0 {
			continue
		}
		emailsList := make([]string, 0, len(emails))
		for email := range emails {
			emailsList = append(emailsList, email)
		}
		sort.Strings(emailsList)
		emailsList = append(emailsList, "")
		copy(emailsList[1:], emailsList)
		emailsList[0] = accounts[account][0]
		accountEmailsList = append(accountEmailsList, emailsList)
	}

	return accountEmailsList
}
