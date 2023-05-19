BEGIN;

CREATE TABLE users (
    id UUID PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    user_name VARCHAR(255) NOT NULL UNIQUE,
    user_role VARCHAR(255) NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    date_added TIMESTAMP NOT NULL,
    date_updated TIMESTAMP NOT NULL,
    concurrency_stamp VARCHAR(255) NOT NULL
);

CREATE TABLE password_reset_tokens (
    id UUID PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    expiration_date TIMESTAMP NOT NULL,
    date_added TIMESTAMP NOT NULL,
    CONSTRAINT fk_user
        FOREIGN KEY(email)
            REFERENCES users(email)
            ON DELETE CASCADE
);

COMMIT;