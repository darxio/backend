CREATE EXTENSION IF NOT EXISTS CITEXT;
DROP TABLE users CASCADE;
DROP TABLE sessions CASCADE;
DROP TABLE groups CASCADE;
DROP TABLE user_groups CASCADE;
DROP TABLE ingredient_types CASCADE;
DROP TABLE group_ingredient_types CASCADE;
DROP TABLE excluded_ingredients CASCADE;
DROP TABLE products CASCADE;
DROP TABLE product_ingredients CASCADE;

-------------------------------------------
-- users & sessions
-------------------------------------------

CREATE TABLE users (
  id       BIGSERIAL      PRIMARY KEY,
  username CITEXT         NOT NULL UNIQUE,
  password TEXT           NOT NULL
);

CREATE TABLE sessions (
    user_id     INT REFERENCES users(id),
    cookie      TEXT DEFAULT '',

    CONSTRAINT sessions_pkey PRIMARY KEY (user_id, cookie)
);

-------------------------------------------
-- groups
-------------------------------------------

CREATE TABLE groups (
  id       BIGSERIAL      PRIMARY KEY,
  name     CITEXT         NOT NULL UNIQUE,
  about    TEXT           NOT NULL DEFAULT ''
);

CREATE TABLE user_groups (
  user_id       BIGINT    REFERENCES users(id),
  group_id      BIGINT    REFERENCES groups(id),

  CONSTRAINT usergroups_pkey PRIMARY KEY (user_id, group_id)
);

-------------------------------------------
-- ingredients
-------------------------------------------

CREATE TABLE ingredient_types (
  id       BIGSERIAL      PRIMARY KEY,
  type     TEXT           NOT NULL DEFAULT '' UNIQUE
);

CREATE TABLE ingredients (
  id       BIGSERIAL      PRIMARY KEY,
  name     CITEXT         NOT NULL UNIQUE,
  about    TEXT           NOT NULL DEFAULT '',
  type_id  BIGINT         REFERENCES ingredient_types(id)
);

CREATE TABLE group_ingredient_types (
  group_id      BIGINT    REFERENCES groups(id),
  type_id       BIGINT    REFERENCES ingredient_types(id),

  CONSTRAINT grouptypes_pkey PRIMARY KEY (group_id, type_id)
);

CREATE TABLE excluded_ingredients (
  ingredient_id      BIGINT    REFERENCES ingredients(id),
  user_id            BIGINT    REFERENCES users(id),

  CONSTRAINT exclingredients_pkey PRIMARY KEY (ingredient_id, user_id)
);

-------------------------------------------
-- products
-------------------------------------------

CREATE TABLE products (
  id       BIGSERIAL      PRIMARY KEY,
  name     CITEXT         NOT NULL,
  barcode  TEXT            NOT NULL 
  -- nutrition facts
);

CREATE TABLE product_ingredients (
  product_id      BIGINT    REFERENCES products(id),
  ingredient_id   BIGINT    REFERENCES ingredients(id),

  CONSTRAINT productingredients_pkey PRIMARY KEY (product_id, ingredient_id)
);

INSERT INTO products(name, barcode) VALUES ('snickers', 11111);