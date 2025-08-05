/*This is a sample working of Rate Limiter which show the working of Token Bucket Algorithm
Where the Algorithm says 
1. Initialize the token bucket with fixed number of tokens.
2. For each request remove a token from the bucket.
3. If there are no tokens left in the bucket  reject the request.
4. Add tokens to the bucket at a fixed rate.*/
package main

import (
	"fmt"
	"time"
)

type TokenBucket struct {
	tokens       int
	capacity     int
	refillRate   int
	lastRefillTs time.Time
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*This method will check if a request can be allowed and also
calculate how many tokens to add */
func (b *TokenBucket) AllowRequest() bool {
	now := time.Now()
	elapsed := now.Sub(b.lastRefillTs).Seconds()

	tokensToAdd := int(elapsed * float64(b.refillRate))
	if tokensToAdd > 0 {
		b.tokens = min(b.capacity, b.tokens+tokensToAdd)
		b.lastRefillTs = now
	}

	if b.tokens > 0 {
		b.tokens--
		return true
	}
	return false
}

func main() {
	bucket := &TokenBucket{
		tokens:       5,
		capacity:     3,
		refillRate:   2,
		lastRefillTs: time.Now(),
	}

	for i := 1; i <= 10; i++{
		allowed := bucket.AllowRequest()

		if allowed{
			fmt.Printf("Request %d allowed at %s\n", i, time.Now().Format("15:04:05.000"))
		}else{
			fmt.Printf("Request %d blocked at %s\n", i, time.Now().Format("15:04:05.000"))
		}
	}
}
