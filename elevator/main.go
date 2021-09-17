package main

import (
	"bufio"
	"elevator/elevator"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	el := &elevator.Elevator{
		Scanner:      bufio.NewScanner(os.Stdin),
		MemoryUp:     make(map[int]bool, 0),
		MemoryDown:   make(map[int]bool, 0),
		InnerQueue:   make([]*elevator.Call, 0),
		OuterQueue:   make([]*elevator.Call, 0),
		Position:     1,
		IsMovingUp:   false,
		IsMovingDown: false,
		Stat:         true,
		Ch:           make(chan string, 10),
		Inner:        make(chan string, 10),
	}

	go el.Launch()

	for el.Scanner.Scan() {
		if el.Stat {
			el.Ch <- el.Scanner.Text()
		} else {
			el.Inner <-el.Scanner.Text()
		}

	}
}
