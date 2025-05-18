package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Runner struct {
	Name string
}

func (r *Runner) Run(cwg, wwg *sync.WaitGroup) {
	defer cwg.Done()
	wwg.Wait()

	start := time.Now()
	fmt.Println(r.Name, "started at", start)
	rand.NewSource(time.Now().UnixNano())
	time.Sleep(time.Duration(rand.Intn(1000)%10) * time.Second)
	finish := time.Now()
	fmt.Println(r.Name, "finished at", finish.Sub(start))
}

func main() {
	cwg := sync.WaitGroup{}
	wwg := sync.WaitGroup{}
	wwg.Add(1)
	cwg.Add(10)
	for i := 0; i < 10; i++ {
		r := &Runner{fmt.Sprintf("Runner %d", i)}
		go r.Run(&cwg, &wwg)
	}

	wwg.Done()
	cwg.Wait()
}
