package bridge

import "testing"

func TestBridgePattern(t *testing.T) {
	h := newAnimal(human, newWalk())
	err := h.move()
	if err != nil {
		t.Error("expected err to be nil but got", err)
	}

	f := newAnimal(fish, newSwim())
	err = f.move()
	if err != nil {
		t.Error("expected err to be nil but got", err)
	}
}
