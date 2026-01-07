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

func (t *Tracker) DeleteItem(out Output, name string) {
	items := t.Items
	for i, item := range items {
		if item.Name == name {
			t.Items = append(items[:i], items[i+1:]...)
			return
		}
	}
	out.Out("not found item")
}

func (t *Tracker) UpdateItem(out Output, id string, name string) {
	items := t.Items
	for _, item := range items {
		if item.ID == id {
			item.Name = name
			return
		}
	}
	out.Out("not found item")
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
