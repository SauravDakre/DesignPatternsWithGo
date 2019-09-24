package prototype

type node interface {
	Clone() node
	Name() string
	Child() []node
}

type element struct {
	name  string
	child []node
}

func (e element) Clone() node {
	n := element{}
	n.name = e.name
	n.child = make([]node, 0)
	for _, v := range e.child {
		n.child = append(n.child, v)
	}
	return n
}

func (e element) Name() string {
	return e.name
}

func (e element) Child() []node {
	return e.child
}
