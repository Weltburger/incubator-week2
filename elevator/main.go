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

type Elevator struct {
	sync.Mutex
	memory       map[int]bool
	innerQueue   []*Call
	outerQueue   []*Call
	position     int
	isMovingUp   bool
	isMovingDown bool
	ch           chan string
}

type Call struct {
	floor int
	up    int
}

func (el *Elevator) Move(call *Call) {
	if el.position == call.floor {
		fmt.Println("This is your current position. Floor: ", el.position)
		el.memory[el.position] = false
	} else if call.up == 0 {
		el.Down(call.floor)
	} else {
		el.Up(call.floor)
	}
}

func (el *Elevator) Up(floor int) {
	el.isMovingDown = false
	el.isMovingUp = true
	f := floor
	for el.position != f {
		if newFloor := el.isCalledUp(el.position); newFloor != 0 && newFloor > f {
			el.memory[f] = true
			f = newFloor
		}
		fmt.Println("Moving up. Current floor: ", el.position)
		time.Sleep(time.Second)
		el.position++
	}
	el.memory[el.position] = false
	fmt.Println("Door is opening. Current floor: ", el.position)
	time.Sleep(time.Second)
	fmt.Println("Door is closing.")
	if newFloor := el.isCalled(); newFloor != 0 {
		el.Move(newFloor)
		return
	}
	el.isMoving = false
	fmt.Println("Done. Current floor: ", el.position)
}

func (el *Elevator) Down(floor int) {

}

func (el *Elevator) isCalledUp(floor int) {
	if el.memory[floor] {

	}
}

func main() {

	el := &Elevator{
		memory:       make(map[int]bool, 0),
		position:     1,
		innerQueue:   make([]*Call, 0),
		outerQueue:   make([]*Call, 0),
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
