package algorithm

import "github.com/ms56bc/rate-limter/pkg"

type TokenBucketRateLimiter struct {
}

func NewTokenBucket() *TokenBucketRateLimiter {
	return &TokenBucketRateLimiter{}
}

func (sw *TokenBucketRateLimiter) HasLimitReached(entityId pkg.EntityId) bool {
	return false
}
