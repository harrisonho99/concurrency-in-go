package atomic

import (
	"concurency-in-go/c1/fmt"
	"sync"
)

func CriticalSectionAccess() {
	var memroryAccess sync.Mutex
	var value int

	go func() {
		memroryAccess.Lock()
		value++
		memroryAccess.Unlock()

	}()

	memroryAccess.Lock()
	fmt.Printf("value is: %d\n", value)
	memroryAccess.Unlock()
}
