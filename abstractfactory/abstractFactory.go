package abstractfactory

import "errors"

type gui interface {
	render() int
}
type guiFactory interface {
	createGUI(string) (gui, error)
}

func createGUIFactory(name string) (guiFactory, error) {
	switch name {
	case "windows":
		return new(windowsGUIFactory), nil
	case "mac":
		return new(macGUIFactory), nil
	default:
		return nil, errors.New("not supported")
	}
}
