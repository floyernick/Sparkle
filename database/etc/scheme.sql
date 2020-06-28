CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR(25) UNIQUE NOT NULL,
    password VARCHAR(64) NOT NULL,
    access_token VARCHAR(36) NOT NULL
);
CREATE TABLE posts (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    text VARCHAR(150) NOT NULL,
    location_code VARCHAR(11) NOT NULL,
    created_at TIMESTAMP NOT NULL
);