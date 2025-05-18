package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func TestDefChannelWithRange(wg *sync.WaitGroup) {

	waitConsumerGroup := sync.WaitGroup{}
	waitConsumerGroup.Add(5)

	ch := make(chan int)
	for i := 0; i < 5; i++ {
		go func(i int) {
			// 下面两行不是原子操作，所以可能出现receive 3再send 3的情况
			ch <- i
			fmt.Println(i, "send", i)
		}(i)
	}

	go func() {
		defer wg.Done()
		defer close(ch)
		defer waitConsumerGroup.Wait()
	}()

	for c := range ch {
		go func(cc int) {
			defer waitConsumerGroup.Done()
			fmt.Println("receive", cc)
		}(c)
	}
}

func TestDefChannel(wg *sync.WaitGroup) {

	waitConsumerGroup := sync.WaitGroup{}
	waitConsumerGroup.Add(5)

	ch := make(chan int)
	for i := 0; i < 5; i++ {
		go func(i int) {
			// 下面两行不是原子操作，所以可能出现receive 3再send 3的情况
			ch <- i
			fmt.Println(i, "send", i)
		}(i)
	}

	for i := 0; i < 5; i++ {
		go func(ii int) {
			defer waitConsumerGroup.Done()
			c, ok := <-ch
			if !ok {
				return
			}
			fmt.Println(ii, "receive", c)

		}(i)
	}

	defer wg.Done()
	defer close(ch)
	defer waitConsumerGroup.Wait()
}

func TestDefChannelWithSelect(wg *sync.WaitGroup) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	close(ch1)
	close(ch2)
	// 随机执行case，如果都没有数据，则执行default
	select {
	case <-ch1:
		fmt.Println("ch1")
	case <-ch2:
		fmt.Println("ch2")
	default:
		fmt.Println("default")
	}

	wg.Done()
}

func TestDefChannelWithDirection() {
	var inputChan = make(chan<- int)
	var outputChan = make(<-chan int)

	inputChan <- 1

	o := <-outputChan
	fmt.Println("outputChan", o)
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func TestDefChannelWithPrime() {
	size := runtime.GOMAXPROCS(0) * 2
	ch := make(chan int, 512)

	wg := sync.WaitGroup{}
	wg.Add(size)

	for i := 0; i < size; i++ {
		go func() {
			defer wg.Done()
			for i := range ch {
				if isPrime(i) {
					fmt.Println(i)
				}
			}
		}()
	}

	for i := 2; i <= 20000000; i++ {
		ch <- i
	}

	close(ch)
	wg.Wait()
}

func main() {
	// wg := sync.WaitGroup{}
	// wg.Add(1)
	// go TestDefChannel(&wg)
	// go TestDefChannelWithRange(&wg)
	// go TestDefChannelWithSelect(&wg)
	// wg.Wait()

	///////////////////////////////////////////////////////////////

	start := time.Now()
	TestDefChannelWithPrime()
	fmt.Println("time", time.Since(start))
}
