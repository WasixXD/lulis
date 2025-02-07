package main

import (
	"fmt"
	"os"
	"runtime"
	"sync"
)

func ints2slice(s int) []int {
	numberLen := 11
	tmp := make([]int, numberLen)
	for c := 0; c < numberLen && s > 0; c++ {
		tmp[c] = s % 10
		s /= 10
	}
	return tmp
}

func sum(s []int) (sum int) {
	sum = 0

	for _, v := range s {
		sum += v
	}
	return
}

func genCpfs(start int, end int, comm chan []int) {
	for n := start; n <= end; n++ {
		nums := ints2slice(n)

		value := 0

		for i, n := range nums {
			value += (10 - i) * n
		}

		digit1 := (11 - (value % 11)) % 10
		value2 := value + (sum(nums) + (digit1 * 2))
		digit2 := (11 - (value2 % 11)) % 10

		nums[9] = digit1
		nums[10] = digit2

		comm <- nums
	}

}

func minimun(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// cpfsTotal := int64(10e11)
	cpfsTotal := int(10e8) // 0.1%
	wait := sync.WaitGroup{}
	nCpus := runtime.NumCPU() * 2
	amount := cpfsTotal / nCpus

	c := make(chan []int, nCpus)

	go func() {
		for v := range c {
			fmt.Println(v)
		}
	}()

	for n := 0; n < nCpus; n++ {
		// division of labor
		start := n * amount
		end := minimun(start+amount, cpfsTotal)

		wait.Add(1)
		go func(start int, end int) {
			defer wait.Done()
			genCpfs(start, end, c)
		}(start, end)
	}
	wait.Wait()
	os.Exit(0)
}
