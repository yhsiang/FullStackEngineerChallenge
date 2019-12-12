-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE `employees` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back


DROP TABLE `employees`;
