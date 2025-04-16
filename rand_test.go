package rand

import "testing"

func Test_GenLetter(t *testing.T) {
	const n = 10000
	for i := 0; i < n; i++ {
		b := make([]byte, 32)
		GenLetter(b)
	}
}

func BenchmarkGenLetter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b1 := make([]byte, 64)
		GenLetter(b1)
	}
}

func BenchmarkGenLower(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b1 := make([]byte, 64)
		GenLower(b1)
	}
}

func BenchmarkGenUpper(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b1 := make([]byte, 64)
		GenUpper(b1)
	}
}

func BenchmarkGenNum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b1 := make([]byte, 64)
		GenNum(b1)
	}
}

func BenchmarkGenAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b1 := make([]byte, 64)
		GenAll(b1)
	}
}
