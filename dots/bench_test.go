package dots

import "testing"

var weights = [9]int{10, 9, 8, 7, 6, 5, 4, 3, 2}

var sink int
var s0, s1, s2, s3 int

// a has BatchSize (10) rows; the original had 4 rows which didn't match function signatures.
var a = [BatchSize][9]int{
	{3, 1, 5, 9, 9, 0, 7, 1, 2},
	{4, 6, 5, 8, 1, 9, 6, 8, 2},
	{5, 0, 7, 0, 3, 0, 7, 0, 2},
	{7, 9, 3, 1, 9, 0, 4, 2, 2},
	{1, 2, 3, 4, 5, 6, 7, 8, 9},
	{9, 8, 7, 6, 5, 4, 3, 2, 1},
	{0, 1, 0, 1, 0, 1, 0, 1, 0},
	{2, 4, 6, 8, 0, 2, 4, 6, 8},
	{1, 1, 1, 1, 1, 1, 1, 1, 1},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
}

var a4 = [4][9]int{
	{3, 1, 5, 9, 9, 0, 7, 1, 2},
	{4, 6, 5, 8, 1, 9, 6, 8, 2},
	{5, 0, 7, 0, 3, 0, 7, 0, 2},
	{7, 9, 3, 1, 9, 0, 4, 2, 2},
}

func Benchmark_Naive(b *testing.B) {
	for b.Loop() {
		sink = Naive(a, weights)
	}
}

func Benchmark_Unroll(b *testing.B) {
	for b.Loop() {
		sink = Unroll(a, weights)
	}
}

func Benchmark_BCE(b *testing.B) {
	for b.Loop() {
		sink = BCE(a, weights)
	}
}

func Benchmark_FullUnroll(b *testing.B) {
	for b.Loop() {
		sink = FullUnroll(a4, weights)
	}
}

func Benchmark_Dot4(b *testing.B) {
	for b.Loop() {
		s0, s1, s2, s3 = Dot4(a, weights)
	}
}
