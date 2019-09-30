package composite

import (
	"fmt"
	"strconv"
)

type part interface {
	working() bool
}

type bicycle struct {
	parts []part
}

func newBicycle() bicycle {
	b := bicycle{
		parts: make([]part, 0),
	}
	return b
}

func (b *bicycle) addPart(p part) {
	b.parts = append(b.parts, p)
}

func (b bicycle) isReady() bool {
	ready := true
	for _, individualPart := range b.parts {

		fmt.Println(individualPart)
		individualPartWorking := individualPart.working()
		ready = ready && individualPartWorking
	}
	return ready
}

type tyre struct {
	airPressure int
}

func newTyre(ap int) tyre {
	return tyre{
		airPressure: ap,
	}
}

func (t tyre) working() bool {
	return t.airPressure > 30
}

func (t tyre) String() string {
	return "Part name: tyre, Air Pressure:" + strconv.Itoa(t.airPressure)
}

func newChain() chain {
	return chain{}
}

type chain struct{}

func (c chain) working() bool {
	return true
}

func (c chain) String() string {
	return "Part name: Chain"
}
