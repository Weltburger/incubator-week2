package slice

import (
	"reflect"
	"sync"
)

type ConcurrentSlice struct {
	mux sync.Mutex
	slice []interface{}
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

func (cs *ConcurrentSlice) Len() int {
	return len(cs.slice)
}

func (cs *ConcurrentSlice) Cap() int {
	return cap(cs.slice)
}
