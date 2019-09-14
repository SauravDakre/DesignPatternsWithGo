package singleton

import (
	"testing"
)

func TestSingleton(t *testing.T) {
	c1 := New()
	if c1 == nil {
		t.Error("expecting pointer of singleton struct received ", c1)
	}

	c1.AddOne()
	newCount := c1.GetCount()
	if newCount != (1) {
		t.Error("expecting new count to be", (1), "but got", newCount)
	}

	c2 := New()
	if c2 == nil {
		t.Error("expecting pointer of singleton struct received ", c2)
	}

	c2.AddOne()
	newCount = c2.GetCount()
	if newCount != (2) {
		t.Error("expecting new count to be", (2), "but got", newCount)
	}
}

func TestSingletonWithMulitpleGoRoutines(t *testing.T) {
	c := New()
	done := make(chan bool, 10)
	initialCount := c.GetCount()
	for i := 0; i < 10; i++ {
		go func() {
			c = New()
			c.AddOne()
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
	if c.GetCount() != (10 + initialCount) {
		t.Error("expected count to be,", (10 + initialCount), " but got", c.GetCount())
	}
}
