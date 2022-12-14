package deadlock

import (
	"concurrency_in_go/c1/fmt"
	"sync"
	"time"
)

type value struct {
	mu    sync.Mutex
	value int
}

var wg sync.WaitGroup

func printSum(v1, v2 *value) {
	defer wg.Done()
	v1.mu.Lock()
	time.Sleep(time.Second * 1)
	v2.mu.Lock()
	fmt.Printf("sum is: %d\n", v1.value+v2.value)
	v1.mu.Unlock()
	v2.mu.Unlock()
}

// see image 1.1
func Gen() {

	var a, b value
	wg.Add(2)
	// both a and b are lock each other
	go printSum(&a, &b)
	go printSum(&b, &a)

	wg.Wait()
}
