package article

import (
	"errors"
	"testing"

	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/mock"
	assert "github.com/stretchr/testify/require"
)

func TestArticle_NewController(t *testing.T) {
	t.Parallel()
	t.Run("when Controller is created", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		r := new(MockRepository)

		controller := NewController(r)
		is.NotNil(controller)
	})
}

func TestArticle_Create(t *testing.T) {
	t.Parallel()
	t.Run("when article creation fails", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		r := new(MockRepository)
		defer r.AssertExpectations(t)

		r.On("Create", mock.Anything).
			Return(errors.New("error")).
			Once()

		c := NewController(r)

		article := &Article{}
		is.Nil(faker.FakeData(article))
		is.NotNil(c.Create(article))
	})
	t.Run("when article creation succeeds", func(t *testing.T) {
		t.Parallel()
	})
}
