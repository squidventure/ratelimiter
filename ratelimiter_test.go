package ratelimiter

import (
	"testing"
)

func TestArrive(t *testing.T) {
	for i := int64(0); i < MaxConnections; i++ {
		if shouldArrive := TheRateLimiter.Arrive(); !shouldArrive {
			t.Fatalf("error: expected successful arrive but got false")
		}
	}
	if shouldArrive := TheRateLimiter.Arrive(); shouldArrive {
		t.Fatalf("error: expected rejection from TheRateLimiter.Arrive() but got true")
	}
}

func TestDepart(t *testing.T) {
	for i := int64(0); i <= MaxConnections; i++ {
		TheRateLimiter.Depart()
	}
	want := int64(0)
	if got := TheRateLimiter.Count(); want != got {
		t.Fatalf("error: expected TheRateLimiter to be empty but it has %d entries", got)
	}
}

func TestArriveDepartInter(t *testing.T) {
	for i := 0; i < 100; i++ {
		func() {
			TheRateLimiter.Arrive()
			defer TheRateLimiter.Depart()
		}()
	}
	want := int64(0)
	if got := TheRateLimiter.Count(); want != got {
		t.Fatalf("error: expected empty rate limiter at idle but got %d entries", got)
	}
}
