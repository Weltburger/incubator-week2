package main

import (
	c "concurent-slice/slice"
	"fmt"
	"sync"
	"time"
)

func main() {
	cs := new(c.ConcurrentSlice)
	wg := new(sync.WaitGroup)

	start := time.Now()

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			cs.Add(i)
		}(i)
	}
	wg.Wait()

	fmt.Println(time.Since(start))
	//fmt.Println(cs.slice)
	//fmt.Println(len(cs.slice))
}
