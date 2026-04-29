package main

import (
	"fmt"
	"lulis/cpf"
	"runtime"
	"sync"
	"time"
)

const MAX_CPFS = 10e8 - 1 // 1_000_000_000 - 1 = 999_999_999
var sink int

func main() {
	numCores := runtime.NumCPU()
	cpfPerBatch := MAX_CPFS / numCores
	wait := sync.WaitGroup{}
	now := time.Now()
	for i := range numCores {
		start := i * cpfPerBatch
		end := start + cpfPerBatch - 1
		if i == numCores-1 {
			end = MAX_CPFS
		}

		wait.Add(1)

		go func(start, end int) {
			defer wait.Done()
			cpf.CalcRange(start, end)
		}(start, end)

	}
	wait.Wait()
	fmt.Println("Multiple threads ", time.Since(now))

	start := time.Now()
	sink = cpf.CalcRange(0, MAX_CPFS)
	fmt.Println("SingleThreaded", time.Since(start))

	now = time.Now()
	sink = cpf.GenCpfs(0, MAX_CPFS)
	fmt.Println("Old implementation", time.Since(start))

}
