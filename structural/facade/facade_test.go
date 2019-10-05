package facade

import "testing"

func TestFacade(t *testing.T) {
	mf := newManagerFacade()
	err := mf.manage()
	if err != nil {
		t.Error("expected error to be nil but got", err)
	}

	order, ok := orders[1]
	if !ok {
		t.Error("expected order to be present")
	}

	if len(order.menus) != 1 {
		t.Error("expected menu to be present")
	}

	m := order.menus[0]
	if m.name != "dal khichdi" {
		t.Error("expected menu to be dal khichdi but got ", m.name)
	}

	if !order.menuDelivered {
		t.Error("expected menu to be delivered")
	}
}
