
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE project_users (
  `id` INT(11) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  `user_id` INT(11) UNSIGNED NOT NULL,
  `project_id` INT(11) UNSIGNED NOT NULL,
  `role` VARCHAR(32) NOT NULL,
  UNIQUE unique_index_users_on_user_id_and_project_id (user_id, project_id),
  CONSTRAINT project_users_fk_user_id
    FOREIGN KEY (user_id)
    REFERENCES users (id)
    ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT project_users_fk_project_id
    FOREIGN KEY (project_id)
    REFERENCES projects (id)
    ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF exists project_users;
