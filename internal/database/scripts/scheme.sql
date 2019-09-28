CREATE EXTENSION IF NOT EXISTS CITEXT;

-------------------------------------------
-- users & sessions
-------------------------------------------

CREATE TABLE users (
  id       BIGSERIAL      PRIMARY KEY,
  username CITEXT         NOT NULL UNIQUE,
  password TEXT           NOT NULL
);

CREATE TABLE sessions (
    user_id     INT REFERENCES users(id),
    cookie      TEXT DEFAULT '',

    CONSTRAINT sessions_pkey PRIMARY KEY (user_id, cookie)
);

-------------------------------------------
-- groups
-------------------------------------------

CREATE TABLE groups (
  id       BIGSERIAL      PRIMARY KEY,
  name     CITEXT         NOT NULL UNIQUE,
  about    TEXT           NOT NULL DEFAULT ''
);

CREATE TABLE user_groups (
  user_id       BIGINT    REFERENCES users(id),
  group_id      BIGINT    REFERENCES groups(id),

  CONSTRAINT usergroups_pkey PRIMARY KEY (user_id, group_id)
);

-------------------------------------------
-- ingredients
-------------------------------------------

CREATE TABLE ingredient_types (
  id       BIGSERIAL      PRIMARY KEY,
  type     TEXT           NOT NULL DEFAULT '' UNIQUE
);

CREATE TABLE ingredients (
  id       BIGSERIAL      PRIMARY KEY,
  name     CITEXT         NOT NULL UNIQUE,
  about    TEXT           NOT NULL DEFAULT '',
  type_id  BIGINT         REFERENCES ingredient_types(id)
);

CREATE TABLE group_ingridient_types (
  group_id      BIGINT    REFERENCES groups(id),
  type_id       BIGINT    REFERENCES ingredient_types(id),

  CONSTRAINT grouptypes_pkey PRIMARY KEY (group_id, type_id)
);

CREATE TABLE excluded_ingredients (
  ingredient_id      BIGINT    REFERENCES ingredients(id),
  user_id            BIGINT    REFERENCES users(id),

  CONSTRAINT exclingredients_pkey PRIMARY KEY (ingredient_id, user_id)
);
