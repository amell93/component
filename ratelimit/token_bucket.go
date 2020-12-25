package ratelimit

import (
	"errors"
	"sync"
	"time"
)

type TokenBucket struct {
	rate         int64 //固定的token放入速率, r/s
	capacity     int64 //桶的容量
	tokens       int64 //桶中当前token数量
	lastTokenSec int64 //上次向桶中放令牌的时间的时间戳，单位为秒

	mu sync.Mutex
}

var ParametersErr = errors.New("rate or capacity err")

func NewTokenBucket(rate, capacity int64) (*TokenBucket, error) {
	if rate <= 0 || capacity <= 0 {
		return nil, ParametersErr
	}

	tb := TokenBucket{
		rate:         rate,
		capacity:     capacity,
		lastTokenSec: time.Now().Unix(),
	}

	return &tb, nil
}

func (bucket *TokenBucket) Take() bool {

	now := time.Now().Unix()

	bucket.mu.Lock()
	defer bucket.mu.Unlock()

	bucket.tokens = bucket.tokens + (now-bucket.lastTokenSec)*bucket.rate // 先添加令牌
	if bucket.tokens > bucket.capacity {
		bucket.tokens = bucket.capacity
	}
	bucket.lastTokenSec = now
	if bucket.tokens > 0 {
		// 还有令牌，领取令牌
		bucket.tokens--
		return true
	} else {
		// 没有令牌,则拒绝
		return false
	}
}

func (bucket *TokenBucket) Set(rate, cap int64) error {

	if rate <= 0 || cap <= 0 {
		return ParametersErr
	}

	bucket.mu.Lock()
	defer bucket.mu.Unlock()

	bucket.rate = rate
	bucket.capacity = cap
	bucket.tokens = 0
	bucket.lastTokenSec = time.Now().Unix()

	return nil
}
