CREATE TABLE ingredient (
  id        SERIAL PRIMARY KEY,
  recipe_id INT REFERENCES recipe (id)     NOT NULL,
  name      TEXT                           NOT NULL,
  measure   TEXT,
  quantity  SMALLINT                       NOT NULL
);

ALTER TABLE ingredient ADD CONSTRAINT ingredient_unique_name_recipe UNIQUE (recipe_id, name);