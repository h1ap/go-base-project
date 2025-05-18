package mutex

import (
	"fmt"
	"sync"
)

func CountWithMutex() {
	a := 0
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(1000)
	lock := sync.Mutex{}

	for i := 0; i < 1000; i++ {
		go func() {
			defer waitGroup.Done()
			lock.Lock()
			defer lock.Unlock()
			a += 1
		}()
	}

	waitGroup.Wait()
	fmt.Println(a)
}
