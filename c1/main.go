package main

import (
	"concurrency_in_go/c1/fmt"
	"concurrency_in_go/c1/starvation"
)

func main() {
	// atomic.CriticalSectionAccess()
	// deadlock.Gen()
	// livelock.Gen()
	starvation.Gen()
}

func init() {
	fmt.SetColorProfile(fmt.NoticeColor)
}
