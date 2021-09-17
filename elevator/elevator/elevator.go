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
	MemoryUp     map[int]bool
	MemoryDown   map[int]bool
	InnerQueue   []*Call
	OuterQueue   []*Call
	Position     int
	IsMovingUp   bool
	IsMovingDown bool
	Stat         bool
	Ch           chan string
	Inner        chan string
}

type Call struct {
	Floor   int
	GoingUp int
}

func (el *Elevator) clearingUp() {
	queue1 := findIndexesUp(el.InnerQueue, el.Position)
	queue2 := findIndexesUp(el.OuterQueue, el.Position)
	if len(queue1) != 0 || len(queue2) != 0 {
		fmt.Println("Door is opening. Current Floor: ", el.Position)
		time.Sleep(time.Second * 2)
		fmt.Println("Door is closing")
		time.Sleep(time.Second * 2)
		if len(queue1) != 0 {
			el.InnerQueue = removeItems(el.InnerQueue, queue1)
		}
		if len(queue2) != 0 {
			el.OuterQueue = removeItems(el.OuterQueue, queue2)
		}
	}
}

func (el *Elevator) clearingDown() {
	queue1 := findIndexesDown(el.InnerQueue, el.Position)
	queue2 := findIndexesDown(el.OuterQueue, el.Position)
	if len(queue1) != 0 || len(queue2) != 0 {
		fmt.Println("Door is opening. Current Floor: ", el.Position)
		time.Sleep(time.Second * 2)
		fmt.Println("Door is closing")
		time.Sleep(time.Second * 2)
		if len(queue1) != 0 {
			el.InnerQueue = removeItems(el.InnerQueue, queue1)
		}
		if len(queue2) != 0 {
			el.OuterQueue = removeItems(el.OuterQueue, queue2)
		}
	}
}

func (el *Elevator) Move(call *Call) {
	if el.Position == call.Floor {
		fmt.Println("This is your current Position. Floor: ", el.Position)
		el.MemoryUp[el.Position] = false
		el.MemoryDown[el.Position] = false
		if newCall := el.isCalled(); newCall != nil {
			el.Move(newCall)
			return
		}
	} else if call.Floor < el.Position {
		el.Down(call)
	} else {
		el.Up(call)
	}
}

func (el *Elevator) Up(call *Call) {
	el.IsMovingDown = false
	el.IsMovingUp = true
	f := call.Floor
	for el.Position != f {
		if el.MemoryUp[el.Position] {
			el.clearingUp()

			el.MemoryUp[el.Position] = false
			el.Stat = false
			newInnerCall := el.innerCall()
			if newInnerCall != nil {
				fmt.Println(newInnerCall)
				el.InnerQueue = append(el.InnerQueue, newInnerCall)
			}
		}
		fmt.Println("Moving up. Current Floor: ", el.Position)
		time.Sleep(time.Second * 2)
		el.Position++
	}
	el.MemoryUp[el.Position] = false
	el.MemoryDown[el.Position] = false
	queue1 := findIndexesUp(el.InnerQueue, el.Position)
	queue2 := findIndexesUp(el.OuterQueue, el.Position)
	if len(queue1) != 0 || len(queue2) != 0 {
		if len(queue1) != 0 {
			el.InnerQueue = removeItems(el.InnerQueue, queue1)
		}
		if len(queue2) != 0 {
			el.OuterQueue = removeItems(el.OuterQueue, queue2)
		}
	}

	fmt.Println("Door is opening. Current Floor: ", el.Position)
	time.Sleep(time.Second * 2)
	fmt.Println("Door is closing.")
	el.IsMovingUp = false
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

func (el *Elevator) Down(call *Call) {
	el.IsMovingUp = false
	el.IsMovingDown = true
	f := call.Floor
	for el.Position != f {
		if el.MemoryDown[el.Position] {
			el.clearingDown()

			el.MemoryDown[el.Position] = false
			el.Stat = false
			newInnerCall := el.innerCall()
			if newInnerCall != nil {
				fmt.Println(newInnerCall)
				el.InnerQueue = append(el.InnerQueue, newInnerCall)
			}
		}
		fmt.Println("Moving down. Current Floor: ", el.Position)
		time.Sleep(time.Second * 2)
		el.Position--
	}
	el.MemoryUp[el.Position] = false
	el.MemoryDown[el.Position] = false
	queue1 := findIndexesDown(el.InnerQueue, el.Position)
	queue2 := findIndexesDown(el.OuterQueue, el.Position)
	if len(queue1) != 0 || len(queue2) != 0 {
		if len(queue1) != 0 {
			el.InnerQueue = removeItems(el.InnerQueue, queue1)
		}
		if len(queue2) != 0 {
			el.OuterQueue = removeItems(el.OuterQueue, queue2)
		}
	}

	fmt.Println("Door is opening. Current Floor: ", el.Position)
	time.Sleep(time.Second * 2)
	fmt.Println("Door is closing.")
	el.IsMovingDown = false
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

func findIndexesUp(arr []*Call, n int) []int {
	res := make([]int, 0)
	dec := 0
	for i, v := range arr {
		if v.Floor == n {
			if v.GoingUp == 1 {
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
		if v.Floor == n {
			if v.GoingUp == 0 {
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
	fmt.Println("Choose the Floor:")
	ticker := time.NewTicker(time.Second * 5)
	c := new(Call)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func(c *Call) {
		defer wg.Done()
		for !el.Stat {
			select {
			case str := <-el.Inner:
				i, err := strconv.Atoi(str)
				if err != nil {
					fmt.Println(err)
					continue
				}
				el.MemoryUp[i] = true
				if i > el.Position {
					c.Floor = i
					c.GoingUp = 1
					el.Stat = true
					break
				} else if i < el.Position {
					c.Floor = i
					c.GoingUp = 0
					el.Stat = true
					break
				} else {
					fmt.Println("This is your current Floor!")
					c = nil
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
	}(c)
	wg.Wait()
	if c.Floor > 0 {
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
					if el.IsMovingUp || el.IsMovingDown {
						if b == 0 {
							el.MemoryDown[f] = true
						} else {
							el.MemoryUp[f] = true
						}
						el.OuterQueue = append(el.OuterQueue, c)
					} else {
						el.Mutex.Lock()
						if b == 0 {
							el.MemoryDown[f] = true
						} else {
							el.MemoryUp[f] = true
						}
						el.Mutex.Unlock()
						go el.Move(c)
					}
				}
			}
		}
	}()
}
