-- migrate:up
CREATE TABLE prices (
    code VARCHAR(10) NOT NULL,
    price DOUBLE NOT NULL,
    currency VARCHAR(3) NOT NULL,
    timestamp TIMESTAMP NOT NULL
);

-- migrate:down

