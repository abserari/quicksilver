package timingwheel

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimingWheel(t *testing.T) {
	var deviation int64
	var num int64 = 100
	sleeptime := time.Millisecond * 20000

	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			start := time.Now().UnixNano()
			<-After(sleeptime)
			end := time.Now().UnixNano()
			duration := end - start - int64(sleeptime)
			if duration < 0 {
				deviation -= duration
			} else {
				deviation += duration
			}
			wg.Done()
		}(i)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Wait()
	fmt.Println(float64(deviation) / float64(num*int64(sleeptime)))

}
