package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"time"
)

func ints2slice(s int) []int {
	tmp := make([]int, 11)
	numberLen := 9
	for c := 0; c < numberLen; c++ {
		tmp[numberLen-c-1] = s % 10
		s /= 10
	}
	return tmp
}

func genCpfs(start int, end int) {
	for n := start; n <= end; n++ {
		nums := make([]int, 11)

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

		nums[9] = digit1
		nums[10] = digit2

	}

}

func minimun(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	cpfsTotal := int(10e10)
	wait := sync.WaitGroup{}
	nCpus := runtime.NumCPU()
	amount := cpfsTotal / nCpus

	start := time.Now()
	for n := 0; n < nCpus; n++ {
		// division of labor
		start := n * amount
		end := minimun(start+amount, cpfsTotal)

		wait.Add(1)
		go func(start int, end int) {
			defer wait.Done()
			genCpfs(start, end)
		}(start, end)
	}

	wait.Wait()
	end := time.Since(start)
	fmt.Printf("Calculated: %d in %v\n", cpfsTotal, end)
	// os.Exit(0)
}
