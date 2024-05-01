-- +migrate Up
CREATE TABLE events(
                       id uuid DEFAULT uuid_generate_v4() PRIMARY KEY ,
                       entity_type int(11) NOT NULL,
                       action_type int(11) NOT NULL,
                       agent_id unsigned int NOT NULL,
                       date TIMESTAMP NOT NULL,
                       is_published bool NOT NULL DEFAULT false,
                       entity VARCHAR(255) NOT NULL,
                       store_id unsigned int NOT NULL,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +migrate Down
DROP TABLE events;