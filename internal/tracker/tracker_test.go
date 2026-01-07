package tracker_test

import (
	"github.com/stretchr/testify/assert"
	"job4j.ru/go-lang-base/internal/tracker"
	"testing"
)

func Test_Tracker(t *testing.T) {
	t.Parallel()

	t.Run("error update - not found", func(t *testing.T) {

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
		err := newTracker.AddItem(itemTwo)
		assert.ErrorIs(t, err, tracker.ErrNotFound)
	})
}
