package composite

import (
	"testing"
)

func TestBicycleShouldBeReady(t *testing.T) {
	b := newBicycle()

	t1 := newTyre(31)
	b.addPart(t1)

	t2 := newTyre(34)
	b.addPart(t2)

	c := newChain()
	b.addPart(c)

	ready := b.isReady()
	if !ready {
		t.Errorf("expected cycle to be ready, but found not ready")
	}
}

func TestBicycleShouldNotBeReady(t *testing.T) {
	b := newBicycle()

	t1 := newTyre(3)
	b.addPart(t1)

	t2 := newTyre(34)
	b.addPart(t2)

	c := newChain()
	b.addPart(c)

	ready := b.isReady()
	if ready {
		t.Errorf("expected cycle to be not ready, but found ready")
	}
}
