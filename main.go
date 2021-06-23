package main

import (
	"fmt"
	"github/com/xzx/zkl/zllock"
	"sync"
	"time"
)

var n = 0

func main() {
	zl := new(zllock.ZKLock)
	err := zl.Init("/lockTest", []string{"127.0.0.1:2181"})
	if err != nil {
		panic(err)
	}
	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			lockUnlock(test)
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(n)
	fmt.Println("Stop...")
}

func lockUnlock(f func()) {
	zl := new(zllock.ZKLock)
	err := zl.CreateLock("/lockTest", 10*time.Second, []string{"127.0.0.1:2181"})
	if err != nil {
		panic(err)
	}
	err = zl.AttempLock()
	if err != nil {
		panic(err)
	}

	f()

	err = zl.Unlock()
	if err != nil {
		panic(err)
	}
}

func test() {
	time.Sleep(1 * time.Second)
	n = n + 1
}
