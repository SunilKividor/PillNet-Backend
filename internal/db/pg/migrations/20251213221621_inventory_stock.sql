-- +goose Up
-- +goose StatementBegin
CREATE TABLE inventory_stock(
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    medicine_id UUID NOT NULL REFERENCES medicines(id),
    batch_number VARCHAR(255) UNIQUE NOT NULL,

    quantity INTEGER NOT NULL CHECK(quantity >= 0),
    received_quantity INTEGER NOT NULL,
    reserved_quantity INTEGER DEFAULT 0,
    damaged_quantity INTEGER DEFAULT 0,

    manufacturer_date DATE NOT NULL,
    expiry_date DATE NOT NULL,
    received_date DATE NOT NULL,

    unit_cost_price DECIMAL(10,2) NOT NULL,
    unit_selling_price DECIMAL(10,2) NOT NULL,
    total_value DECIMAL(12,2),

    location_id UUID REFERENCES storage_locations(id),
    panel_code VARCHAR(10),
    row_number INTEGER,
    rack_code VARCHAR(10),
    bin_number INTEGER,

    supplier_id UUID REFERENCES manufacturers(id),
    status VARCHAR(20) DEFAULT 'active',

    stock_checked_by UUID REFERENCES users(id),
    stock_checked_at TIMESTAMP,
    
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE inventory_stock;
-- +goose StatementEnd
