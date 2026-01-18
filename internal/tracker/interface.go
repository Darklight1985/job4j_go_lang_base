package tracker

import "context"

type Input interface {
	Get() string
}

type Output interface {
	Out(text string)
}

type Store interface {
	Create(ctx context.Context, item Item) error
	List(ctx context.Context) ([]Item, error)
	Get(ctx context.Context, id string) (Item, error)
}
