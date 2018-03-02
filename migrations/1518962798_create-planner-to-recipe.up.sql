CREATE TABLE planner_to_recipe (
  id          SERIAL PRIMARY KEY,
  planner_id  INT REFERENCES "planner" (id) ON DELETE CASCADE NOT NULL,
  recipe_id   INT REFERENCES "recipe" (id) ON DELETE CASCADE NOT NULL
)
