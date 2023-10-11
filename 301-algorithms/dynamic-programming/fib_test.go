package main

import "testing"

// Benchmark -> BenchmarkFib1-12          447574              2617 ns/op               0 B/op
func BenchmarkFib1(b *testing.B) {

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = Fib1(15)
	}
}

// Benchmark -> BenchmarkFib2-12        825429842                1.397 ns/op           0 B/op
func BenchmarkFib2(b *testing.B) {
	n := 15
	memo := make([]int, n)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = Fib2(n, memo)
	}
}

// Benchmark -> BenchmarkFib3-12        155726814                7.593 ns/op           0 B/op
func BenchmarkFib3(b *testing.B) {
	n := 15
	memo := make([]int, n+1)

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = Fib3(n, memo)
	}
}
