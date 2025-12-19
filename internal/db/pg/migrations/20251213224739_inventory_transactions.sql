-- +goose Up
-- +goose StatementBegin
CREATE TABLE inventory_transactions (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    medicine_id UUID NOT NULL REFERENCES medicines(id),
    batch_number VARCHAR(100),

    transaction_type VARCHAR(30) NOT NULL CHECK (
        transaction_type IN (
            'purchase','sale','return','adjustment',
            'transfer','waste','expiry','damage','theft'
        )
    ),

    from_location_id UUID REFERENCES storage_locations(id),
    to_location_id UUID REFERENCES storage_locations(id),

    unit_price DECIMAL(10,2),
    quantity INTEGER,
    total_value DECIMAL(12,2),
    
    performed_by UUID REFERENCES users(id),

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE inventory_transactions;
-- +goose StatementEnd
