package algorithm

import "github.com/ms56bc/rate-limter/pkg"

type FixedWindowRateLimiter struct {
}

func NewFixedWindow() *FixedWindowRateLimiter {
	return &FixedWindowRateLimiter{}
}

func (sw *FixedWindowRateLimiter) HasLimitReached(id pkg.EntityId) bool {
	return false

}
