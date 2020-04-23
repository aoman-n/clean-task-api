
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE users (
  `id` INT(11) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  `display_name` VARCHAR(32) NOT NULL,
  `login_name` VARCHAR(32) NOT NULL,
  `password_digest` VARCHAR(255) NOT NULL,
  `avatar_url` VARCHAR(255) DEFAULT "",
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  UNIQUE unique_index_users_on_login_name (login_name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF exists users;
