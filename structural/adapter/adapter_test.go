package adapter

import "testing"

func TestAdapterPattern(t *testing.T) {
	basic := newBasicCoffee(5, 10, 7)
	premium := newPremiumcCoffeeAdapter(5, 10, 10, 5)

	cost := calculateCost(basic)
	if cost != 22 {
		t.Error("expected cost to be 22 but got", cost)
	}

	cost = calculateCost(premium)
	if cost != 30 {
		t.Error("expected cost to be 30 but got", cost)
	}
}
