package main

import (
	"elevator/elevator"
)

func main() {
	el := elevator.Construct()
	el.Launch()

	for el.Scanner.Scan() {
		if el.Stat {
			el.Ch <- el.Scanner.Text()
		} else {
			el.Inner <-el.Scanner.Text()
		}

	}
}
