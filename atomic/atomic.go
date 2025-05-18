package atomic

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func CountWithAtomic() {
	a := &atomic.Int32{}
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(1000)

	for i := 0; i < 1000; i++ {
		go func() {
			defer waitGroup.Done()
			a.Add(1)
		}()
	}

	waitGroup.Wait()
	fmt.Println(a.Load())
}

func Count() {
	a := 0
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(1000)

	for i := 0; i < 1000; i++ {
		go func() {
			defer waitGroup.Done()
			a += 1
		}()
	}

	waitGroup.Wait()
	fmt.Println(a)
}
