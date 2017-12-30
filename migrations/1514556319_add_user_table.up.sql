CREATE TABLE "user" (
  id         SERIAL PRIMARY KEY,
  username   TEXT UNIQUE NOT NULL,
  email      TEXT UNIQUE NOT NULL,
  password   TEXT        NOT NULL,
  salt       TEXT        NOT NULL,
  algorithm  TEXT        NOT NULL,
  iterations SMALLINT    NOT NULL,
  created_at INT         NOT NULL,
  updated_at INT         NOT NULL,
  last_login INT
);
