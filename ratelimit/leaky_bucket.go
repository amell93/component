package ratelimit

import (
	"math"
	"sync"
	"time"
)

type LeakyBucket struct {
	rate       float64 // the outflow rate
	capacity   float64 // capacity of bucket
	water      float64 // current water remaining
	lastLeakMs int64   // the milliseconds at most recently outflow
	mu         sync.Mutex
}

func NewLeakyBucket(rate, capacity float64) (*LeakyBucket, error) {
	if rate <= 0 || capacity <= 0 {
		return nil, ParametersErr
	}

	lb := LeakyBucket{
		rate:       rate,
		capacity:   capacity,
		water:      0,
		lastLeakMs: time.Now().UnixNano() / 1e6,
	}

	return &lb, nil
}

func (leaky *LeakyBucket) Take() bool {

	now := time.Now().UnixNano() / 1e6
	leaky.mu.Lock()
	defer leaky.mu.Unlock()
	// calculate the amount of water remaining
	leakyWater := leaky.water - (float64(now-leaky.lastLeakMs) * leaky.rate / 1e3)
	leaky.water = math.Max(0, leakyWater)
	leaky.lastLeakMs = now
	if leaky.water+1 <= leaky.capacity {
		leaky.water++
		return true
	} else {
		return false
	}
}

func (leaky *LeakyBucket) Set(rate, capacity float64) error {
	if rate <= 0 || capacity <= 0 {
		return ParametersErr
	}

	leaky.mu.Lock()
	defer leaky.mu.Unlock()

	leaky.rate = rate
	leaky.capacity = capacity
	leaky.water = 0
	leaky.lastLeakMs = time.Now().UnixNano() / 1e6

	return nil
}
