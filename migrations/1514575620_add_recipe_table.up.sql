CREATE TABLE recipe (
  id           SERIAL PRIMARY KEY,
  user_id      INT REFERENCES "user" (id) NOT NULL,
  name         TEXT NOT NULL,
  instructions TEXT NOT NULL,
  yield        SMALLINT NOT NULL,
  prep_time    SMALLINT NOT NULL,
  cook_time    SMALLINT NOT NULL,
  description  TEXT
);
