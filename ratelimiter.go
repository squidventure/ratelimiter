package ratelimiter

import (
	"net/http"
	"strings"
	"sync/atomic"

	"github.com/squidventure/set"
)

var (
	TheRateLimiter RateLimiter = &BasicRateLimiter{}

	MaxConnections            int64 = 100
	prefixesBypassRateLimiter       = set.NewSet[string](true)
	suffixesBypassRateLimiter       = set.NewSet[string](true)
)

func RegisterBypassPrefix(prefix string) {
	prefixesBypassRateLimiter.Add(prefix)
}

func RegisterBypassSuffix(suffix string) {
	suffixesBypassRateLimiter.Add(suffix)
}

func PathShouldBypassRateLimiter(path string) bool {
	for _, prefix := range prefixesBypassRateLimiter.Slice() {
		if strings.HasPrefix(path, prefix) {
			return true
		}
	}
	for _, suffix := range suffixesBypassRateLimiter.Slice() {
		if strings.HasSuffix(path, suffix) {
			return true
		}
	}
	return false
}

type RateLimiter interface {
	Arrive() bool
	Depart()
	Count() int64
}

type BasicRateLimiter struct {
	n atomic.Int64
}

func (r *BasicRateLimiter) Arrive() bool {
	n := r.n.Add(1)
	return n <= MaxConnections
}

func (r *BasicRateLimiter) Depart() {
	r.n.Add(-1)
}

func (r *BasicRateLimiter) Count() int64 {
	return r.n.Load()
}

func Limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		shouldArrive := TheRateLimiter.Arrive()
		if shouldArrive || PathShouldBypassRateLimiter(r.URL.Path) {
			next.ServeHTTP(w, r)
			return
		}
		http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
	})
}
