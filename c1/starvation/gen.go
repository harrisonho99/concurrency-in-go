package starvation

import (
	"concurrency_in_go/c1/fmt"
	"sync"
	"time"
)

func Gen() {
	var wg sync.WaitGroup
	var sharedLock sync.Mutex
	const runtime = time.Second

	greedyWorker := func() {
		defer wg.Done()
		var count int

		for begin := time.Now(); time.Since(begin) < runtime; {
			sharedLock.Lock()
			time.Sleep(time.Nanosecond * 3)
			sharedLock.Unlock()
			count++
		}

		fmt.Printf("Greedy worker was able to execute %v work loops.\n", count)
	}
	politeWorker := func() {
		defer wg.Done()
		var count int

		for begin := time.Now(); time.Since(begin) < runtime; {
			sharedLock.Lock()
			time.Sleep(time.Nanosecond * 1)
			sharedLock.Unlock()

			sharedLock.Lock()
			time.Sleep(time.Nanosecond * 1)
			sharedLock.Unlock()

			sharedLock.Lock()
			time.Sleep(time.Nanosecond * 1)
			sharedLock.Unlock()

			count++
		}

		fmt.Printf("Polite worker was able to execute %v work loops.\n", count)
	}

	wg.Add(2)
	go greedyWorker()
	go politeWorker()
	wg.Wait()
}
