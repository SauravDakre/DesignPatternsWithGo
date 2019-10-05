package facade

import (
	"errors"
)

type menu struct {
	id       int
	name     string
	prepared bool
}

type table struct {
	id            int
	menus         []menu
	menuDelivered bool
}

var orders = make(map[int]*table)

type waiter struct {
}

func newWaiter() waiter {
	return waiter{}
}

func (w waiter) getOrders() *table {
	m := make([]menu, 0)
	m = append(m, menu{
		id:       1,
		name:     "dal khichdi",
		prepared: false,
	})
	t := table{
		id:            1,
		menus:         m,
		menuDelivered: false,
	}
	orders[t.id] = &t
	return &t
}

func (w *waiter) deliverOrders(t *table) error {
	allItemPrepared := true
	for k := range orders {
		for _, menu := range orders[k].menus {
			allItemPrepared = allItemPrepared && menu.prepared
		}
		orders[k].menuDelivered = allItemPrepared
	}
	if !allItemPrepared {
		return errors.New("menu not prepared")
	}
	t.menuDelivered = true
	return nil
}

type kitchen struct {
}

func (k kitchen) createFood() {
	for _, v := range orders {
		for i := 0; i < len(v.menus); i++ {
			v.menus[i].prepared = true
		}

	}
}

type managerFacade struct {
	w *waiter
	k *kitchen
}

func newManagerFacade() managerFacade {
	return managerFacade{
		w: &waiter{},
		k: &kitchen{},
	}
}

func (m managerFacade) manage() error {
	t := m.w.getOrders()
	m.k.createFood()
	return m.w.deliverOrders(t)
}
