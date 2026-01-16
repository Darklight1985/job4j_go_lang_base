package tracker

import (
	"strings"
)

type Tracker struct {
	Items []Item
}

func NewTracker() *Tracker {
	return &Tracker{}
}

func (t *Tracker) AddItem(item Item) error {
	_, ok := t.indexOf(item.ID)
	if ok {
		return ErrNotFound
	}

	t.Items = append(t.Items, item)
	return nil
}

func (t *Tracker) GetItems() []Item {
	res := make([]Item, len(t.Items))
	copy(res, t.Items)
	return res
}

func (t *Tracker) DeleteItem(id string) bool {
	items := t.Items
	i, ok := t.indexOf(id)
	if ok {
		t.Items = append(items[:i], items[i+1:]...)
		return true
	}
	return false
}

func (t *Tracker) UpdateItem(id string, name string) bool {
	items := t.Items
	i, ok := t.indexOf(id)
	if ok {
		items[i].Name = name
		return true
	}
	return false
}

func (t *Tracker) GetItem(subStr string) (error, Item) {
	items := t.Items
	for _, item := range items {
		if strings.Contains(item.Name, subStr) {
			return nil, item
		}
	}
	return ErrNotFound, Item{}
}

func (t *Tracker) indexOf(id string) (int, bool) {
	for i, item := range t.Items {
		if item.ID == id {
			return i, true
		}
	}
	return -1, false
}
