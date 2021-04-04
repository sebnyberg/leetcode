package p0399evaluatedivision

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_calcEquation(t *testing.T) {
	for _, tc := range []struct {
		equations [][]string
		values    []float64
		queries   [][]string
		want      []float64
	}{
		{[][]string{
			{"x1", "x2"},
			{"x2", "x3"},
			{"x1", "x4"},
			{"x2", "x5"},
		}, []float64{
			3, 0.5, 3.4, 5.6,
		}, [][]string{
			{"x2", "x4"},
			{"x1", "x5"},
			{"x1", "x3"},
			{"x5", "x5"},
			{"x5", "x1"},
			{"x3", "x4"},
			{"x4", "x3"},
			{"x6", "x6"},
			{"x0", "x0"},
		}, []float64{
			1.13333, 16.8, 1.5, 1, 0.05952, 2.26667, 0.44118, -1, -1,
		}},
		{[][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}}, []float64{1.5, 2.5, 5.0}, [][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}}, []float64{3.75, 0.4, 5, 0.2}},
		{[][]string{{"a", "b"}}, []float64{0.5}, [][]string{{"a", "b"}, {"b", "a"}, {"a", "c"}, {"x", "y"}}, []float64{0.5, 2, -1, -1}},
		{[][]string{{"a", "b"}, {"b", "c"}}, []float64{2.0, 3.0}, [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}, []float64{6, 0.5, -1, 1, -1}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.equations), func(t *testing.T) {
			require.InEpsilonSlice(t, tc.want, calcEquation(tc.equations, tc.values, tc.queries), 0.1)
		})
	}
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	quotientSet := &QuotientSet{
		divisor:  make(map[string]string, len(equations)),
		quotient: make(map[string]float64, len(equations)),
	}
	for i, eq := range equations {
		quotientSet.Union(eq[0], eq[1], values[i])
	}

	res := make([]float64, len(queries))
	for i, query := range queries {
		switch {
		case !quotientSet.SameSet(query[0], query[1]):
			res[i] = -1
		default:
			_, qa := quotientSet.FindQuotient(query[0], 1)
			_, qb := quotientSet.FindQuotient(query[1], 1)
			totalQuotient := qa / qb
			res[i] = totalQuotient
		}
	}

	return res
}

type QuotientSet struct {
	divisor  map[string]string  // parent of a variable
	quotient map[string]float64 // distance from a variable to its parent
}

func (d *QuotientSet) ensureExists(a string) {
	if _, exists := d.divisor[a]; !exists {
		d.divisor[a] = a
		d.quotient[a] = 1
	}
}

func (d *QuotientSet) Union(a, b string, quotient float64) {
	d.ensureExists(a)
	d.ensureExists(b)

	ra, qa := d.FindQuotient(a, 1)
	rb, qb := d.FindQuotient(b, 1)

	if ra == rb { // relationship already exists
		return
	}

	// Link the two roots together
	d.quotient[ra] = (quotient * qb) / qa
	d.divisor[ra] = rb
}

func (d *QuotientSet) SameSet(a, b string) bool {
	if _, exists := d.divisor[a]; !exists {
		return false
	}
	if _, exists := d.divisor[b]; !exists {
		return false
	}
	ra, _ := d.FindQuotient(a, 1)
	rb, _ := d.FindQuotient(b, 1)
	return ra == rb
}

// FindQuotient returns the root of "a" and quotient between "a" and its root.
func (d *QuotientSet) FindQuotient(a string, dist float64) (string, float64) {
	if d.divisor[a] == a {
		return a, dist
	}
	return d.FindQuotient(d.divisor[a], dist*d.quotient[a])
}
