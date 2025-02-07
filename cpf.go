package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

func strings2ints(s []string) []int {
	tmp := make([]int, len(s))

	for i, v := range s {
		num, _ := strconv.Atoi(v)
		tmp[i] = num
	}
	return tmp
}

var cache map[int]int

func ints2slice(s int) []int {
	numberLen := 9
	tmp := make([]int, numberLen)
	for c := 0; c < numberLen && s > 0; c++ {
		tmp[numberLen-1-c] = s % 10
		s = int(s / 10)
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

func ints2byte(s []int) []byte {
	tmp := make([]byte, len(s))
	for i, v := range s {
		tmp[i] = byte(v + '0')
	}
	return tmp
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
		nums = append(nums, digit1, digit2)

		comm <- nums
	}

}

func minimun(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func main() {
	// cpfsTotal := int64(10e11)
	cpfsTotal := int64(10e8) // 0.1%
	wait := sync.WaitGroup{}
	nCpus := int64(12)

	c := make(chan []int)

	go func() {
		for v := range c {
			fmt.Println(v)
		}
	}()

	for n := range nCpus {
		// division of labor
		amount := cpfsTotal / nCpus
		start := (n + 1) * amount
		end := minimun(start+amount, cpfsTotal)
		cpfsTotal -= amount

		wait.Add(1)
		go func(start int, end int) {
			defer wait.Done()
			genCpfs(start, end, c)
		}(int(start), int(end))
	}

	wait.Wait()
	os.Exit(0)
}
