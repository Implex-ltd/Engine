package ratelimiter

import (
	"sync"
	"time"
)

type Limiter struct {
	Amount int
	Span   time.Duration
	mu     sync.Mutex
	count  int
}
