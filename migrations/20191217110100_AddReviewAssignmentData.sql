-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

INSERT INTO review_assignments (`reviewee`, `reviewer`, `updated_at`, `created_at`) VALUES (1, 2, "2019-12-16 09:49:18", "2019-12-16 09:49:18");
INSERT INTO review_assignments (`reviewee`, `reviewer`, `updated_at`, `created_at`) VALUES (1, 3, "2019-12-16 09:49:18", "2019-12-16 09:49:18");
INSERT INTO review_assignments (`reviewee`, `reviewer`, `updated_at`, `created_at`) VALUES (3, 1, "2019-12-16 09:49:18", "2019-12-16 09:49:18");
-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DELETE FROM review_assignments WHERE id = 1;
DELETE FROM review_assignments WHERE id = 2;
DELETE FROM review_assignments WHERE id = 3;
