package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

const (
	CPF_DIGITS = 9
	BATCH_SIZE = 4
)

func calcCpfRange(start, end int) int {
	var digits [CPF_DIGITS]int
	var local int

	for k := start; k <= end; k += BATCH_SIZE {
		var batchCpf [BATCH_SIZE][CPF_DIGITS]int
		for x := 0; x < BATCH_SIZE; x++ {
			// Instead of creating an array every loop
			// we use just one and increments its bits one by one
			// 0 0 0 0 0 0 0 0 1
			// 0 0 0 0 0 0 0 0 2
			// ...
			// 0 0 0 0 0 0 0 1 0
			for i := 8; i >= 0; i-- {
				digits[i]++
				if digits[i] < 10 {
					break
				}
				digits[i] = 0
			}
			batchCpf[x] = digits
		}

		// Holds A+B+C+D+E+F+G+H+I
		baseSum := [BATCH_SIZE]int{}
		sum := [BATCH_SIZE]int{}

		for i, j := 0, 10; i < 9; i, j = i+1, j-1 {
			// Loop unrolling
			digit := batchCpf[0][i]
			sum[0] += digit * j
			baseSum[0] += digit

			digit = batchCpf[1][i]
			sum[1] += digit * j
			baseSum[1] += digit

			digit = batchCpf[2][i]
			sum[2] += digit * j
			baseSum[2] += digit

			digit = batchCpf[3][i]
			sum[3] += digit * j
			baseSum[3] += digit
		}

		var finalDigits [BATCH_SIZE][2]int
		for x := 0; x < BATCH_SIZE; x++ {
			// digit1 is (10A + 9B + 8C + 7D + 6E + 5F + 4G + 3H + 2I) % 11
			rem := sum[x] % 11
			digit1 := remTable[rem]

			// digit2 is 11A + 10B + 9C + 8D + 7E + 6F + 5G + 4H + 3I + 2J
			// but can also be written as:
			// (10A + 9B + 8C + 7D + 6E + 5F + 4G + 3H + 2I) + (A+B+C+D+E+F+G+H+I) + 2 * digit1
			// we save one loop
			sum[x] += baseSum[x] + 2*digit1

			// its has come a time that the division is one of the most
			// expensive operations on this algorithm. Unfortunely even
			// compiler optimization wont save us
			rem = sum[x] % 11
			digit2 := remTable[rem]

			finalDigits[x] = [2]int{digit1, digit2}
			local += digit1 + digit2
		}
	}

	return local
}

func main() {
	f, _ := os.Create("cpu.prof")

	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal(err)
	}

	now := time.Now()
	calcCpfRange(0, 100_000_000)

	pprof.StopCPUProfile()
	f.Close()
	fmt.Println("Finished in: ", time.Since(now))
}
