package base_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"job4j.ru/go-lang-base/internal/base"
)

func Test_Validate(t *testing.T) {
	t.Parallel()

	t.Run("validate Request not empty", func(t *testing.T) {
		t.Parallel()

		var first *base.ValidateRequest

		first = &base.ValidateRequest{
			Title:       "Title",
			Description: "Description",
		}

		rsl := base.Validate(first)

		assert.Equal(t, []string{"no UserID"}, rsl)
	})

	t.Run("validate request empty", func(t *testing.T) {
		t.Parallel()

		var first *base.ValidateRequest

		first = &base.ValidateRequest{}

		rsl := base.Validate(first)

		assert.Equal(t, []string{"no UserID", "no Title", "no Description"}, rsl)
	})

	t.Run("validate request nil", func(t *testing.T) {
		t.Parallel()

		var first *base.ValidateRequest

		rsl := base.Validate(first)

		assert.Equal(t, []string{"no data"}, rsl)
	})
}
