package trifles

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestRateLimiter(t *testing.T) {
	rl := newRateLimiter(10, time.Millisecond*1, time.Millisecond*10)
	// Make 10 requests
	clientID := "a"
	for i := 0; i < 10; i++ {
		require.True(t, rl.call(clientID))
	}

	// 11'th call should be blocked
	require.False(t, rl.call(clientID))

	// Wait until 2 ms has passed
	time.Sleep(2 * time.Millisecond)

	// Should be OK now
	require.True(t, rl.call(clientID))

	// Wait 120 ms, check that the client was cleared out from the map
	time.Sleep(20 * time.Millisecond)
	require.Len(t, rl.clients, 0)
}
