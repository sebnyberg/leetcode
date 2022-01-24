package trifles

import (
	"sync"
	"time"
)

type rateLimiter struct {
	ringSize      int
	minDuration   time.Duration
	resetDuration time.Duration
	clients       map[string]*client
	mtx           sync.RWMutex
}

func newRateLimiter(
	ringSize int,
	minDuration time.Duration,
	resetDuration time.Duration,
) *rateLimiter {
	return &rateLimiter{
		ringSize:      ringSize,
		minDuration:   minDuration,
		resetDuration: resetDuration,
		clients:       map[string]*client{},
	}
}

type client struct {
	mtx      sync.RWMutex
	requests []time.Time
	reset    func()
	idx      int
}

func (rl *rateLimiter) registerClient(clientID string) {
	timer := time.AfterFunc(rl.resetDuration, func() {
		rl.mtx.Lock()
		defer rl.mtx.Unlock()
		delete(rl.clients, clientID)
	})

	cl := &client{
		requests: make([]time.Time, rl.ringSize),
		reset:    func() { timer.Reset(rl.resetDuration) },
	}
	rl.mtx.Lock()
	defer rl.mtx.Unlock()
	rl.clients[clientID] = cl
}

func (rl *rateLimiter) call(clientID string) bool {
	rl.mtx.RLock()
	if _, exists := rl.clients[clientID]; !exists {
		rl.mtx.RUnlock()
		rl.registerClient(clientID)
	} else {
		rl.mtx.RUnlock()
	}
	cl := rl.clients[clientID]
	cl.mtx.Lock()
	defer cl.mtx.Unlock()
	now := time.Now()
	if now.Sub(cl.requests[cl.idx]) < rl.minDuration {
		return false
	}
	cl.requests[cl.idx] = now
	cl.idx = (cl.idx + 1) % rl.ringSize
	cl.reset()
	return true
}
