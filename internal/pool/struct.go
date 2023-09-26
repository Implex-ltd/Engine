package pool

import (
	"sync"

	"github.com/Implex-ltd/engine/internal/browser"
)

type Pool struct {
	Size, Curr int
	Mut        *sync.Mutex
	Workers    []*browser.Instance
}
