package factory

import "errors"

type polygon interface {
	area() int
}

type rectangle struct {
	l, b int
}

type square struct {
	a int
}

type polygonFactory func(params ...interface{}) polygon

func newSquare(params ...interface{}) polygon {
	return square{
		a: params[0].(int),
	}
}

func (s square) area() int {
	return s.a * s.a
}

func newRectangle(params ...interface{}) polygon {
	return rectangle{
		l: params[0].(int),
		b: params[1].(int),
	}
}

func (r rectangle) area() int {
	return r.l * r.b
}

var polygonFactories = make(map[string]polygonFactory)

func registerFactories(name string, factory polygonFactory) (err error) {
	if factory == nil {
		err = errors.New("polygon factory for " + name + " doesn't exist")
		return
	}
	_, exists := polygonFactories[name]
	if exists {
		err = errors.New("polygon factory for " + name + " already exist")
		return
	}
	polygonFactories[name] = factory
	return
}

func createPolygon(params ...interface{}) (polygon, error) {
	polygonName := params[0].(string)

	factory, ok := polygonFactories[polygonName]
	if !ok {
		err := errors.New("polygon factory for " + polygonName + " doesn't exist")
		return nil, err
	}

	p := factory(params[1:]...)
	return p, nil
}
