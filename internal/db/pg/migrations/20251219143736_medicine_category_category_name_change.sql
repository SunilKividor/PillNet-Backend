-- +goose Up
-- +goose StatementBegin
ALTER TABLE medicines
RENAME COLUMN category TO category_id;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE medicines
RENAME COLUMN category_id TO category;
-- +goose StatementEnd
