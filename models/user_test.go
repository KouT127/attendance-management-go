package models

import (
	"github.com/Songmu/flextime"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser(t *testing.T) {
	t.Run("Testing User data access", func(t *testing.T) {
		assert.Nil(t, SetTestDatabase())
		userId := "eru4kawei42lasedi356ladfkjfity3"

		t.Run("Should create user", func(t *testing.T) {
			user, err := GetOrCreateUser(userId)
			assert.Nil(t, err)
			assert.Equal(t, userId, user.Id)
		})

		t.Run("Should not create user when id is empty", func(t *testing.T) {
			user, err := GetOrCreateUser("")
			assert.NotNil(t, err)
			assert.Nil(t, user)
		})

		t.Run("Should get user", func(t *testing.T) {
			user, err := GetOrCreateUser(userId)
			assert.Nil(t, err)
			assert.Equal(t, userId, user.Id)
		})

		t.Run("Should update user", func(t *testing.T) {
			userId := "asdiekawei42lasedi356ladfkjfity3"
			user, err := createTestUser(userId)
			assert.Nil(t, err)
			assert.NotNil(t, user)

			user.Name = "updated"
			user.Email = "updated@test.com"
			user.ImageUrl = "updated.com"
			assert.Nil(t, UpdateUser(user))

			gotUser, err := getUser(engine, userId)
			assert.Nil(t, err)
			assert.Equal(t, gotUser.Name, "updated")
			assert.Equal(t, gotUser.Email, "updated@test.com")
			assert.Equal(t, gotUser.ImageUrl, "updated.com")
		})

		t.Run("Should not update user", func(t *testing.T) {
			u := &User{
				Id: "",
			}
			assert.NotNil(t, UpdateUser(u))
		})
	})
}

func createTestUser(userId string) (*User, error) {
	u := &User{
		Id:        userId,
		Name:      "insert user",
		Email:     "insert@test.com",
		ImageUrl:  "insert.com",
		CreatedAt: flextime.Now(),
		UpdatedAt: flextime.Now(),
	}

	if _, err := engine.Insert(u); err != nil {
		return nil, err
	}
	return u, nil
}
