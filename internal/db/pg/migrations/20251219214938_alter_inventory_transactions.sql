-- +goose Up
-- +goose StatementBegin
ALTER TABLE inventory_transactions
ADD COLUMN inventory_id UUID NOT NULL REFERENCES inventory_stock(id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
