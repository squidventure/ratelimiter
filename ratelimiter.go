package ratelimiter

import "sync/atomic"

var (
	MaxConnections int64 = 100
)

type RateLimiter interface {
	Arrive() bool
	Depart()
	Count() int64
}

type BasicRateLimiter struct {
	n atomic.Int64
}

func (r BasicRateLimiter) Arrive() bool {
	n := r.n.Add(1)
	return n <= MaxConnections
}

func (r BasicRateLimiter) Depart() {
	r.n.Add(-1)
}

func (r BasicRateLimiter) Count() int64 {
	return r.n.Load()
}
