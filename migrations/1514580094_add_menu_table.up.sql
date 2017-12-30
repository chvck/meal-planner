CREATE TABLE menu (
  id          SERIAL PRIMARY KEY,
  user_id     INT REFERENCES "user" (id) NOT NULL,
  name        TEXT                       NOT NULL,
  description TEXT
);
