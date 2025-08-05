package main

import "time"

type TokenBucket struct {
	tokens       int
	capacity     int
	refillRate   int
	lastRefillTs time.Time
}

func min(a, b int)int{
	if a < b{
		return a
	}
	return b
}

func (b *TokenBucket) AllowRequest() bool{
	now := time.Now()
	elapsed := now.Sub(b.lastRefillTs).Seconds()

	tokensToAdd := int(elapsed * float64(b.refillRate))
	if tokensToAdd > 0{
		b.tokens = min(b.capacity, b.tokens + tokensToAdd)
		b.lastRefillTs = now
	}

	if b.tokens > 0{
		b.tokens--
		return true
	}
	return false
}
