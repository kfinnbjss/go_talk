package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
)

func memConsumed() uint64 {
	runtime.GC()
	var s runtime.MemStats
	runtime.ReadMemStats(&s)
	return s.Sys
}

func ByteCountSI(b uint64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB",
		float64(b)/float64(div), "kMGTPE"[exp])
}

func getMemUsedByGoroutines(numGoroutines int) uint64 {
	var c <-chan interface{}
	var wg sync.WaitGroup
	noop := func() { wg.Done(); <-c } // <1>

	wg.Add(numGoroutines)
	before := memConsumed() // <3>
	for i := numGoroutines; i > 0; i-- {
		go noop()
	}
	wg.Wait()
	after := memConsumed() // <4>
	return after - before
}

func main() {
	fmt.Println("number_of_goroutines\ttotal_memory\t\tmomery_per_goroutine")
	for i := 1; i < 7; i++ {
		numGoroutines := int(math.Pow10(i))
		mem := getMemUsedByGoroutines(numGoroutines)
		fmt.Printf("%.0e\t\t\t%s  \t\t%.3fkb\n", float64(numGoroutines), ByteCountSI(mem), float64(mem)/float64(numGoroutines)/1000)
	}
}
