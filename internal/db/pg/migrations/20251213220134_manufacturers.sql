-- +goose Up
-- +goose StatementBegin
CREATE TABLE manufacturers (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY, 
    name VARCHAR(255) NOT NULL,
    license_number VARCHAR(255),
    contact_person VARCHAR(100),
    email VARCHAR(255),
    phone VARCHAR(20),
    address TEXT,
    country VARCHAR(100),

    reliability_score DECIMAL(3,2) DEFAULT 0.5,
    avg_delivery_time_days INTEGER,
    quality_rating DECIMAL(3,2),

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE manufacturers;
-- +goose StatementEnd
