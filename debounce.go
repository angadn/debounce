package debounce

import (
	"sync"
	"time"
)

type Block struct {
	lock  sync.Mutex
	timer *time.Timer
}

func (block *Block) Do(duration time.Duration, f func()) {
	block.lock.Lock()
	defer block.lock.Unlock()

	if block.timer != nil {
		stop := block.timer.Stop()
		if !stop {
			block.timer = nil
			return
		}
	}

	block.timer = time.AfterFunc(duration, func() {
		block.lock.Lock()
		defer block.lock.Unlock()
		f()
		block.timer = nil
	})
}
