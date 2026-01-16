package tracker

import (
	"fmt"
	"github.com/google/uuid"
)

type Usecase interface {
	Done(in Input, out Output, tracker *Tracker)
}

type AddUsecase struct{}

func (u AddUsecase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("enter name:")
	name := in.Get()
	id := uuid.New().String()
	_, error := tracker.AddItem(Item{Name: name, ID: id})
	if error != nil {
		out.Out("item with such id already exists")
	}
}

type GetUsecase struct{}

func (u GetUsecase) Done(_ Input, out Output, tracker *Tracker) {
	for _, item := range tracker.Items {
		out.Out(item.toString())
	}
}

type DeleteUsecase struct{}

func (u DeleteUsecase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("enter id item:")
	name := in.Get()
	if !tracker.DeleteItem(name) {
		out.Out("not found item")
	}
}

type UpdateUsecase struct{}

func (u UpdateUsecase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("enter id item:")
	id := in.Get()
	out.Out("enter new name item:")
	name := in.Get()
	if !tracker.UpdateItem(id, name) {
		out.Out("not found item")
	}
}

type GetItemUsecase struct{}

func (u GetItemUsecase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("enter part of name item:")
	name := in.Get()
	error, item := tracker.GetItem(name)
	if error != nil {
		out.Out("not found item")
	}
	out.Out(fmt.Sprintf("item ID: %s Name: %s", item.ID, item.Name))
}
