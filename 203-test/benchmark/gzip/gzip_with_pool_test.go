package gzip

import (
	"fmt"
	"testing"
)

func BenchmarkGzipWithPool(b *testing.B) {
	gzip := NewGzipWithPool()

	fmt.Println("BenchmarkGzipWithPool is running...")

	// Önceki testlerden kalan dataları temizler
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		err := gzip.Zip("Ders kodu 203, Benchmark testi 1 -- 08.10.2023")

		if err != nil {
			b.Fatal(err)
		}
	}
}
