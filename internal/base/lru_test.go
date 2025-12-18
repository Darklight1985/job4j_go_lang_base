package base_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"job4j.ru/go-lang-base/internal/base"
)

func Test_Lru(t *testing.T) {
	t.Parallel()

	t.Run("put one node - true", func(t *testing.T) {
		t.Parallel()

		in := base.NewLruCache(6)
		in.Put("six", "six")

		assert.Equal(t, "six", *in.Get("six"))
	})

	t.Run("get from lru - true", func(t *testing.T) {
		t.Parallel()

		in := base.NewLruCache(6)
		in.Put("six", "6")
		in.Put("five", "5")
		in.Put("four", "4")
		in.Put("tree", "3")
		in.Put("two", "2")
		in.Put("one", "1")

		assert.Equal(t, "1", in.Head.Value)
		assert.Equal(t, "5", *in.Get("five"))
		assert.Equal(t, "5", in.Head.Value)
	})

	t.Run("exceed size - true", func(t *testing.T) {
		t.Parallel()

		in := base.NewLruCache(6)
		in.Put("six", "6")
		in.Put("five", "5")
		in.Put("four", "4")
		in.Put("tree", "3")
		in.Put("two", "2")
		in.Put("one", "1")
		in.Put("new", "new")

		assert.Equal(t, "5", in.Tail.Value)
		assert.Equal(t, "new", in.Head.Value)
	})
}
