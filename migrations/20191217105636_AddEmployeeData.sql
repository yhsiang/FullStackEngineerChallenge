-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

INSERT INTO employees (`name`, `updated_at`, `created_at`) VALUES ("test1", "2019-12-16 09:49:18", "2019-12-16 09:49:18");
INSERT INTO employees (`name`, `updated_at`, `created_at`) VALUES ("test2", "2019-12-16 09:49:18", "2019-12-16 09:49:18");
INSERT INTO employees (`name`, `updated_at`, `created_at`) VALUES ("test3", "2019-12-16 09:49:18", "2019-12-16 09:49:18");
-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DELETE FROM employees WHERE id = 1;
DELETE FROM employees WHERE id = 2;
DELETE FROM employees WHERE id = 3;
