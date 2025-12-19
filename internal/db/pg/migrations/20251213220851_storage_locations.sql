-- +goose Up
-- +goose StatementBegin
CREATE TABLE storage_locations(
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    location_type VARCHAR(50) NOT NULL,

    temperature_controlled BOOLEAN DEFAULT false,
    temperature_min DECIMAL(5,2),
    temperature_max DECIMAL(5,2),
    humidity_controlled BOOLEAN DEFAULT false,
    humidity_min DECIMAL(5,2),
    humidity_max DECIMAL(5,2),

    requires_key_access BOOLEAN DEFAULT false,
    
    capacity DECIMAL(10,2),
    current_utilization_percentage DECIMAL(10,2),

    floor_number INTEGER,
    section VARCHAR(50),

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE storage_locations;
-- +goose StatementEnd
