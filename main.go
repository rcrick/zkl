package main

import (
	"fmt"
	"github/com/xzx/zkl/zllock"
	"sync"
)

var n = 0

func main() {
	zl := new(zllock.ZKLock)
	err := zl.Init("/lockTest", []string{"127.0.0.1:2181"})
	if err != nil {
		panic(err)
	}
	wg := sync.WaitGroup{}
	wg.Add(80)
	for i := 0; i < 80; i++ {
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
	err := zl.CreateLock("/lockTest", []string{"127.0.0.1:2181"})
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
	// time.Sleep(time.Millisecond * 10)
	n = n + 1
}
