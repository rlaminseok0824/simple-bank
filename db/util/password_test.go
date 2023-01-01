package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := RandomString(6)

	hasshedPassword1, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hasshedPassword1)

	err = CheckPassword(password, hasshedPassword1)
	require.NoError(t, err)

	wrongPassword := RandomString(6)
	err = CheckPassword(wrongPassword, hasshedPassword1)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	hasshedPassword2, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hasshedPassword2)
	require.NotEqual(t, hasshedPassword1, hasshedPassword2)
}
