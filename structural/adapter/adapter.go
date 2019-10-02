package adapter

type coffee interface {
	cost() int
}

type basicCoffee struct {
	sugar       int
	milk        int
	coffePowder int
}

func newBasicCoffee(sugar, milk, coffePowder int) basicCoffee {
	return basicCoffee{
		sugar:       sugar,
		milk:        milk,
		coffePowder: coffePowder,
	}
}

func (b basicCoffee) cost() int {
	return b.sugar + b.milk + b.coffePowder
}

type premiumCoffee struct {
	beans int
	milk  int
	sugar int
}

func newPremiumcCoffee(sugar, milk, beans int) premiumCoffee {
	return premiumCoffee{
		sugar: sugar,
		milk:  milk,
		beans: beans,
	}
}

func (b premiumCoffee) cost(additionalCharge int) int {
	return b.sugar + b.milk + b.beans + additionalCharge
}

type premiumCoffeeAdapter struct {
	additionalCharge int
	c                premiumCoffee
}

func newPremiumcCoffeeAdapter(sugar, milk, beans, additionalCharge int) premiumCoffeeAdapter {
	return premiumCoffeeAdapter{
		additionalCharge: additionalCharge,
		c: premiumCoffee{
			sugar: sugar,
			milk:  milk,
			beans: beans,
		},
	}
}

func (b premiumCoffeeAdapter) cost() int {
	return b.c.cost(b.additionalCharge)
}

func calculateCost(c coffee) int {
	return c.cost()
}
