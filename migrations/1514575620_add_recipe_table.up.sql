CREATE TABLE recipe (
  id           SERIAL PRIMARY KEY,
  user_id      INT REFERENCES "user" (id) NOT NULL,
  name         TEXT NOT NULL,
  instructions TEXT NOT NULL,
  yield        SMALLINT,
  prep_time    SMALLINT,
  cook_time    SMALLINT,
  description  TEXT
);
