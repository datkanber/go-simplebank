package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := RandomString(6)

	HashedPassword1, err := HashedPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, HashedPassword1)

	err = CheckPassword(password, HashedPassword1)
	require.NoError(t, err)

	wrongPassword := RandomString(6)
	err = CheckPassword(wrongPassword, HashedPassword1)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	HashedPassword2, err := HashedPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, HashedPassword2)
	require.NotEqual(t, HashedPassword1, HashedPassword2)
}
