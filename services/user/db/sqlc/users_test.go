package user

import (
	"context"
	"database/sql"
	"testing"

	"github.com/omkarbhostekar/brewgo/services/user/util"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	CreateRandomUser(t)
}

func CreateRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Name: util.RandomString(6),
		Email: util.RandomEmail(),
		Password: util.RandomPassword(),
		PhoneNumber: util.RandomPhoneNumber(),
	}
	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.Name, user.Name)
	require.Equal(t, arg.Password, user.Password)
	require.Equal(t, arg.Email, user.Email)
	require.Zero(t, user.PasswordChangedAt)
	require.NotZero(t, user.CreatedAt)
	return user
}

func TestGetUserByPhoneNumber(t *testing.T) {
	user1 := CreateRandomUser(t)
	user2, err := testQueries.GetUserByPhoneNumber(context.Background(), user1.PhoneNumber)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Name, user2.Name)
	require.Equal(t, user1.Password, user2.Password)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.PhoneNumber, user2.PhoneNumber)
	require.Equal(t, user1.CreatedAt, user2.CreatedAt)
	require.Equal(t, user1.PasswordChangedAt, user2.PasswordChangedAt)
}

func TestGetUserByPhoneNumberNotFound(t *testing.T) {
	user, err := testQueries.GetUserByPhoneNumber(context.Background(), util.RandomPhoneNumber())
	require.Error(t, err)
	require.Empty(t, user)
}

func TestUpdateUser(t *testing.T) {
	user1 := CreateRandomUser(t)
	arg := UpdateUserParams{
		ID: user1.ID,
		Name: sql.NullString{String: util.RandomString(6), Valid: true},
		Email: sql.NullString{String: util.RandomEmail(), Valid: true},
		PhoneNumber: sql.NullString{String: util.RandomPhoneNumber(), Valid: true},
	}
	user2, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, arg.Name.String, user2.Name)
	require.Equal(t, arg.Email.String, user2.Email)
	require.Equal(t, arg.PhoneNumber.String, user2.PhoneNumber)
	require.Equal(t, user1.Password, user2.Password)
	require.Equal(t, user1.CreatedAt, user2.CreatedAt)
	require.NotEqual(t, user1.Name, user2.Name)
	require.NotEqual(t, user1.Email, user2.Email)
	require.NotEqual(t, user1.PhoneNumber, user2.PhoneNumber)
}