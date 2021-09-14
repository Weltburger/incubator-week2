package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Очередь вызовов (внутренняя и внешняя)
// Кнопка вызова вверх и вниз
// Память для очередей вверх и вниз? memory

type Elevator struct {
	sync.Mutex
	memoryUp     map[int]bool
	memoryDown   map[int]bool
	innerQueue   []*Call
	outerQueue   []*Call
	position     int
	isMovingUp   bool
	isMovingDown bool
	ch           chan string
}

type Call struct {
	floor   int
	goingUp bool
}

func (el *Elevator) Move(call *Call) {
	if el.position == call.floor {
		fmt.Println("This is your current position. Floor: ", el.position)
		el.memoryUp[el.position] = false
		el.memoryDown[el.position] = false
	} else if call.floor < el.position {
		el.Down(call)
	} else {
		el.Up(call)
	}
}

func (el *Elevator) Up(call *Call) {
	el.isMovingDown = false
	el.isMovingUp = true
	f := call.floor
	for el.position != f {
		if el.memoryUp[el.position] {
			queue1 := findIndexesUp(el.innerQueue, el.position)
			queue2 := findIndexesUp(el.outerQueue, el.position)
			if len(queue1) != 0 || len(queue2) != 0 {
				fmt.Println("Door is opening. Current floor: ", el.position)
				time.Sleep(time.Second)
				fmt.Println("Door is closing")
				time.Sleep(time.Second)
				if len(queue1) != 0 {
					removeItems(el.innerQueue, queue1)
				}
				if len(queue2) != 0 {
					removeItems(el.outerQueue, queue2)
				}
			}
			el.memoryUp[el.position] = false
		}
		fmt.Println("Moving up. Current floor: ", el.position)
		time.Sleep(time.Second)
		el.position++
	}
	el.memoryUp[el.position] = false
	fmt.Println("Door is opening. Current floor: ", el.position)
	time.Sleep(time.Second)
	fmt.Println("Door is closing.")
	el.isMovingUp = false
	if newCall := el.isCalled(); newCall != nil {
		el.Move(newCall)
		return
	}
	fmt.Println("Done. Current floor: ", el.position)
}

func (el *Elevator) Down(call *Call) {

}

func (el *Elevator) isCalled() *Call {
	if len(el.innerQueue) != 0 {
		call := el.innerQueue[0]
		el.innerQueue = el.innerQueue[1:]
		return call
	}
	if len(el.outerQueue) != 0 {
		call := el.outerQueue[0]
		el.outerQueue = el.outerQueue[1:]
		return call
	}
	return nil
}

func findIndexesUp(arr []*Call, n int) []int {
	res := make([]int, 0)
	dec := 0
	for i, v := range arr {
		if v.floor == n {
			if v.goingUp {
				res = append(res, i-dec)
				dec++
			}
		}
	}
	return res
}

func removeItems(arr []*Call, indexes []int) []*Call {
	if len(indexes) != 0 {
		for _, v := range indexes {
			if v == len(arr)-1 {
				arr = arr[:v]
			} else {
				arr = append(arr[:v], arr[v+1:]...)
			}
		}
		return arr
	}

	return arr
}

func main() {

	el := &Elevator{
		memory:       make(map[int]bool, 0),
		position:     1,
		innerQueue:   make([]int, 0),
		outerQueue:   make([]int, 0),
		isMovingUp:   false,
		isMovingDown: false,
		ch:           make(chan string, 10),
	}

	go func() {
		for {
			select {
			case floor := <-el.ch:
				arr := strings.Split(floor, " ")
				f, err := strconv.Atoi(arr[0])
				b, err := strconv.Atoi(arr[1])
				if err != nil {
					fmt.Println(err)
					continue
				}
				c := &Call{
					floor: f,
					up:    b,
				}
				if el.isMovingUp || el.isMovingDown {
					el.outerQueue = append(el.outerQueue, c)
				} else {
					el.Mutex.Lock()
					el.memory[f] = true
					if c.up == 0 {
						el.isMovingDown = true
						el.isMovingUp = false
					} else {
						el.isMovingUp = true
						el.isMovingDown = false
					}
					el.Mutex.Unlock()
					go el.Move(c)
				}
			}
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		el.ch <- scanner.Text()
	}
}
