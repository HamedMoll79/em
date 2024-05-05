-- +migrate Up
CREATE TABLE events(
                       id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
                       entity_type INT NOT NULL,
                       action_type INT NOT NULL,
                       agent_id INT NOT NULL,
                       date TIMESTAMP NOT NULL,
                       is_published bool NOT NULL DEFAULT FALSE,
                       entity VARCHAR(255) NOT NULL,
                       store_id INT NOT NULL,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +migrate Down
DROP TABLE events;