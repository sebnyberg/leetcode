package p0721accountsmerge

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_accountsMerge(t *testing.T) {
	for _, tc := range []struct {
		accounts [][]string
		want     [][]string
	}{
		{
			[][]string{{"John", "johnsmith@mail.com", "john_newyork@mail.com"}, {"John", "johnsmith@mail.com", "john00@mail.com"}, {"Mary", "mary@mail.com"}, {"John", "johnnybravo@mail.com"}},
			[][]string{{"John", "john00@mail.com", "john_newyork@mail.com", "johnsmith@mail.com"}, {"Mary", "mary@mail.com"}, {"John", "johnnybravo@mail.com"}},
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.accounts), func(t *testing.T) {
			require.EqualValues(t, tc.want, accountsMerge(tc.accounts))
		})
	}
}

func accountsMerge(accounts [][]string) [][]string {
	dsu := NewDSU(len(accounts))

	// Collect emails for accounts and names for emails
	emailAccounts := make(map[string][]int)
	for i, account := range accounts {
		for _, email := range account[1:] {
			emailAccounts[email] = append(emailAccounts[email], i)
		}
	}

	// Union all accounts who share the same emails
	for _, accs := range emailAccounts {
		for i, acc := range accs {
			if i > 0 {
				dsu.Union(acc, accs[0])
			}
		}
	}

	accountEmails := make(map[int]map[string]struct{})
	for i, account := range accounts {
		j := dsu.Find(i)
		if _, exists := accountEmails[j]; !exists {
			accountEmails[j] = make(map[string]struct{})
		}
		for _, email := range account[1:] {
			accountEmails[j][email] = struct{}{}
		}
	}

	resp := make([][]string, 0, len(accountEmails))
	var respIdx int
	for accountIndex, emails := range accountEmails {
		resp = append(resp, make([]string, 0, len(emails)+1))
		resp[respIdx] = append(resp[respIdx], accounts[accountIndex][0])
		for email := range emails {
			resp[respIdx] = append(resp[respIdx], email)
		}
		sort.Strings(resp[respIdx][1:])
		respIdx++
	}

	// Find names for groups of emails
	return resp
}

type DSU struct {
	parent []int
	size   []int
}

func NewDSU(n int) *DSU {
	dsu := &DSU{
		parent: make([]int, n),
		size:   make([]int, n),
	}
	for i := range dsu.parent {
		dsu.parent[i] = i
		dsu.size[i] = 1
	}
	return dsu
}

func (d *DSU) Find(x int) int {
	if d.parent[x] != x {
		return d.Find(d.parent[x])
	}
	return x
}

func (d *DSU) Union(x, y int) {
	i, j := d.Find(x), d.Find(y)

	if d.size[i] < d.size[j] {
		d.parent[i] = j
		d.size[j] += d.size[i]
	} else {
		d.parent[j] = i
		d.size[i] += d.size[j]
	}
}
