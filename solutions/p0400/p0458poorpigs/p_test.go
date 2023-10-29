package p0458poorpigs

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_poorPigs(t *testing.T) {
	for _, tc := range []struct {
		buckets       int
		minutesToDie  int
		minutesToTest int
		want          int
	}{
		{1000, 15, 60, 5},
		{4, 15, 15, 2},
		{4, 15, 30, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.buckets), func(t *testing.T) {
			require.Equal(t, tc.want, poorPigs(tc.buckets, tc.minutesToDie, tc.minutesToTest))
		})
	}
}

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	// With one pig, and one experiment, we can cover 2 buckets.
	// With one pig, and two experiments, we can cover 3 buckets.
	// With two pigs, and one experiment, we can cover 4 buckets.
	// For each bucket (except one), at least one pig must try it. If two or more
	// pigs try the same bucket, then all pigs will die, which also gives away
	// which bucket is poisonous.
	// So, for each bucket, assign a binary non-zero number. Each pig represents
	// one position in the number. So, bucket 1 can be 001, bucket two, 011, three
	// 010, and so on. This means that the number of buckets we can test with
	// x pigs and one experiment is 2^x.
	//
	// What about more than one experiment?
	// Whenever a pig dies using the method above (at least one pig covering each
	// bucket), then the buckets position becomes immediately known and the trial
	// ends. But in the event that all pigs survive, they will need to try once
	// again in the next round.
	//
	// With this definition, each pig has two attempts and three states:
	// Either the pig never drinks a certain bucket, or the pig drinks the bucket
	// and dies in the first attempt, or the pig drinks and dies from the bucket
	// in the third attempt.
	//
	// This means that there are now 3 states for each pig instead of 2.
	//
	// So the solution to the problem is finding the minimum x such that
	// (experiments+1)^x >= buckets
	//
	// This gives us x*ln(experiments+1) >= ln(buckets)
	// x >= ln(buckets) / ln(experiments+1)
	// x = ceil(ln(buckets) / ln(experiments+1))
	experiments := minutesToTest / minutesToDie
	return int(math.Ceil(math.Log2(float64(buckets)) / math.Log2(float64(experiments+1))))
}
