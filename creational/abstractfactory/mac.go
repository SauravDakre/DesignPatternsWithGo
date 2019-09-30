package abstractfactory

import "errors"

type macGUIFactory struct{}

func (w macGUIFactory) createGUI(name string) (gui, error) {
	switch name {
	case "button":
		return new(macButton), nil
	case "checkbox":
		return new(macCheckbox), nil
	default:
		return nil, errors.New("not supported")
	}
}

type macButton struct{}

func (w macButton) render() int {
	return 3
}

type macCheckbox struct{}

func (w macCheckbox) render() int {
	return 4
}
