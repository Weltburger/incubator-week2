package tests

import (
	c "cuncurrent-slice-2/slice"
	"testing"
	"time"
)

//go test -bench=BenchmarkInsertIntMap -benchmem -benchtime=100000x
func BenchmarkAppendConcurrentSlice(b *testing.B) {
	slice := new(c.ConcurrentSlice)
	slice.Init()
	go slice.Listen()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice.WaitGroup.Add(1)
		go func(i int) {
			defer slice.WaitGroup.Done()
			slice.Add(i)
		}(i)
	}
	slice.WaitGroup.Wait()
}


type testpair struct {
	iter int
	size int
}

var tests = []testpair{
	{ 700, 700 },
	{ 1500, 1500 },
	{ 10000, 10000 },
}



func TestConcurrentSlice(t *testing.T) {
	for _, pair := range tests {
		slice := new(c.ConcurrentSlice)
		slice.Init()
		go slice.Listen()
		for i := 0; i < pair.iter; i++ {
			slice.WaitGroup.Add(1)
			go func(i int) {
				defer slice.WaitGroup.Done()
				slice.Add(i)
			}(i)
		}
		slice.WaitGroup.Wait()
		if slice.Len() != pair.size {
			t.Error(
				"For", pair.iter,
				"expected", pair.size,
				"got", slice.Len(),
			)
		}
	}
}
