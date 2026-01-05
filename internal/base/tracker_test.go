package base_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"job4j.ru/go-lang-base/internal/base"
)

func Test_Tracker(t *testing.T) {
	t.Parallel()

	t.Run("check link leak", func(t *testing.T) {
		t.Parallel()

		tracker := base.NewTracker()

		item := base.Item{
			ID:   "1",
			Name: "First Item",
		}
		tracker.AddItem(item)

		res := tracker.GetItems()
		res[0].Name = "Second Item"

		assert.Equal(t,
			[]base.Item{item},
			tracker.GetItems(),
		)
	})

	t.Run("add few item", func(t *testing.T) {
		t.Parallel()

		tracker := base.NewTracker()

		item := base.Item{
			ID:   "1",
			Name: "First Item",
		}
		tracker.AddItem(item)
		tracker.AddItem(item)

		res := tracker.GetItems()

		assert.Equal(t,
			res[0].ID,
			res[1].ID,
		)
	})

}
