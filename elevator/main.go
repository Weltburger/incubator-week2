package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Elevator struct {
	sync.Mutex
	scanner      *bufio.Scanner
	memoryUp     map[int]bool
	memoryDown   map[int]bool
	innerQueue   []*Call
	outerQueue   []*Call
	position     int
	isMovingUp   bool
	isMovingDown bool
	stat         bool
	ch           chan string
	inner        chan string
}

type Call struct {
	floor   int
	goingUp int
}

func (el *Elevator) Move(call *Call) {
	if el.position == call.floor {
		fmt.Println("This is your current position. Floor: ", el.position)
		el.memoryUp[el.position] = false
		el.memoryDown[el.position] = false
		if newCall := el.isCalled(); newCall != nil {
			el.Move(newCall)
			return
		}
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
				time.Sleep(time.Second * 2)
				fmt.Println("Door is closing")
				time.Sleep(time.Second * 2)
				if len(queue1) != 0 {
					el.innerQueue = removeItems(el.innerQueue, queue1)
				}
				if len(queue2) != 0 {
					el.outerQueue = removeItems(el.outerQueue, queue2)
				}
			}
			el.memoryUp[el.position] = false

			el.stat = false
			newInnerCall := el.innerCall()
			if newInnerCall != nil {
				fmt.Println(newInnerCall)
				el.innerQueue = append(el.innerQueue, newInnerCall)
			}
		}
		fmt.Println("Moving up. Current floor: ", el.position)
		time.Sleep(time.Second * 2)
		el.position++
	}
	el.memoryUp[el.position] = false
	el.memoryDown[el.position] = false
	fmt.Println("Door is opening. Current floor: ", el.position)
	time.Sleep(time.Second * 2)
	fmt.Println("Door is closing.")
	el.isMovingUp = false
	el.stat = false
	newInnerCall := el.innerCall()
	if newInnerCall != nil {
		fmt.Println(newInnerCall)
		el.innerQueue = append(el.innerQueue, newInnerCall)
	}
	if newCall := el.isCalled(); newCall != nil {
		el.Move(newCall)
		return
	}
	fmt.Println("Done. Current floor: ", el.position)
}

func (el *Elevator) Down(call *Call) {
	el.isMovingUp = false
	el.isMovingDown = true
	f := call.floor
	for el.position != f {
		if el.memoryDown[el.position] {
			queue1 := findIndexesDown(el.innerQueue, el.position)
			queue2 := findIndexesDown(el.outerQueue, el.position)
			if len(queue1) != 0 || len(queue2) != 0 {
				fmt.Println("Door is opening. Current floor: ", el.position)
				time.Sleep(time.Second * 2)
				fmt.Println("Door is closing")
				time.Sleep(time.Second * 2)
				if len(queue1) != 0 {
					el.innerQueue = removeItems(el.innerQueue, queue1)
				}
				if len(queue2) != 0 {
					el.outerQueue = removeItems(el.outerQueue, queue2)
				}
			}
			el.memoryDown[el.position] = false

			el.stat = false
			newInnerCall := el.innerCall()
			if newInnerCall != nil {
				fmt.Println(newInnerCall)
				el.innerQueue = append(el.innerQueue, newInnerCall)
			}
		}
		fmt.Println("Moving down. Current floor: ", el.position)
		time.Sleep(time.Second * 2)
		el.position--
	}
	el.memoryUp[el.position] = false
	el.memoryDown[el.position] = false
	fmt.Println("Door is opening. Current floor: ", el.position)
	time.Sleep(time.Second * 2)
	fmt.Println("Door is closing.")
	el.isMovingDown = false
	el.stat = false
	newInnerCall := el.innerCall()
	if newInnerCall != nil {
		fmt.Println(newInnerCall)
		el.innerQueue = append(el.innerQueue, newInnerCall)
	}
	if newCall := el.isCalled(); newCall != nil {
		el.Move(newCall)
		return
	}
	fmt.Println("Done. Current floor: ", el.position)
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
			if v.goingUp == 1 {
				res = append(res, i-dec)
				dec++
			}
		}
	}
	return res
}

func findIndexesDown(arr []*Call, n int) []int {
	res := make([]int, 0)
	dec := 0
	for i, v := range arr {
		if v.floor == n {
			if v.goingUp == 0 {
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

func (el *Elevator) innerCall() *Call {
	fmt.Println("Choose the floor:")
	ticker := time.NewTicker(time.Second * 5)
	c := new(Call)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func(c *Call) {
		defer wg.Done()
		for !el.stat {
			select {
			case str := <- el.inner:
				i, err := strconv.Atoi(str)
				if err != nil {
					fmt.Println(err)
					continue
				}
				el.memoryUp[i] = true
				if i > el.position {
					c.floor = i
					c.goingUp = 1
					el.stat = true
					break
				} else if i < el.position {
					c.floor = i
					c.goingUp = 0
					el.stat = true
					break
				} else {
					fmt.Println("This is your current floor!")
					c = nil
					el.stat = true
					break
				}
			case t := <-ticker.C:
				fmt.Println("Time is up", t)
				c = nil
				el.stat = true
				break
			}
		}
	}(c)
	wg.Wait()
	if c.floor > 0 {
		return c
	}
	return nil
}

func (el *Elevator) launch() {
	go func() {
		for {
			select {
			case floor := <-el.ch:
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
						floor:   f,
						goingUp: b,
					}
					if el.isMovingUp || el.isMovingDown {
						if b == 0 {
							el.memoryDown[f] = true
						} else {
							el.memoryUp[f] = true
						}
						el.outerQueue = append(el.outerQueue, c)
					} else {
						el.Mutex.Lock()
						if b == 0 {
							el.memoryDown[f] = true
						} else {
							el.memoryUp[f] = true
						}
						el.Mutex.Unlock()
						go el.Move(c)
					}
				}
			}
		}
	}()
}

func main() {
	rand.Seed(time.Now().UnixNano())

	el := &Elevator{
		scanner: bufio.NewScanner(os.Stdin),
		memoryUp:     make(map[int]bool, 0),
		memoryDown:   make(map[int]bool, 0),
		innerQueue:   make([]*Call, 0),
		outerQueue:   make([]*Call, 0),
		position:     1,
		isMovingUp:   false,
		isMovingDown: false,
		stat:         true,
		ch:           make(chan string, 10),
		inner:        make(chan string, 10),
	}

	go el.launch()

	for el.scanner.Scan() {
		if el.stat {
			el.ch <- el.scanner.Text()
		} else {
			el.inner <-el.scanner.Text()
		}

	}
}
