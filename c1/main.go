package main

import (
	"concurency-in-go/c1/deadlock"
)

func main() {
	// atomic.CriticalSectionAccess()
	deadlock.GenDeadlock()
}
