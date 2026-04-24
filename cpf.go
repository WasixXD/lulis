package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
	"time"
)

var pow10table = []int{
	1e00,
	1e01,
	1e02,
	1e03,
	1e04,
	1e05,
	1e06,
	1e07,
	1e08,
	1e09,
	1e10,
}

func div10(n uint64) uint64 {
	return (n * 0xCCCCCCCCCCCCCCCD) >> 67
}

func mod10(n uint64) uint64 {
	q := div10(n)
	return n - q*10
}

func genCpfs(start int, end int) {
	var local int

	for n := start; n <= end; n++ {
		value := 0
		sum := 0

		for i := 0; i < 9; i++ {
			num := (n / pow10table[8-i]) % 10
			value += (10 - i) * num
			sum += num
		}

		digit1 := (11 - (value % 11)) % 10
		value += sum + digit1*2
		digit2 := (11 - (value % 11)) % 10

		local += digit1 + digit2

		// fmt.Printf("%d%d%d\n", n, digit1, digit2)
	}
}

func main() {
	f, _ := os.Create("cpu.prof")
	defer f.Close()

	cpfsTotal := int(10e9)
	wait := sync.WaitGroup{}
	nCpus := runtime.NumCPU()
	amount := cpfsTotal / nCpus

	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal(err)
	}
	defer pprof.StopCPUProfile()

	start := time.Now()
	for n := 0; n < nCpus; n++ {
		// division of labor
		start := (n + 1) * amount
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
}
