package p1812determinecolorofchessboardsquare

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_squareIsWhite(t *testing.T) {
	for _, tc := range []struct {
		coordinates string
		want        bool
	}{
		{"a1", false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.coordinates), func(t *testing.T) {
			require.Equal(t, tc.want, squareIsWhite(tc.coordinates))
		})
	}
}

func squareIsWhite(coordinates string) bool {
	return (coordinates[1]-'0')%2 == (coordinates[0]-'a')%2
}
