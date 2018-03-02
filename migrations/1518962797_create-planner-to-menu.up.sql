CREATE TABLE planner_to_menu (
  id          SERIAL PRIMARY KEY,
  planner_id  INT REFERENCES "planner" (id) ON DELETE CASCADE NOT NULL,
  menu_id     INT REFERENCES "menu" (id) ON DELETE CASCADE NOT NULL
)
