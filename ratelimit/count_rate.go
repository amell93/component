package ratelimit

import (
	"errors"
	"sync"
	"time"
)

type CountRate struct {
	rate  int           //阀值
	begin time.Time     //计数开始时间
	cycle time.Duration //计数周期
	count int           //收到的请求数
	lock  sync.Mutex    //锁
}

var RateOrCycleErr = errors.New("rate and cycle must greater than zero")

// NewCountRate returns a CountRate by the given value rate and cycle.
// it returns an RateOrCycleErr if the rate or cycle not greater than zero.
func NewCountRate(rate int, cycle time.Duration) (*CountRate, error) {

	if rate <= 0 || cycle <= 0 {
		return nil, RateOrCycleErr
	}

	cr := CountRate{
		rate:  rate,
		begin: time.Now(),
		cycle: cycle,
	}

	return &cr, nil

}

// Take returns the result of add count.
// if returns true when add the count success
// else returns false.
func (cRate *CountRate) Take() bool {
	cRate.lock.Lock()
	defer cRate.lock.Unlock()

	// 判断收到请求数是否达到阀值
	if cRate.count == cRate.rate-1 {
		now := time.Now()
		// 达到阀值后，判断是否是请求周期内
		if now.Sub(cRate.begin) >= cRate.cycle {
			cRate.Reset(now)
			return true
		}
		return false
	} else {
		cRate.count++
		return true
	}
}

// Set set the CountRate by given value rate and cycle
func (cRate *CountRate) Set(rate int, cycle time.Duration) error {

	if rate <= 0 || cycle <= 0 {
		return RateOrCycleErr
	}

	cRate.lock.Lock()
	defer cRate.lock.Unlock()

	cRate.rate = rate
	cRate.begin = time.Now()
	cRate.cycle = cycle
	cRate.count = 0

	return nil
}

// Reset reset the begin time and count
func (cRate *CountRate) Reset(begin time.Time) {
	cRate.lock.Lock()
	defer cRate.lock.Unlock()

	cRate.begin = begin
	cRate.count = 0
}
