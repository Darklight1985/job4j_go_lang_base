package tracker

import "github.com/google/uuid"

type Usecase interface {
	Done(in Input, out Output, tracker *Tracker)
}

type AddUsecase struct{}

func (u AddUsecase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("enter name:")
	name := in.Get()
	id := uuid.New().String()
	tracker.AddItem(Item{Name: name, ID: id})
}

type GetUsecase struct{}

func (u GetUsecase) Done(_ Input, out Output, tracker *Tracker) {
	for _, item := range tracker.Items {
		out.Out(item.toString())
	}
}

type DeleteUsecase struct{}

func (u DeleteUsecase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("enter name item:")
	name := in.Get()
	tracker.DeleteItem(out, name)
}

type UpdateUsecase struct{}

func (u UpdateUsecase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("enter id item:")
	id := in.Get()
	out.Out("enter new name item:")
	name := in.Get()
	tracker.UpdateItem(out, id, name)
}

type GetItemUsecase struct{}

func (u GetItemUsecase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("enter part of name item:")
	name := in.Get()
	tracker.GetItem(out, name)
}
