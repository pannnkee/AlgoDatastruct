package ratelimit

import (
	"sync/atomic"
	"time"
)

type RateLimiter struct {
	rate uint64
	allowance uint64
	max uint64
	unit uint64
	lastCheck uint64
}

func New(rate int, pre time.Duration) *RateLimiter {
	nano := uint64(pre)
	if nano < 1 {
		nano = uint64(time.Second)
	}
	if rate < 1 {
		rate = 1
	}
	return &RateLimiter{
		rate:      uint64(rate),
		allowance: uint64(rate) * nano,
		max:       uint64(rate) * nano,
		unit:      nano,
		lastCheck: unixNano(),
	}
}

func (this *RateLimiter) Limit() bool {
	now := unixNano()
	// 计算上一次调用到现在过了多少纳秒
	passed := now - atomic.SwapUint64(&this.lastCheck, now)
	rate := atomic.LoadUint64(&this.rate)
	current := atomic.AddUint64(&this.allowance, passed*rate)

	if max := atomic.LoadUint64(&this.max); current > max {
		atomic.AddUint64(&this.allowance, max-current)
		current = max
	}

	if current < this.unit {
		return true
	}
	atomic.AddUint64(&this.allowance, -this.unit)
	return  false
}

// unixNano 当前时间（纳秒）
func unixNano() uint64 {
	return uint64(time.Now().UnixNano())
}