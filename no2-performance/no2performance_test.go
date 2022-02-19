package no2performance_test

import (
	codeA "citcall-go-test/no2-performance/code-a"
	codeB "citcall-go-test/no2-performance/code-b"
	"testing"
)

func TestCodeA(t *testing.T) {
	codeA.Test()
}

func TestCodeB(t *testing.T) {
	codeB.Test()
}

func BenchmarkCodeA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		codeA.Test()
	}
}

func BenchmarkCodeB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		codeB.Test()
	}
}
