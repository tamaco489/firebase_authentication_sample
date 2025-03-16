-- Firebase Authentication で認証したユーザを作成する
-- name: CreateUserFirebaseAuthentication :exec
INSERT INTO `user_firebase_authentications` (
  `id`,
  `uid`
) VALUES (
  sqlc.arg('id'),
  sqlc.arg('uid')
);
