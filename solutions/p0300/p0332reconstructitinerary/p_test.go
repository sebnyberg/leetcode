package p0332reconstructitinerary

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findItinerary(t *testing.T) {
	for _, tc := range []struct {
		tickets [][]string
		want    []string
	}{
		{
			[][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}},
			[]string{"JFK", "MUC", "LHR", "SFO", "SJC"},
		},
		{
			[][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}},
			[]string{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"},
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.tickets), func(t *testing.T) {
			require.Equal(t, tc.want, findItinerary(tc.tickets))
		})
	}
}

func findItinerary(tickets [][]string) []string {
	flights := make(map[string][]string)
	for _, ticket := range tickets {
		from, to := ticket[0], ticket[1]
		flights[from] = append(flights[from], to)
	}
	visits := make(map[string][]bool)
	for k := range flights {
		sort.Strings(flights[k])
		visits[k] = make([]bool, len(flights[k]))
	}
	res := visit(flights, visits, "JFK", []string{}, 1, len(tickets))
	return res
}

func visit(flights map[string][]string, visits map[string][]bool, airport string, prefix []string, nvisited, ntotal int) []string {
	prefix = append(prefix, airport)
	if nvisited == ntotal+1 {
		res := make([]string, len(prefix))
		copy(res, prefix)
		return res
	}
	// Try each possible ticket
	for i, destination := range flights[airport] {
		if visits[airport][i] {
			continue
		}
		// Attempt to fly to airport
		visits[airport][i] = true
		res := visit(flights, visits, destination, prefix, nvisited+1, ntotal)
		if len(res) == ntotal+1 {
			return res
		}
		visits[airport][i] = false
	}
	prefix = prefix[:len(prefix)-1]
	return prefix
}
