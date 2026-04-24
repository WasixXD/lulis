package main

import "testing"

func BenchmarkAritmeticLoop(b *testing.B) {
	var n [9]int = [9]int{0, 0, 0, 0, 0, 0, 0, 0, 1}
	for b.Loop() {
		sum := 0
		baseSum := 0
		for i := 0; i < 9; i++ {
			digit := n[i]
			sum += digit * (10 - i)
			baseSum += digit
		}
	}
}

func BenchmarkDoubleVarLoop(b *testing.B) {
	var n [9]int = [9]int{0, 0, 0, 0, 0, 0, 0, 0, 1}
	for b.Loop() {
		sum := 0
		baseSum := 0
		for i, j := 0, 10; i < 9; i, j = i+1, j-1 {
			digit := n[i]
			sum += digit * j
			baseSum += digit
		}
	}
}

func BenchmarkNoBoundCheck(b *testing.B) {

	var n [9]int = [9]int{0, 0, 0, 0, 0, 0, 0, 0, 1}

	for b.Loop() {
		sum := 0
		baseSum := 0
		for i, digit := range &n {
			sum += digit * (10 - i)
			baseSum += digit
		}
	}

}
