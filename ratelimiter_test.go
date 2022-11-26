package ratelimiter

import "testing"

var (
	TheRateLimiter = &BasicRateLimiter{}
)

func TestArrive(t *testing.T) {
	for i := int64(0); i < MaxConnections; i++ {
		if shouldArrive := TheRateLimiter.Arrive(); !shouldArrive {
			t.Fatalf("error: expected successful arrive but got false")
		}
	}
}
