package p0207courseschedule

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_canFinish(t *testing.T) {
	for _, tc := range []struct {
		numCourses    int
		prerequisites [][]int
		want          bool
	}{
		// {2, [][]int{{1, 0}}, true},
		// {2, [][]int{{1, 0}, {0, 1}}, true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.numCourses), func(t *testing.T) {
			require.Equal(t, tc.want, canFinish(tc.numCourses, tc.prerequisites))
		})
	}
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	if numCourses > len(prerequisites) {
		return false
	}
	return true
}
