-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE `review_assignments` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `reviewee` int(11) NOT NULL,
  `reviewer` int(11) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`reviewee`) REFERENCES employees(`id`),
  FOREIGN KEY (`reviewer`) REFERENCES employees(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE `review_assignments`;