package main

import (
	"fmt"
	"github.com/kofalt/go-memoize"
	"time"
)

func main() {
	// Any expensive call that you wish to cache
	expensive := func() (any, error) {
		time.Sleep(3 * time.Second)
		return "some data", nil
	}

	// Cache expensive calls in memory for 90 seconds, purging old entries every 10 minutes.
	cache := memoize.NewMemoizer(90*time.Second, 10*time.Minute)

	// This will call the expensive func
	result, err, cached := cache.Memoize("key1", expensive)
	fmt.Println(result, err, cached)
	// This will be cached
	// result, err, cached = cache.Memoize("key1", expensive)

	// This uses a new cache key, so expensive is called again
	// result, err, cached = cache.Memoize("key2", expensive)
}
