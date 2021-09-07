package slice

import (
	"reflect"
	"sync"
)

type ConcurrentSlice struct {
	mux   sync.Mutex
	slice []interface{}
}

type ConcurrentSliceItem struct {
	Index int
	Value interface{}
}

func (cs *ConcurrentSlice) mutexLocked() bool {
	state := reflect.ValueOf(&cs.mux).Elem().FieldByName("state")
	return state.Int() == 1
}

func (cs *ConcurrentSlice) Add(val interface{}) {
	cs.mux.Lock()
	defer cs.mux.Unlock()
	cs.slice = append(cs.slice, val)
}

func (cs *ConcurrentSlice) Get(i int) interface{} {
	cs.mux.Lock()
	defer cs.mux.Unlock()
	return cs.slice[i]
}

func (cs *ConcurrentSlice) Iter() <-chan ConcurrentSliceItem {
	c := make(chan ConcurrentSliceItem)

	f := func() {
		cs.mux.Lock()
		defer cs.mux.Unlock()
		for index, value := range cs.slice {
			c <- ConcurrentSliceItem{index, value}
		}
		close(c)
	}
	go f()

	return c
}

func (cs *ConcurrentSlice) Len() int {
	return len(cs.slice)
}

func (cs *ConcurrentSlice) Cap() int {
	return cap(cs.slice)
}
