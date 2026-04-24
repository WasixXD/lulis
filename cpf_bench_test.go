package main

import "testing"

var pow10bench = pow10table

func BenchmarkWeightedSum(b *testing.B) {
	n := 123456789
	for b.Loop() {
		value := 0
		sum := 0
		for i := 0; i < 9; i++ {
			num := (n / pow10bench[8-i]) % 10
			value += (10 - i) * num
			sum += num
		}
	}
}

func BenchmarkDigitCalc(b *testing.B) {
	n := uint64(123456789)

	for b.Loop() {
		x := n

		sum := uint64(0)
		value := uint64(0)

		for i := 0; i < 9; i++ {
			num := mod10(x)
			x = div10(x)

			value += uint64(i+2) * num
			sum += num
		}
	}
}
