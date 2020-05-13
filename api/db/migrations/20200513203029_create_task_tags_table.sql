
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE task_tags (
  `id` INT(11) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  `task_id` INT(11) UNSIGNED NOT NULL,
  `tag_id` INT(11) UNSIGNED NOT NULL,
  CONSTRAINT task_tags_fk_task_id
    FOREIGN KEY (task_id)
    REFERENCES tasks (id)
    ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT task_tags_fk_tag_id
    FOREIGN KEY (tag_id)
    REFERENCES tags (id)
    ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS task_tags;
