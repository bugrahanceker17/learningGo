package gzip

import (
	"fmt"
	"testing"
)

func BenchmarkGzipWithoutPool(b *testing.B) {
	gzip := NewGzipWithoutPool()

	fmt.Println("BenchmarkGzipWithoutPool is running...")

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		err := gzip.Zip("Ders kodu 203, Benchmark testi 2 -- 08.10.2023")

		if err != nil {
			b.Fatal(err)
		}
	}
}
