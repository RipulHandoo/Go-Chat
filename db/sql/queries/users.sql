-- name: CreateUser :one
INSERT INTO users(Email,password,username)
VALUES($1,$2,$3)
RETURNING id, Email, password, username;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE Email = $1;

-- name: DeleteUser :one
DELETE FROM users WHERE id = $1 RETURNING id, Email, password, username;

