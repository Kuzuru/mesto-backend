CREATE TABLE users
(
    id      SERIAL PRIMARY KEY NOT NULL UNIQUE,
    auth_id UUID               NOT NULL DEFAULT gen_random_uuid(),
    name    VARCHAR(127)       NOT NULL,
    about   VARCHAR(127)       NOT NULL,
    avatar  VARCHAR(255)       NOT NULL
);

SELECT * FROM users;