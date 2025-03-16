-- ユーザーを作成する
-- name: CreateUser :exec
INSERT INTO `users` (
  `id`,
  `username`,
  `email`,
  `role`,
  `status`,
  `last_login_at`
) VALUES (
  sqlc.arg('id'),
  sqlc.arg('username'),
  sqlc.arg('email'),
  sqlc.arg('role'),
  sqlc.arg('status'),
  sqlc.arg('last_login_at')
);

-- uidを指定して対象のユーザ情報を取得する
-- name: GetUserByID :one
SELECT
  `username`,
  `email`,
  `role`,
  `status`,
  `last_login_at`
FROM `users` WHERE `id` = ?;
