
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE tasks (
  `id` INT(11) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  `name` VARCHAR(255) NOT NULL,
  `due_on` DATETIME DEFAULT NULL,
  `status` VARCHAR(32) NOT NULL DEFAULT "waiting",
  `project_id` INT(11) UNSIGNED NOT NULL,
  CONSTRAINT tasks_fk_project_id
    FOREIGN KEY (project_id)
    REFERENCES projects (id)
    ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF exists tasks;
