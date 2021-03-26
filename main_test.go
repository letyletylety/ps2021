package main

import "testing"

func BenchmarkForNormal(b *testing.B) {
	n := int(5e6)
	b.Logf("%d\n", n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = i
	}

	sum := 0
	for i := 0; i < n; i++ {
		sum += i
	}
	b.Logf("%d\n", sum)
}

func BenchmarkForRange(b *testing.B) {
	n := int(5e6)
	b.Logf("%d\n", n)

	a := make([]int, n)
	for i := range a {
		a[i] = i
	}

	sum := 0
	for _, v := range a {
		sum += v
	}
	b.Logf("%d\n", sum)
}
