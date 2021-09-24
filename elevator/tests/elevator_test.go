package tests

import (
	"bufio"
	"elevator/elevator"
	"os"
	"testing"
)

type testData struct {
	innerQueue []*elevator.Call
	outerQueue []*elevator.Call
	position int
}

var tests = []testData{
	{
		[]*elevator.Call{{7, 1}, {3,0}, {2,1}},
		[]*elevator.Call{{6,1}, {5,0}},
		3,
	},
	{
		[]*elevator.Call{{7, 1}, {3,0}, {2,1}},
		[]*elevator.Call{{10,1}, {2,0}, {4,1}, {5,0}},
		2,
	},
	{
		[]*elevator.Call{{3, 1}, {5,1}, {4, 0}, {1,0}, {8, 1}},
		[]*elevator.Call{{9,1}, {6,0}, {3,1}, {5,0}},
		6,
	},
}

func TestElevator(t *testing.T) {
	for i, data := range tests {
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
		el.InnerQueue = data.innerQueue
		el.OuterQueue = data.outerQueue
		el.Move(&elevator.Call{
			Floor:   1,
			GoingUp: 0,
		}, 0)

		if el.Position != data.position {
			t.Error(
				"For ", i,
				"expected", data.position,
				"got",el.Position,
			)
		}
	}
}
