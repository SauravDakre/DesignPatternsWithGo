package bridge

type movelogic interface {
	move() error
}

type animal struct {
	animalType int
	movelogic
}

func newAnimal(animalType int, logic movelogic) animal {
	return animal{
		animalType: animalType,
		movelogic:  logic,
	}
}

const (
	human = iota
	fish
)

type walk struct{}

func (w walk) move() error {
	return nil
}

func newWalk() walk {
	return walk{}
}

type swim struct{}

func (s swim) move() error {
	return nil
}

func newSwim() swim {
	return swim{}
}
