package base_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"job4j.ru/go-lang-base/internal/base"
)

func Test_Lru(t *testing.T) {
	t.Parallel()

	t.Run("[1, 2, 3] - true", func(t *testing.T) {
		t.Parallel()

		in := base.NewLruCache(6)
		in.Put("six", "six")

		assert.Equal(t, "six", *in.Get("six"))
	})

	t.Run("[1, 2, 4] - true", func(t *testing.T) {
		t.Parallel()

		in := base.NewLruCache(6)
		in.Put("six", "six")
		in.Put("five", "five")
		in.Put("four", "four")
		in.Put("tree", "tree")
		in.Put("two", "two")
		in.Put("one", "one")

		assert.Equal(t, "five", *in.Get("five"))
	})
}
