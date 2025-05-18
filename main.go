package main

import (
	"go-base-project/atomic"
	"go-base-project/mutex"
)

func main() {
	atomic.Count()
	atomic.CountWithAtomic()
	mutex.CountWithMutex()
}
