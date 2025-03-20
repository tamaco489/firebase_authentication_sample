-- Firebase Authentication で認証したユーザを作成する
-- name: CreateUserFirebaseAuthentication :exec
INSERT INTO `user_firebase_authentications` (
  `id`,
  `uid`
) VALUES (
  sqlc.arg('id'),
  sqlc.arg('uid')
);

-- 指定したFirebaseのユーザIDのレコードが存在しているかを判定する
-- name: GetUIDByFirebaseUID :one
SELECT uid FROM user_firebase_authentications WHERE id = sqlc.arg('firebase_uid');
