package services

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

func TestAuthClient_Auth(t *testing.T) {
	assertNoAuth := func() {
		_, err := c.Auth.GetAuthenticatedUserID(ctx)
		assert.True(t, errors.Is(err, NotAuthenticatedError{}))
		_, err = c.Auth.GetAuthenticatedUser(ctx)
		assert.True(t, errors.Is(err, NotAuthenticatedError{}))
	}

	assertNoAuth()

	err := c.Auth.Login(ctx, usr.ID)
	require.NoError(t, err)

	uid, err := c.Auth.GetAuthenticatedUserID(ctx)
	require.NoError(t, err)
	assert.Equal(t, usr.ID, uid)

	u, err := c.Auth.GetAuthenticatedUser(ctx)
	require.NoError(t, err)
	assert.Equal(t, u.ID, usr.ID)

	err = c.Auth.Logout(ctx)
	require.NoError(t, err)

	assertNoAuth()
}

func TestAuthClient_PasswordHashing(t *testing.T) {
	pw := "testcheckpassword"
	hash, err := c.Auth.HashPassword(pw)
	assert.NoError(t, err)
	assert.NotEqual(t, hash, pw)
	err = c.Auth.CheckPassword(pw, hash)
	assert.NoError(t, err)
}

func TestAuthClient_RandomToken(t *testing.T) {
	length := c.Config.App.PasswordToken.Length
	a, err := c.Auth.RandomToken(length)
	require.NoError(t, err)
	b, err := c.Auth.RandomToken(length)
	require.NoError(t, err)
	assert.Len(t, a, length)
	assert.Len(t, b, length)
	assert.NotEqual(t, a, b)
}

func TestAuthClient_EmailVerificationToken(t *testing.T) {
	t.Run("valid token", func(t *testing.T) {
		email := "test@localhost.com"
		token, err := c.Auth.GenerateEmailVerificationToken(email)
		require.NoError(t, err)

		tokenEmail, err := c.Auth.ValidateEmailVerificationToken(token)
		require.NoError(t, err)
		assert.Equal(t, email, tokenEmail)
	})

	t.Run("invalid token", func(t *testing.T) {
		badToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RAbG9jYWxob3N0LmNvbSIsImV4cCI6MTkxNzg2NDAwMH0.ScJCpfEEzlilKfRs_aVouzwPNKI28M3AIm-hyImQHUQ"
		_, err := c.Auth.ValidateEmailVerificationToken(badToken)
		assert.Error(t, err)
	})

	t.Run("expired token", func(t *testing.T) {
		c.Config.App.EmailVerificationTokenExpiration = -time.Hour
		email := "test@localhost.com"
		token, err := c.Auth.GenerateEmailVerificationToken(email)
		require.NoError(t, err)

		_, err = c.Auth.ValidateEmailVerificationToken(token)
		assert.Error(t, err)

		c.Config.App.EmailVerificationTokenExpiration = time.Hour * 12
	})
}
