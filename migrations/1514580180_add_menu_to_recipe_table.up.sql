CREATE TABLE menu_to_recipe (
  id        SERIAL PRIMARY KEY,
  menu_id   INT REFERENCES menu (id)   NOT NULL,
  recipe_id INT REFERENCES recipe (id) NOT NULL
);
