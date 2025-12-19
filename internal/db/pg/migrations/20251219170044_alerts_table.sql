-- +goose Up
-- +goose StatementBegin
CREATE TABLE alerts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    alert_type TEXT,            -- LOW_STOCK, NEAR_EXPIRY
    medicine_id UUID,
    location_id UUID,

    message TEXT,

    status TEXT DEFAULT 'OPEN', -- OPEN, ACK, RESOLVED

    triggered_at TIMESTAMP DEFAULT now(),
    resolved_at TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE alerts;
-- +goose StatementEnd
