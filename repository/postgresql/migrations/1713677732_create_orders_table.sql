-- +migrate Up
CREATE TYPE action_type AS ENUM ('create', 'update', 'delete')
CREATE TABLE orders(
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY ,
    store_id VARCHAR(255) NOT NULL,
    web_hook_url VARCHAR(255 NOT NULL,
    meta_data jsonb,
    amount bigint NOT NULL,
    user_id varchar(255) not null,
    type action_type NOT NULL,
    is_published BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
);

-- +migrate Down
DROP TABLE orders;
