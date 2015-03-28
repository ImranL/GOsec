package main

import (
	"fmt"
	"gosec/lib/pmtu"
	"runtime"
)

func main() {
	// Go MP!
	runtime.GOMAXPROCS(runtime.NumCPU())

	fmt.Println("Hello World v2!")

	// test PMTU library:
	pmtu.PmtuTestHarness()
}
