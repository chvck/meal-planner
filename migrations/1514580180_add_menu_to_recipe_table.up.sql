CREATE TABLE menu_to_recipe (
  id        SERIAL PRIMARY KEY,
  menu_id   INT REFERENCES menu (id) ON DELETE CASCADE NOT NULL,
  recipe_id INT REFERENCES recipe (id) ON DELETE CASCADE NOT NULL
);
