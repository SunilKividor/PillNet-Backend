-- +goose Up
-- +goose StatementBegin
CREATE TABLE medicines(
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    generic_name VARCHAR(255) NOT NULL,
    trade_name VARCHAR(255) NOT NULL,
    category_id UUID REFERENCES medicine_category(id),

    dosage_form VARCHAR NOT NULL,
    strength VARCHAR(100),
    route_of_administration VARCHAR(50),

    is_prescription_required BOOLEAN DEFAULT true,
    is_controlled_substance BOOLEAN DEFAULT false,
    schedule_classification VARCHAR(10),
    
    storage_condition VARCHAR(50),
    storage_temperature_min DECIMAL(5,2),
    storage_temperature_max DECIMAL(5,2),
    requires_refrigeration BOOLEAN DEFAULT false,
    light_sensitive BOOLEAN DEFAULT false,
    moisture_sensitive BOOLEAN DEFAULT false,

    abc_classification CHAR(1) CHECK(abc_classification IN ('A','B','C')),
    ved_classification VARCHAR(10) CHECK(ved_classification IN ('Vital','Essential','Desirable')),
    fsn_classification VARCHAR(15) CHECK(fsn_classification IN ('Fast','Slow','Non-moving')),

    therapeutic_class VARCHAR(100),
    pharmacological_class TEXT,
    indications TEXT,
    contraindication TEXT,
    side_effects TEXT,

    unit_price DECIMAL(10,2),
    mrp DECIMAL(10,2),

    is_active BOOLEAN DEFAULT true,
    discontinuation_reason TEXT,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE medicines;
-- +goose StatementEnd
