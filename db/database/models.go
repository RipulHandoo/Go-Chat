// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package database

import ()

type User struct {
	ID       int64
	Email    string
	Password string
	Username string
}

type UserFollower struct {
	FollowerID  int64
	FollowingID int64
}
