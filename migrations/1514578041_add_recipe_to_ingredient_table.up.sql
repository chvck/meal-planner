CREATE TABLE recipe_to_ingredient (
  id            SERIAL PRIMARY KEY,
  recipe_id     INT REFERENCES recipe (id)     NOT NULL,
  ingredient_id INT REFERENCES ingredient (id) NOT NULL,
  measure_id    INT REFERENCES measure (id)    NOT NULL,
  quantity      SMALLINT                       NOT NULL
);
