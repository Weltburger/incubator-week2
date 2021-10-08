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
		[]*elevator.Call{{7, true}, {3,false}, {2,true}},
		[]*elevator.Call{{6,true}, {5,false}},
		3,
	},
	{
		[]*elevator.Call{{7, true}, {3,false}, {2,true}},
		[]*elevator.Call{{10,true}, {2,false}, {4,true}, {5,false}},
		2,
	},
	{
		[]*elevator.Call{{3, true}, {5,true}, {4, false}, {1,false}, {8, true}},
		[]*elevator.Call{{9,true}, {6,false}, {3,true}, {5,false}},
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
			GoingUp: false,
		})

		if el.Position != data.position {
			t.Error(
				"For ", i,
				"expected", data.position,
				"got",el.Position,
			)
		}
	}
}
