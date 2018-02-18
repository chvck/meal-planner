CREATE TABLE ingredient (
  id        SERIAL PRIMARY KEY,
  recipe_id INT REFERENCES recipe (id)     NOT NULL,
  name      TEXT                           NOT NULL,
  measure   TEXT,
  quantity  DECIMAL                        NOT NULL,
  UNIQUE(recipe_id, name)
);