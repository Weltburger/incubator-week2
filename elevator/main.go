package main

import (
	"bufio"
	"elevator/elevator"
	"os"
	"sync"
)

func main() {
	var once sync.Once
	el := &elevator.Elevator{
		Scanner:      bufio.NewScanner(os.Stdin),
		InnerQueue:   make([]*elevator.Call, 0),
		OuterQueue:   make([]*elevator.Call, 0),
		Position:     1,
		IsMoving:     false,
		Stat:         true,
		Ch:           make(chan string, 10),
		Inner:        make(chan string, 10),
	}

	once.Do(el.Launch)

	for el.Scanner.Scan() {
		if el.Stat {
			el.Ch <- el.Scanner.Text()
		} else {
			el.Inner <-el.Scanner.Text()
		}

	}
}
