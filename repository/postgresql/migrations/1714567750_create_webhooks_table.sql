-- +migrate Up
CREATE TABLE webhooks(
                       id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
                       url VARCHAR(255) NOT NULL,
                       store_id unsigned int NOT NULL,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +migrate Down
DROP TABLE webhooks;