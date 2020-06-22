CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR(25) UNIQUE NOT NULL,
    password VARCHAR(64) NOT NULL,
    access_token VARCHAR(36) NOT NULL
);