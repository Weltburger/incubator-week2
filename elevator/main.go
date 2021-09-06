package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

type Elevator struct {
	sync.Mutex
	sync.WaitGroup
	memory map[int]bool
	position int
	isMoving bool
	ch chan string
}

func (el *Elevator) Move(floor int) {
	if el.position == floor {
		fmt.Println("This is your current floor")
		el.memory[el.position] = false
		if newFloor := el.isCalled(); newFloor != 0 {
			el.Move(newFloor)
		}
		return
	}
	if el.position > floor {
		el.Down(floor)
		return
	} else {
		el.Up(floor)
		return
	}
}

func (el *Elevator) Up(floor int) {
	el.isMoving = true
	f := floor
	for el.position != f {
		if newFloor := el.isCalled(); newFloor != 0 && newFloor > f {
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
	el.isMoving = true
	f := floor
	for el.position != f {
		fmt.Println("Moving down. Current floor: ", el.position)
		time.Sleep(time.Second)
		el.memory[el.position] = false
		el.position--
		if el.memory[el.position] {
			fmt.Println("Door is opening. Current floor: ", el.position)
			time.Sleep(time.Second)
			fmt.Println("Door is closing.")
		}
	}
	el.memory[el.position] = false
	if newFloor := el.isCalled(); newFloor != 0 {
		el.Move(newFloor)
		return
	}
	el.isMoving = false
	fmt.Println("Done. Current floor: ", el.position)
}

func (el *Elevator) isCalled() int {
	maxFloor := 0
	for i := range el.memory {
		if el.memory[i] && i > maxFloor {
			maxFloor = i
		}
	}
	return maxFloor
}

func (el *Elevator) Set(floor int) {
	el.position = floor
}

func main() {

	el := &Elevator{
		memory:    make(map[int]bool, 0),
		position:  1,
		isMoving:  false,
		ch : make(chan string, 10),
	}

	go func() {
		for {
			select {
			case floor := <-el.ch:
				i, err := strconv.Atoi(floor)
				if err != nil {
					log.Fatal(err)
				}
				if el.isMoving {
					el.Mutex.Lock()
					el.memory[i] = true
					el.Mutex.Unlock()
				} else {
					go el.Move(i)
				}
			default:
				time.Sleep(time.Millisecond * 500)
			}
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		el.ch <- scanner.Text()
	}
}
