package decorator

import "testing"

func TestDecorator(t *testing.T) {
	s := "world"

	var fn stringManipulator = identity
	fn = toCamelCase(fn)
	r := fn(s)
	if r != "World" {
		t.Error("expected World got", r)
	}

	fn = identity
	fn = appendPrefix(toCamelCase(fn))
	r = fn(s)
	if r != "Hello World" {
		t.Error("expected Hello World got", r)
	}

	fn = identity
	fn = appendCustomPrefix("play")(toCamelCase(fn))
	r = fn("golang")
	if r != "Play Golang" {
		t.Error("expected Play Golang got", r)
	}
}
