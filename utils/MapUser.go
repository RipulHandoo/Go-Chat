package utils

import "github.com/RipulHandoo/goChat/db/database"

// RegisterUser represents a user for registration purposes.
type RegisterUser struct {
	ID       int64
	Email    string
	Password string
	Username string
}

// MapRegisteredUser maps a database.User to a RegisterUser structure.
func MapRegisteredUser(user database.User) RegisterUser {
	return RegisterUser{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
		Username: user.Username,
	}
}

// LoginUser represents a user for login purposes.
type LoginUser struct {
	ID       int64
	Email    string
	Username string
}

// MapLoginUser maps a database.User to a LoginUser structure.
func MapLoginUser(user database.User) LoginUser {
	return LoginUser{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
	}
}

// DeleteUser represents a user for deletion purposes.
type DeleteUser struct {
	ID       int64
	Email    string
	Username string
}

// MapDeleteUser maps a database.User to a DeleteUser structure.
func MapDeleteUser(user database.User) DeleteUser {
	return DeleteUser{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
	}
}
