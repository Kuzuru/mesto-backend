CREATE TABLE users
(
    id      SERIAL PRIMARY KEY NOT NULL UNIQUE,
    auth_id UUID               NOT NULL DEFAULT gen_random_uuid(),
    name    VARCHAR(100)       NOT NULL,
    about   VARCHAR(100)       NOT NULL,
    avatar  VARCHAR(100)       NOT NULL
);

INSERT INTO users (name, about, avatar)
VALUES ('Mark Fialko', 'JavaScript Frontend Developer', 'https://avatars.githubusercontent.com/u/49472529?v=4');

SELECT *
FROM users;
