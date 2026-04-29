package cpf

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
	"time"
)

const (
	CPF_DIGITS = 9
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

var remTable = [11]int{
	0:  0,
	1:  0,
	2:  9,
	3:  8,
	4:  7,
	5:  6,
	6:  5,
	7:  4,
	8:  3,
	9:  2,
	10: 1,
}

// its slower but easier to debug
func DebugCPF(start, end int) int {
	const BATCH_SIZE = 4
	var digits [CPF_DIGITS]int
	var local int

	for k := start; k < end; k += BATCH_SIZE {
		var batchCpf [BATCH_SIZE][CPF_DIGITS]int
		for x := 0; x < BATCH_SIZE && k+x < end; x++ {
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

		for x := 0; x < BATCH_SIZE; x++ {
			// digit1 is (10A + 9B + 8C + 7D + 6E + 5F + 4G + 3H + 2I) % 11
			rem := sum[x] % 11
			digit1 := remTable[rem]

			// digit2 is 11A + 10B + 9C + 8D + 7E + 6F + 5G + 4H + 3I + 2J
			// but can also be written as:
			// (10A + 9B + 8C + 7D + 6E + 5F + 4G + 3H + 2I) + (A+B+C+D+E+F+G+H+I) + 2 * digit1
			// we save one loop
			// tmp := sum[x]
			sum[x] += baseSum[x] + 2*digit1

			// its has come a time that the division is one of the most
			// expensive operations on this algorithm. Unfortunely even
			// compiler optimization wont save us
			rem = sum[x] % 11
			digit2 := remTable[rem]

			// fmt.Println(batchCpf[x], digit1, digit2, "baseSum", baseSum[x], "sum1", tmp, "sum2", sum[x])
			local += digit1 + digit2
		}
		// if k == 40 {
		// 	break
		// }
	}

	return local
}

func intToDigits(start int, digits *[CPF_DIGITS]int) {
	for i := 0; i < 9; i++ {
		num := (start / pow10table[8-i]) % 10
		digits[i] = num
	}
}

func increaseDigit(digits *[CPF_DIGITS]int) {
	for i := 8; i >= 0; i-- {
		digits[i]++
		if digits[i] < 10 {
			break
		}
		digits[i] = 0
	}
}

func calcBothSum(digits *[CPF_DIGITS]int) (int, int) {
	baseSum := 0
	sum := 0

	for i, j := 0, 10; i < CPF_DIGITS; i, j = i+1, j-1 {
		digit := digits[i]
		baseSum += digit
		sum += digit * j
	}

	return baseSum, sum
}

func calc2Digits(baseSum, sum int) (int, int) {
	rem := sum % 11
	d1 := remTable[rem]

	// 					 v Fun *2 multiplication. No really improvement
	sum += baseSum + (d1 << 1)

	rem = sum % 11
	d2 := remTable[rem]

	return d1, d2
}

// CalcRange
func CalcRange(start, end int) int {
	const BatchSize = 10
	var digits [CPF_DIGITS]int
	var local int

	intToDigits(start, &digits)
	// first batch lets calc the first 9
	for current := start + 1; current < BatchSize; current++ {
		increaseDigit(&digits)
		baseSum, sum := calcBothSum(&digits)
		d1, d2 := calc2Digits(baseSum, sum)
		local += d1 + d2
	}

	// now calculate by 10
	for current := BatchSize; current <= end; current += BatchSize {
		increaseDigit(&digits)
		baseSum0, sum0 := calcBothSum(&digits)
		for x := 0; x < BatchSize && current+x <= end; x++ {
			baseSum, sum := baseSum0+x, sum0+2*x
			d1, d2 := calc2Digits(baseSum, sum)
			local += d1 + d2
			if x < BatchSize-1 && current+x < end {
				increaseDigit(&digits)
			}
		}
	}
	return local
}

// LegacySoA uses a struct-of-arrays memory layout instead of array-of-structs.
// func LegacySoA(start, end int) int {
// 	var digits [CPF_DIGITS]int
// 	var local int

// 	for k := start; k <= end; k += BatchSize {
// 		var batchCpf [CPF_DIGITS][BatchSize]int
// 		for x := 0; x < BatchSize; x++ {
// 			for i := 8; i >= 0; i-- {
// 				digits[i]++
// 				if digits[i] < 10 {
// 					break
// 				}
// 				digits[i] = 0
// 			}
// 			for i := 0; i < CPF_DIGITS; i++ {
// 				batchCpf[i][x] = digits[i]
// 			}
// 		}

// 		baseSum := [BatchSize]int{}
// 		sum := [BatchSize]int{}

// 		for i := 0; i < CPF_DIGITS; i++ {
// 			for x := 0; x < BatchSize; x++ {
// 				digit := batchCpf[i][x]
// 				sum[x] += digit * Weights[i]
// 				baseSum[x] += digit
// 			}
// 		}

// 		for x := 0; x < BatchSize; x++ {
// 			rem := sum[x] % 11
// 			digit1 := remTable[rem]

// 			sum[x] += baseSum[x] + (digit1 << 1)

// 			rem = sum[x] % 11
// 			digit2 := remTable[rem]

// 			local += digit1 + digit2
// 		}
// 	}

// 	return local
// }

func GenCpfs(start int, end int) int {
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
	}
	return local
}

func old() {
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
		start := (n + 1) * amount
		end := min(start+amount, cpfsTotal)

		wait.Add(1)
		go func(start int, end int) {
			defer wait.Done()
			GenCpfs(start, end)
		}(start, end)
	}

	wait.Wait()
	end := time.Since(start)
	fmt.Printf("Calculated: %d in %v\n", cpfsTotal, end)
}
