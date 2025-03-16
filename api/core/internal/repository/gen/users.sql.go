// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: users.sql

package repository

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :exec
INSERT INTO ` + "`" + `users` + "`" + ` (
  ` + "`" + `id` + "`" + `,
  ` + "`" + `username` + "`" + `,
  ` + "`" + `email` + "`" + `,
  ` + "`" + `role` + "`" + `,
  ` + "`" + `status` + "`" + `,
  ` + "`" + `last_login_at` + "`" + `
) VALUES (
  ?,
  ?,
  ?,
  ?,
  ?,
  ?
)
`

type CreateUserParams struct {
	ID          string       `json:"id"`
	Username    string       `json:"username"`
	Email       string       `json:"email"`
	Role        UsersRole    `json:"role"`
	Status      UsersStatus  `json:"status"`
	LastLoginAt sql.NullTime `json:"last_login_at"`
}

// ユーザーを作成する
//
//	INSERT INTO `users` (
//	  `id`,
//	  `username`,
//	  `email`,
//	  `role`,
//	  `status`,
//	  `last_login_at`
//	) VALUES (
//	  ?,
//	  ?,
//	  ?,
//	  ?,
//	  ?,
//	  ?
//	)
func (q *Queries) CreateUser(ctx context.Context, db DBTX, arg CreateUserParams) error {
	_, err := db.ExecContext(ctx, createUser,
		arg.ID,
		arg.Username,
		arg.Email,
		arg.Role,
		arg.Status,
		arg.LastLoginAt,
	)
	return err
}

const getUserByID = `-- name: GetUserByID :one
SELECT
  ` + "`" + `username` + "`" + `,
  ` + "`" + `email` + "`" + `,
  ` + "`" + `role` + "`" + `,
  ` + "`" + `status` + "`" + `,
  ` + "`" + `last_login_at` + "`" + `
FROM ` + "`" + `users` + "`" + ` WHERE ` + "`" + `id` + "`" + ` = ?
`

type GetUserByIDRow struct {
	Username    string       `json:"username"`
	Email       string       `json:"email"`
	Role        UsersRole    `json:"role"`
	Status      UsersStatus  `json:"status"`
	LastLoginAt sql.NullTime `json:"last_login_at"`
}

// uidを指定して対象のユーザ情報を取得する
//
//	SELECT
//	  `username`,
//	  `email`,
//	  `role`,
//	  `status`,
//	  `last_login_at`
//	FROM `users` WHERE `id` = ?
func (q *Queries) GetUserByID(ctx context.Context, db DBTX, id string) (GetUserByIDRow, error) {
	row := db.QueryRowContext(ctx, getUserByID, id)
	var i GetUserByIDRow
	err := row.Scan(
		&i.Username,
		&i.Email,
		&i.Role,
		&i.Status,
		&i.LastLoginAt,
	)
	return i, err
}
