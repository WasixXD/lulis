package main

import (
	"testing"
)

var sink int
var a = [4][9]int{
	{3, 1, 5, 9, 9, 0, 7, 1, 2},
	{4, 6, 5, 8, 1, 9, 6, 8, 2},
	{5, 0, 7, 0, 3, 0, 7, 0, 2},
	{7, 9, 3, 1, 9, 0, 4, 2, 2},
}
var (
	s0, s1, s2, s3 int
)

func Benchmark_DotNaive(b *testing.B) {

	for b.Loop() {
		sink = DotNaive(a, weights)
	}
}

func Benchmark_DotUnroll(b *testing.B) {

	for b.Loop() {
		sink = DotUnroll(a, weights)
	}
}

func Benchmark_DotBCE(b *testing.B) {

	for b.Loop() {
		sink = DotBCE(a, weights)
	}
}
func Benchmark_DotFullUnroll(b *testing.B) {

	for b.Loop() {
		sink = DotFullUnroll(a, weights)
	}
}

func Benchmark_Dot4(b *testing.B) {
	for b.Loop() {
		s0, s1, s2, s3 = Dot4(a, weights)
	}
}
