package elevator

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Elevator struct {
	sync.Mutex
	Scanner      *bufio.Scanner
	InnerQueue   []*Call
	OuterQueue   []*Call
	Position     int
	IsMoving     bool
	Stat         bool
	Ch           chan string
	Inner        chan string
}

type Call struct {
	Floor   int
	GoingUp int
}

func (el *Elevator) clearing(state int) bool {
	queue1, queue2 := false, false
	el.Mutex.Lock()
	el.InnerQueue, queue1 = findAndDeleteIndexes(el.InnerQueue, el.Position, state)
	el.OuterQueue, queue2 = findAndDeleteIndexes(el.OuterQueue, el.Position, state)
	el.Mutex.Unlock()

	return queue1 || queue2
}

func (el *Elevator) openCloseMessage() {
	fmt.Println("Door is opening. Current Floor: ", el.Position)
	time.Sleep(time.Second * 2)
	fmt.Println("Door is closing")
	time.Sleep(time.Second * 2)
}

func (el *Elevator) Move(call *Call) {
	var way int
	if el.Position > call.Floor {
		way = 0
	} else {
		way = 1
	}

	el.IsMoving = true
	f := call.Floor
	for el.Position != f {
		if el.clearing(way) {
			el.openCloseMessage()
			el.Stat = false
			newInnerCall := el.innerCall()
			if newInnerCall != nil {
				fmt.Println(newInnerCall)
				el.InnerQueue = append(el.InnerQueue, newInnerCall)
			}
		}
		fmt.Println("Moving. Current Floor: ", el.Position)
		time.Sleep(time.Second * 2)
		if way == 1 {
			el.Position++
		} else {
			el.Position--
		}
	}

	el.clearing(way)
	el.openCloseMessage()

	el.IsMoving = false
	el.Stat = false
	newInnerCall := el.innerCall()
	if newInnerCall != nil {
		fmt.Println(newInnerCall)
		el.InnerQueue = append(el.InnerQueue, newInnerCall)
	}
	if newCall := el.isCalled(); newCall != nil {
		el.Move(newCall)
		return
	}
	fmt.Println("Done. Current Floor: ", el.Position)
}

func (el *Elevator) isCalled() *Call {
	if len(el.InnerQueue) != 0 {
		call := el.InnerQueue[0]
		el.InnerQueue = el.InnerQueue[1:]
		return call
	}
	if len(el.OuterQueue) != 0 {
		call := el.OuterQueue[0]
		el.OuterQueue = el.OuterQueue[1:]
		return call
	}
	return nil
}

func findAndDeleteIndexes(arr []*Call, n, state int) ([]*Call, bool) {
	res := make([]int, 0)
	isNEmpty := false
	dec := 0
	for i, v := range arr {
		if v.Floor == n {
			if v.GoingUp == state {
				res = append(res, i-dec)
				dec++
				isNEmpty = true
			}
		}
	}

	if len(res) != 0 {
		for _, v := range res {
			if v == len(arr)-1 {
				arr = arr[:v]
			} else {
				arr = append(arr[:v], arr[v+1:]...)
			}
		}
	}
	return arr, isNEmpty
}

func (el *Elevator) innerCall() *Call {
	fmt.Println("Choose the Floor:")
	ticker := time.NewTicker(time.Second * 5)
	c := new(Call)
	for !el.Stat {
		select {
		case str := <-el.Inner:
			i, err := strconv.Atoi(str)
			if err != nil {
				fmt.Println(err)
				continue
			}
			if i == el.Position {
				fmt.Println("This is your current Floor!")
				c = nil
				el.Stat = true
				break
			} else {
				c.Floor = i
				el.Stat = true
				break
			}
		case t := <-ticker.C:
			fmt.Println("Time is up", t)
			c = nil
			el.Stat = true
			break
		}
	}

	if c != nil {
		return c
	}
	return nil
}

func (el *Elevator) Launch() {
	go func() {
		for {
			select {
			case floor := <-el.Ch:
				arr := strings.Split(floor, " ")
				if len(arr) > 1 {
					f, err := strconv.Atoi(arr[0])
					if err != nil {
						fmt.Println(err)
						continue
					}
					b, err := strconv.Atoi(arr[1])
					if err != nil {
						fmt.Println(err)
						continue
					}
					c := &Call{
						Floor:   f,
						GoingUp: b,
					}
					if el.IsMoving {
						el.Mutex.Lock()
						el.OuterQueue = append(el.OuterQueue, c)
						el.Mutex.Unlock()
					} else {
						go el.Move(c)
					}
				}
			}
		}
	}()
}
