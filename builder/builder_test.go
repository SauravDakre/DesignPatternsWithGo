package builder

import (
	"testing"
)

func TestValidDate(t *testing.T) {
	b := builder{
		day:   29,
		month: 2,
		year:  2020,
	}
	_, err := b.new()
	if err != nil {
		t.Error("expected nil but got", err)
	}

	b = builder{
		day:   31,
		month: 1,
		year:  2019,
	}
	_, err = b.new()
	if err != nil {
		t.Error("expected nil but got", err)
	}

	b = builder{
		day:   28,
		month: 2,
		year:  2019,
	}
	_, err = b.new()
	if err != nil {
		t.Error("expected nil but got", err)
	}

	b = builder{
		day:   30,
		month: 9,
		year:  2019,
	}
	_, err = b.new()
	if err != nil {
		t.Error("expected nil but got", err)
	}
}

func TestInvalidDate(t *testing.T) {
	b := builder{
		day:   29,
		month: 2,
		year:  2019,
	}
	_, err := b.new()
	if err == nil {
		t.Error("expected err but got", err)
	}

	b = builder{
		day:   31,
		month: 2,
		year:  2019,
	}
	_, err = b.new()
	if err == nil {
		t.Error("expected err but got", err)
	}

	b = builder{
		day:   32,
		month: 1,
		year:  2019,
	}
	_, err = b.new()
	if err == nil {
		t.Error("expected err but got", err)
	}
}
