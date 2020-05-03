
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE tags (
  `id` INT(11) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  `name` VARCHAR(30) NOT NULL,
  `color` VARCHAR(7) NOT NULL,
  `project_id` INT(11) UNSIGNED NOT NULL,
  CONSTRAINT tags_fk_project_id
    FOREIGN KEY (project_id)
    REFERENCES projects (id)
    ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS tags;
