package store_test

import (
	"testing"

	"github.com/Fox1N69/rest-api/internal/app/model"
	"github.com/Fox1N69/rest-api/store"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T){
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	u, err := s.User().Create(&model.User{
		Email: "user@example.org",
	})
	assert.NoError(t, err)
	assert.NotNil(t, u)
}