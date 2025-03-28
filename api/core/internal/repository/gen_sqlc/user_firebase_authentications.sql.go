// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: user_firebase_authentications.sql

package repository

import (
	"context"
)

const createUserFirebaseAuthentication = `-- name: CreateUserFirebaseAuthentication :exec
INSERT INTO ` + "`" + `user_firebase_authentications` + "`" + ` (
  ` + "`" + `id` + "`" + `,
  ` + "`" + `uid` + "`" + `
) VALUES (
  ?,
  ?
)
`

type CreateUserFirebaseAuthenticationParams struct {
	ID  string `json:"id"`
	Uid string `json:"uid"`
}

// Firebase Authentication で認証したユーザを作成する
func (q *Queries) CreateUserFirebaseAuthentication(ctx context.Context, db DBTX, arg CreateUserFirebaseAuthenticationParams) error {
	_, err := db.ExecContext(ctx, createUserFirebaseAuthentication, arg.ID, arg.Uid)
	return err
}

const getUIDByFirebaseUID = `-- name: GetUIDByFirebaseUID :one
SELECT uid FROM user_firebase_authentications WHERE id = ?
`

// 指定したFirebaseのユーザIDのレコードが存在しているかを判定する
func (q *Queries) GetUIDByFirebaseUID(ctx context.Context, db DBTX, firebaseUid string) (string, error) {
	row := db.QueryRowContext(ctx, getUIDByFirebaseUID, firebaseUid)
	var uid string
	err := row.Scan(&uid)
	return uid, err
}
