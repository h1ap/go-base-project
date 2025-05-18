package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Product struct {
	inventoryRemaining *atomic.Int32
	mutex              *sync.Mutex
	consumerCond       *sync.Cond
	producerCond       *sync.Cond
}

func NewProduct() *Product {
	mutex := &sync.Mutex{}
	inventory := &atomic.Int32{}
	inventory.Store(0)

	return &Product{
		inventoryRemaining: inventory,
		mutex:              mutex,
		consumerCond:       sync.NewCond(mutex),
		producerCond:       sync.NewCond(mutex),
	}
}

func main() {
	consumer := 5
	producer := 3
	maxInventory := int32(10)

	p := NewProduct()

	// 启动消费者
	for i := 0; i < consumer; i++ {
		go func(id int) {
			for {
				p.mutex.Lock()
				for p.inventoryRemaining.Load() <= 0 {
					fmt.Printf("消费者 %d 等待产品...\n", id)
					p.consumerCond.Wait()
				}
				p.inventoryRemaining.Add(-1)
				fmt.Printf("消费者 %d 消费了一个产品，剩余 %d\n", id, p.inventoryRemaining.Load())
				p.producerCond.Signal() // 通知生产者可以生产
				p.mutex.Unlock()

				// 模拟消费时间
				time.Sleep(time.Millisecond * 500)
			}
		}(i)
	}

	// 启动生产者
	for i := 0; i < producer; i++ {
		go func(id int) {
			for {
				p.mutex.Lock()
				for p.inventoryRemaining.Load() >= maxInventory {
					fmt.Printf("生产者 %d 等待空间...\n", id)
					p.producerCond.Wait()
				}
				p.inventoryRemaining.Add(1) // 生产一个产品
				fmt.Printf("生产者 %d 生产了一个产品，现有 %d\n", id, p.inventoryRemaining.Load())
				p.consumerCond.Signal() // 通知消费者可以消费
				p.mutex.Unlock()

				// 模拟生产时间
				time.Sleep(time.Millisecond * 300)
			}
		}(i)
	}

	// 让主线程不退出
	select {}
}
