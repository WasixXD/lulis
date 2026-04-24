package main

import (
	"log"
	"os"
	"runtime/pprof"
	"strconv"
)

func toDigits(arr []string) []int {
	digits := make([]int, len(arr))
	for i := range arr {
		num, _ := strconv.Atoi(arr[i])
		digits[i] = num
	}
	return digits
}

func calc2Digits(digits *[9]int) []string {
	var local int

	sum := 0
	// holds A+B+C+D+E+F+G+H+I
	baseSum := 0
	for i, j := 0, 10; i < 9; i, j = i+1, j-1 {
		digit := digits[i]
		sum += digit * j
		baseSum += digit
	}

	// digit1 is (10A + 9B + 8C + 7D + 6E + 5F + 4G + 3H + 2I) % 11
	rem := sum % 11
	digit1 := 11 - rem
	if rem == 0 || rem == 1 {
		digit1 = 0
	}

	// digit2 is 11A + 10B + 9C + 8D + 7E + 6F + 5G + 4H + 3I + 2J
	// but can also be written as:
	// (10A + 9B + 8C + 7D + 6E + 5F + 4G + 3H + 2I) + (A+B+C+D+E+F+G+H+I) + 2 * digit1
	// we save one loop
	sum += baseSum + 2*digit1

	rem = sum % 11
	digit2 := 11 - rem
	if rem == 0 || rem == 1 {
		digit2 = 0
	}

	local += digit1 + digit2
	// fmt.Println(digits, digit1, digit2)

	return []string{}
}

func main() {
	f, _ := os.Create("cpu.prof")

	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal(err)
	}

	var digits [9]int
	for i := 1; i < 999_999_999; i++ {

		for i := 8; i >= 0; i-- {
			digits[i]++
			if digits[i] < 10 {
				break
			}
			digits[i] = 0
		}

		calc2Digits(&digits)
	}

	pprof.StopCPUProfile()
	f.Close()
}
