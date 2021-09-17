package main

import (
	"cuncurrent-slice-2/slice"
	"fmt"
)

func main() {
	cs := &slice.ConcurrentSlice{}
	cs.Init()

	go cs.Listen()

	for i := 0; i < 100000; i++ {
		cs.WaitGroup.Add(1)
		go func(i int) {
			defer cs.WaitGroup.Done()
			cs.Add(i)
		}(i)
	}
	cs.WaitGroup.Wait()

	fmt.Println(cs.Len())
	fmt.Println(cs.Cap())
}
