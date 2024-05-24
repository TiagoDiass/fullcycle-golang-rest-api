package entity

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser(
		"Tiago",
		"tiago@gmail.com",
		"fake-password",
	)

	require.Nil(t, err)
	require.NotEmpty(t, user.ID)
	require.Equal(t, user.Name, "Tiago")
	require.Equal(t, user.Email, "tiago@gmail.com")
	require.NotEmpty(t, user.Password)
	require.NotEqual(t, user.Password, "fake-password")
}

func TestValidatePassword(t *testing.T) {
	user, _ := NewUser(
		"Tiago",
		"tiago@gmail.com",
		"my-beautiful-password",
	)

	passwordsMatch := user.ValidatePassword("wrong-password")
	require.Equal(t, passwordsMatch, false)

	passwordsMatch = user.ValidatePassword("my-beautiful-password")
	require.Equal(t, passwordsMatch, true)
}
