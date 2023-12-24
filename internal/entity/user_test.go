package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	email := "dave.mustaine@hotmail.com"
	pass := "SuperSecurePa$$word"

	user, err := NewUser(email, pass)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.Equal(t, user.Email, email)
	assert.NotEmpty(t, user.Password)
}

func TestValidatePassword(t *testing.T) {
	email := "dave.mustaine@hotmail.com"
	pass := "SuperSecurePa$$word"

	user, err := NewUser(email, pass)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.True(t, user.ValidatePassword(pass))
	assert.False(t, user.ValidatePassword("anotherPass"))
}
