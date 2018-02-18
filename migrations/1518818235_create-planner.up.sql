CREATE TYPE mealtime AS ENUM('breakfast', 'lunch', 'dinner', 'snack');

CREATE TABLE planner (
  "id"          SERIAL PRIMARY KEY,
  "user_id"     INT REFERENCES "user" ("id") NOT NULL,
  "when"        INT NOT NULL,
  "for"         mealtime NOT NULL,
  UNIQUE("when", "for", "user_id")
);
