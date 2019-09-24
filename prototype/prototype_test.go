package prototype

import (
	"reflect"
	"testing"
)

func TestClone(t *testing.T) {
	e2 := element{name: "2"}
	e1 := element{
		name:  "1",
		child: []node{e2},
	}
	e3 := e1.Clone()

	if !reflect.DeepEqual(e1, e3) {
		t.Error("Expected two nodes to be same")
	}
}

func TestName(t *testing.T) {
	e2 := element{name: "2"}
	e1 := element{
		name:  "1",
		child: []node{e2},
	}
	e3 := e1.Clone()

	if !(e1.Name() == e3.Name()) {
		t.Error("Expected two nodes to be same")
	}
}
