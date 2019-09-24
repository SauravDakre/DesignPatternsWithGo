package abstractfactory

import (
	"testing"
)

func TestMacButton(t *testing.T) {
	gf, err := createGUIFactory("mac")
	if err != nil {
		t.Error("Expected error to be nil but received", err)
	}

	b, err := gf.createGUI("button")
	if err != nil {
		t.Error("Expected error to be nil but received", err)
	}

	v := b.render()
	if v != 3 {
		t.Error("Expected value to be 3 but received", v)
	}
}

func TestWindowsButton(t *testing.T) {
	gf, err := createGUIFactory("windows")
	if err != nil {
		t.Error("Expected error to be nil but received", err)
	}

	b, err := gf.createGUI("button")
	if err != nil {
		t.Error("Expected error to be nil but received", err)
	}

	v := b.render()
	if v != 1 {
		t.Error("Expected value to be 1 but received", v)
	}
}

func TestWindowsCheckbox(t *testing.T) {
	gf, err := createGUIFactory("windows")
	if err != nil {
		t.Error("Expected error to be nil but received", err)
	}

	b, err := gf.createGUI("checkbox")
	if err != nil {
		t.Error("Expected error to be nil but received", err)
	}

	v := b.render()
	if v != 2 {
		t.Error("Expected value to be 2 but received", v)
	}
}

func TestMacCheckbox(t *testing.T) {
	gf, err := createGUIFactory("mac")
	if err != nil {
		t.Error("Expected error to be nil but received", err)
	}

	b, err := gf.createGUI("checkbox")
	if err != nil {
		t.Error("Expected error to be nil but received", err)
	}

	v := b.render()
	if v != 4 {
		t.Error("Expected value to be 4 but received", v)
	}
}
