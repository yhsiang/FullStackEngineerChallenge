-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE `performance_reviews` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `content` TEXT NOT NULL,
  `assign_id` int(11) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
FOREIGN KEY (`assign_id`) REFERENCES review_assignments(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE `performance_reviews`;