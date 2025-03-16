
-- +migrate Up
CREATE TABLE IF NOT EXISTS `user_firebase_authentications` (
  `id` VARCHAR(255) PRIMARY KEY COMMENT '外部認証プロバイダのユニークなID',
  `uid` VARCHAR(255) NOT NULL COMMENT 'usersテーブルのID',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (`uid`) REFERENCES `users`(`id`)
);

-- +migrate Down
DROP TABLE IF EXISTS `user_firebase_authentications`;
