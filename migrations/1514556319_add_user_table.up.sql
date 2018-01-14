CREATE TABLE "user" (
  id         SERIAL PRIMARY KEY,
  username   TEXT UNIQUE NOT NULL,
  email      TEXT UNIQUE NOT NULL,
  password   TEXT        NOT NULL,
  created_at INT         NOT NULL,
  updated_at INT         NOT NULL,
  last_login INT
);
