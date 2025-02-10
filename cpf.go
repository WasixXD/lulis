package main

import (
	"fmt"
	"math"
	"os"
	"runtime"
	"sync"
	"time"
)

func genCpfs(start int, end int) {
	for n := start; n <= end; n++ {
		value := 0
		sum := 0

		for i := 0; i < 9; i++ {
			num := (n / int(math.Pow10(i))) % 10
			value += (10 - i) * num
			sum += num
		}

		digit1 := (11 - (value % 11)) % 10
		value += (sum + (digit1 * 2))
		digit2 := (11 - (value % 11)) % 10

		// fast concat
		_ = ((n*10+digit1)*10 + digit2)
	}

}

func main() {
	cpfsTotal := int(10e11)
	wait := sync.WaitGroup{}
	nCpus := runtime.NumCPU()
	amount := cpfsTotal / nCpus

	start := time.Now()
	for n := 0; n < nCpus; n++ {
		// division of labor
		start := n * amount
		end := min(start+amount, cpfsTotal)

		wait.Add(1)
		go func(start int, end int) {
			defer wait.Done()
			genCpfs(start, end)
		}(start, end)
	}

	wait.Wait()
	end := time.Since(start)
	fmt.Printf("Calculated: %d in %v\n", cpfsTotal, end)
	os.Exit(0)

}
