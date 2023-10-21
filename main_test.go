package main

import (
	"testing"
)

// goos: darwin
// goarch: amd64
// pkg: goles_merge
// cpu: Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz
// BenchmarkMergeOne-8            6         171473951 ns/op
// PASS
// ok      goles_merge     16.611s

func BenchmarkMergeOne(b *testing.B) {
	a1 := getALongSortedArray(10000000)
	b1 := getALongSortedArray(10000000)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		mergeOne(a1, b1)
	}
}

// goos: darwin
// goarch: amd64
// pkg: goles_merge
// cpu: Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz
// BenchmarkMergeTwo-8            7         161867971 ns/op
// PASS
// ok      goles_merge     12.451s

func BenchmarkMergeTwo(b *testing.B) {
	a1 := getALongSortedArray(10000000)
	b1 := getALongSortedArray(10000000)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		mergeTwo(a1, b1)
	}
}

func TestGetALongSortedArray(t *testing.T) {
	a := getALongSortedArray(10)
	greeting := len(a)
	if greeting != 10 {
		t.Errorf("Length of results array not equals to %d, current length is %d", 10, greeting)
	}
}

func TestMergeTwoFindLess(t *testing.T) {
	var la []int = []int{100, 200, 300}
	var lk int = 1
	var ar []int = []int{100, 125, 150, 175, 200, 225, 250, 275, 300}
	var s int = 0

	greeting := mergeTwoFindLess(la, lk, ar, &s)

	if !greeting {
		t.Errorf("Result is false!")
	}

	if s != 5 {
		t.Errorf("Result is incorrect: %d, need %d", s, 5)
	}
}
