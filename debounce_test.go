package debounce_test

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	"github.com/angadn/debounce"
)

func TestDebounce(t *testing.T) {
	var (
		deb debounce.Block
		ctr uint64
	)

	for i := 0; i < 3; i++ {
		for j := 0; j < 10; j++ {
			deb.Do(100*time.Millisecond, func() {
				atomic.AddUint64(&ctr, 1)
			})
		}

		time.Sleep(200 * time.Millisecond)
	}

	c := int(atomic.LoadUint64(&ctr))
	fmt.Printf("ctr = %d\n", c)
	if c != 3 {
		t.Fail()
	}
}
