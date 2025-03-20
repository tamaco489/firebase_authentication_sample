// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package repository

import (
	"context"
)

type Querier interface {
	// ユーザーを作成する
	CreateUser(ctx context.Context, db DBTX, arg CreateUserParams) error
	// Firebase Authentication で認証したユーザを作成する
	CreateUserFirebaseAuthentication(ctx context.Context, db DBTX, arg CreateUserFirebaseAuthenticationParams) error
	// 指定したFirebaseのユーザIDのレコードが存在しているかを判定する
	GetUIDByFirebaseUID(ctx context.Context, db DBTX, firebaseUid string) (string, error)
}

var _ Querier = (*Queries)(nil)
