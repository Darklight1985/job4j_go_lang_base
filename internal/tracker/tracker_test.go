package tracker_test

import (
	"github.com/stretchr/testify/assert"
	"job4j.ru/go-lang-base/internal/tracker"
	"testing"
)

func Test_Tracker(t *testing.T) {
	t.Parallel()

	t.Run("item add - successfully", func(t *testing.T) {
		t.Parallel()

		newTracker := tracker.NewTracker()
		item := tracker.Item{
			ID:   "1",
			Name: "First Item",
		}

		itemTwo := tracker.Item{
			ID:   "2",
			Name: "Second Item",
		}

		newTracker.AddItem(item)
		itemAnswer, _ := newTracker.AddItem(itemTwo)
		assert.Equal(t, itemTwo, itemAnswer)
	})

	t.Run("item add - error already exists", func(t *testing.T) {
		t.Parallel()

		newTracker := tracker.NewTracker()
		item := tracker.Item{
			ID:   "1",
			Name: "First Item",
		}

		itemTwo := tracker.Item{
			ID:   "1",
			Name: "Second Item",
		}

		newTracker.AddItem(item)
		_, err := newTracker.AddItem(itemTwo)
		assert.ErrorIs(t, err, tracker.ErrAlreadyExists)
	})
}
