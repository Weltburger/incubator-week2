package slice

import (
	"fmt"
	"sync"
)

type ConcurrentSlice struct {
	ch    chan interface{}
	slice []interface{}
	*sync.WaitGroup
}

func (cs *ConcurrentSlice) Init() {
	cs.ch = make(chan interface{}, 10)
	cs.slice = make([]interface{}, 0)
	cs.WaitGroup = new(sync.WaitGroup)
}

func (cs *ConcurrentSlice) Add(i interface{}) {
	cs.ch <- i
}

func (cs *ConcurrentSlice) Listen() {
	for {
		select {
		case x := <-cs.ch:
			cs.slice = append(cs.slice, x)
		}
	}
}

func (cs *ConcurrentSlice) Get(i int) interface{} {
	return cs.slice[i]
}

func (cs *ConcurrentSlice) Len() int {
	return len(cs.slice)
}

func (cs *ConcurrentSlice) Cap() int {
	return cap(cs.slice)
}

func (cs *ConcurrentSlice) Print() {
	fmt.Println(cs.slice)
}
