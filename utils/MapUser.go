package utils

import "github.com/RipulHandoo/goChat/db/database"

type RegisterUser struct {
	ID       int64
	Email    string
	Password string
	Username string
}

func MapRegisteredUser(user database.User) RegisterUser {
	return RegisterUser{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
		Username: user.Username,
	}
}

type LoginUser struct {
	ID       int64
	Email    string
	Username string
}

func MapLoginUser(user database.User) LoginUser {
	return LoginUser{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
	}
}

type DeleteUser struct {
	ID       int64
	Email    string
	Username string
}

func MapDeleteUser(user database.User) DeleteUser {
	return DeleteUser{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
	}
}
