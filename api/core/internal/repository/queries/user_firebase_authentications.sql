-- ユーザーを作成するクエリ
-- name: CreateUserFirebaseAuthentication :exec
INSERT INTO `user_firebase_authentications` (
  `id`,
  `uid`
) VALUES (
  sqlc.arg('id'),
  sqlc.arg('uid')
);
