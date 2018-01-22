package main

import "testing"

func TestSum(t *testing.T) {
	got := sum(1, 2)
	if got != 3 {
		t.Error("wrong sum")
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum(1, 2)
	}
}
