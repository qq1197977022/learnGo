package fabonacci

import "testing"

func BenchmarkRecursion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Recursion(10)
	}
}

func BenchmarkNoRecusion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NoRecusion(10)
	}
}
