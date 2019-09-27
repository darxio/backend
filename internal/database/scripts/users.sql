SET SYNCHRONOUS_COMMIT = 'off';
CREATE EXTENSION IF NOT EXISTS CITEXT;

CREATE TABLE users (
  id       BIGSERIAL      PRIMARY KEY,
  username CITEXT         NOT NULL UNIQUE,
  password TEXT           NOT NULL
);

CREATE TABLE sessions (
    user_id     INT REFERENCES users(id),
    cookie      TEXT DEFAULT ''

    CONSTRAINT sessions_pkey PRIMARY KEY (user_id, cookie)
);