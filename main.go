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
	tokens       int //it tell how many tokens are currently in the bucket
	capacity     int //it tells max tokens that a bucket can hold
	refillRate   int //how fast tokens are added 
	lastRefillTs time.Time //last time we added token
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
		capacity:     5,
		refillRate:   3,
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


/*Output
Request 1 allowed at 20:20:34.675
Request 2 allowed at 20:20:34.676
Request 3 allowed at 20:20:34.676
Request 4 allowed at 20:20:34.676
Request 5 allowed at 20:20:34.676
Request 6 blocked at 20:20:34.676
Request 7 blocked at 20:20:34.677
Request 8 blocked at 20:20:34.677
Request 9 blocked at 20:20:34.677
Request 10 blocked at 20:20:34.677
*/