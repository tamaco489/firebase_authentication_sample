
-- +migrate Up
CREATE TABLE IF NOT EXISTS `users` (
  `id` VARCHAR(255) PRIMARY KEY COMMENT 'プロダクト固有のユニークなID',
  `username` VARCHAR(255) COMMENT 'ゲーム内のアカウント名',
  `email` VARCHAR(255) COMMENT 'ユーザのメールアドレス',
  `role` ENUM('general', 'admin', 'beta_tester') NOT NULL COMMENT 'ユーザの権限レベル',
  `status` ENUM('active', 'inactive', 'banned') NOT NULL COMMENT 'アカウントが有効か、無効か、強制退会済みかを判別',
  `last_login_at` DATETIME NOT NULL COMMENT '最終ログイン日時',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  INDEX `idx_email` (`email`),
  INDEX `idx_role` (`role`),
  INDEX `idx_status` (`status`)
);

-- +migrate Down
DROP TABLE IF EXISTS `users`;
