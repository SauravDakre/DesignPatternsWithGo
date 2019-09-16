package factory

import "testing"

func TestSquarePolygon(t *testing.T) {
	registerFactories("square", newSquare)
	p, err := createPolygon("square", 5)
	if err != nil {
		t.Error("Expected error to be nil but got", err)
	}

	area := p.area()
	if area != 25 {
		t.Error("Expected area to be 25 but got", area)
	}
}

func TestRectanglePolygon(t *testing.T) {
	registerFactories("rectangle", newRectangle)
	p, err := createPolygon("rectangle", 5, 7)
	if err != nil {
		t.Error("Expected error to be nil but got", err)
	}

	area := p.area()
	if area != 35 {
		t.Error("Expected area to be 25 but got", area)
	}
}

func TestRegisterFactories(t *testing.T) {
	err := registerFactories("notexists", nil)
	if err == nil {
		t.Error("Expected error but received nil")
	}

	err = registerFactories("rectangle", newRectangle)
	if err == nil {
		t.Error("Expected error but received nil")
	}
}
