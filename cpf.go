package main

import (
	"runtime"
	"sync"
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
		nums := ints2slice(n)

		value := 0
		sum := 0

		for i, n := range nums[:9] {
			value += (10 - i) * n
			sum += n
		}

		digit1 := (11 - (value % 11)) % 10
		value += (sum + (digit1 * 2))
		digit2 := (11 - (value % 11)) % 10

		nums[9] = digit1
		nums[10] = digit2

		// comm <- nums
	}

}

func minimun(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	cpfsTotal := int(10e11)
	wait := sync.WaitGroup{}
	nCpus := runtime.NumCPU()
	amount := cpfsTotal / nCpus

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
	// os.Exit(0)
}
