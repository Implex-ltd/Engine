package ratelimiter

import (
	"time"
)

func NewLimiter(amount int, span time.Duration) *Limiter {
	return &Limiter{
		Amount: amount,
		Span:   span,
	}
}

func (L *Limiter) Wait() {
	L.mu.Lock()
	defer L.mu.Unlock()

	for L.count >= L.Amount {
		time.Sleep(time.Second)
	}

	L.count++
	go L.release()
}

func (L *Limiter) release() {
	time.Sleep(L.Span)

	L.mu.Lock()
	defer L.mu.Unlock()
	L.count--
}
