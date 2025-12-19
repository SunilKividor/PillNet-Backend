-- +goose Up
-- +goose StatementBegin
CREATE TABLE medicine_category(
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    parent_category UUID REFERENCES medicine_category(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE medicine_category;
-- +goose StatementEnd
