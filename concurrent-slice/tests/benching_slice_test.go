package tests

import (
	c "concurent-slice/slice"
	"sync"
	"testing"
)

//go test -bench=BenchmarkInsertIntMap -benchmem -benchtime=100000x
func BenchmarkAppendConcurrentSlice(b *testing.B) {
	slice := new(c.ConcurrentSlice)
	wg := new(sync.WaitGroup)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			slice.Add(i)
		}(i)
	}
	wg.Wait()
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
		wg := new(sync.WaitGroup)
		for i := 0; i < pair.iter; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				slice.Add(i)
			}(i)
		}
		wg.Wait()
		if slice.Len() != pair.size {
			t.Error(
				"For", pair.iter,
				"expected", pair.size,
				"got", slice.Len(),
			)
		}
	}
}
