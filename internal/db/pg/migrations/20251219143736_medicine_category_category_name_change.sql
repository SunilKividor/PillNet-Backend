-- +goose Up
-- +goose StatementBegin
-- FLAGGED AS REDUNDANT: ALTER TABLE medicines RENAME COLUMN category TO category_id;
SELECT 1;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 1;
-- +goose StatementEnd
