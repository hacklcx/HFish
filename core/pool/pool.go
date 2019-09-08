package pool

import (
	"sync"
	"github.com/panjf2000/ants"
)

func New(size int) (sync.WaitGroup, *ants.Pool) {
	wg := sync.WaitGroup{}
	pool, _ := ants.NewPool(size)
	return wg, pool
}
