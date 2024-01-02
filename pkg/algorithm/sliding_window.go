package algorithm

import "github.com/ms56bc/rate-limter/pkg"

type SlidingWindowRateLimiter struct {
}

func NewSlidingWindow() *SlidingWindowRateLimiter {
	return &SlidingWindowRateLimiter{}
}

func (sw *SlidingWindowRateLimiter) HasLimitReached(entityId pkg.EntityId) bool {
	return false
}
