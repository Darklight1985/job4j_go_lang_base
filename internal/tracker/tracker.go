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

func (t *Tracker) AddItem(item Item) (Item, error) {
	_, ok := t.indexOf(item.ID)
	if ok {
		return Item{}, ErrAlreadyExists
	}
	t.Items = append(t.Items, item)
	return item, nil
}

func (t *Tracker) GetItems() []Item {
	res := make([]Item, len(t.Items))
	copy(res, t.Items)
	return res
}

func (t *Tracker) DeleteItem(id string) error {
	items := t.Items
	i, ok := t.indexOf(id)
	if !ok {

		return ErrNotFound
	}
	t.Items = append(items[:i], items[i+1:]...)
	return nil
}

func (t *Tracker) UpdateItem(id string, name string) error {
	items := t.Items
	i, ok := t.indexOf(id)
	if !ok {
		return ErrNotFound
	}
	items[i].Name = name
	return nil
}

func (t *Tracker) GetItem(subStr string) (Item, error) {
	items := t.Items
	for _, item := range items {
		if strings.Contains(item.Name, subStr) {
			return item, nil
		}
	}
	return Item{}, ErrNotFound
}

func (t *Tracker) indexOf(id string) (int, bool) {
	for i, item := range t.Items {
		if item.ID == id {
			return i, true
		}
	}
	return -1, false
}
