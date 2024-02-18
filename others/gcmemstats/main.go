package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/trace"
)

func main() {
	// Create a trace file.
	traceFile, err := os.Create("trace.out")
	if err != nil {
		log.Fatal(err)
	}
	defer traceFile.Close()

	// Start tracing.
	if err := trace.Start(traceFile); err != nil {
		log.Fatal(err)
	}
	defer trace.Stop()

	// Allocate a large amount of memory
	ls := make([]int, 1000000)
	if len(ls) == 0 {
		fmt.Println("short ls")
	}

	// Force a garbage collection
	runtime.GC()
	var memStats runtime.MemStats

	runtime.ReadMemStats(&memStats)

	fmt.Printf("Total allocated memory (in bytes): %d\n", memStats.Alloc)
	fmt.Printf("Heap memory (in bytes): %d\n", memStats.HeapAlloc)
	fmt.Printf("Number of garbage collections: %d\n", memStats.NumGC)
}
