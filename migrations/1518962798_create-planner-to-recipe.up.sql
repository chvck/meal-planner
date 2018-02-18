CREATE TABLE planner_to_recipe (
  id          SERIAL PRIMARY KEY,
  planner_id  INT REFERENCES "planner" (id) NOT NULL,
  recipe_id   INT REFERENCES "recipe" (id) NOT NULL
)
