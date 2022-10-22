// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: user.sql

package maindb

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (email, password, role)
VALUES ($1, $2, $3)
RETURNING id, email, password, role, created_at
`

type CreateUserParams struct {
	Email    string
	Password string
	Role     string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (*User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Email, arg.Password, arg.Role)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.Role,
		&i.CreatedAt,
	)
	return &i, err
}

const getUser = `-- name: GetUser :one
SELECT id, email, password, role, created_at FROM users
WHERE id = $1
`

func (q *Queries) GetUser(ctx context.Context, id int32) (*User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.Role,
		&i.CreatedAt,
	)
	return &i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, email, password, role, created_at FROM users
WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.Role,
		&i.CreatedAt,
	)
	return &i, err
}
