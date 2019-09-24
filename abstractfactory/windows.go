package abstractfactory

import "errors"

type windowsGUIFactory struct{}

func (w windowsGUIFactory) createGUI(name string) (gui, error) {
	switch name {
	case "button":
		return new(windowsButton), nil
	case "checkbox":
		return new(windowsChecbox), nil
	default:
		return nil, errors.New("not supported")
	}
}

type windowsButton struct{}

func (w windowsButton) render() int {
	return 1
}

type windowsChecbox struct{}

func (w windowsChecbox) render() int {
	return 2
}
