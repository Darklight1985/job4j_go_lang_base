package tracker

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ConsoleInput struct{}

func (c ConsoleInput) Get() string {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		// block IO
	}
	return scanner.Text()
}

type ConsoleOutput struct{}

func (c ConsoleOutput) Out(text string) {
	fmt.Println(text)
}

type Item struct {
	ID   string
	Name string
}

func (i Item) toString() string {
	return fmt.Sprintf("%s\t%s", i.ID, i.Name)
}

type Tracker struct {
	Items []Item
}

func NewTracker() *Tracker {
	return &Tracker{}
}

func (t *Tracker) AddItem(item Item) {
	t.Items = append(t.Items, item)
}

func (t *Tracker) GetItems() []Item {
	return t.Items
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

func (t *Tracker) GetItem(out Output, subStr string) {
	items := t.Items
	for _, item := range items {
		if strings.Contains(item.Name, subStr) {
			out.Out(fmt.Sprintf("found item ID: %s  Name: %s", item.ID, item.Name))
			return
		}
	}
	out.Out("not found item")
}
