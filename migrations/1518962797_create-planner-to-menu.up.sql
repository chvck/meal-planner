CREATE TABLE planner_to_menu (
  id          SERIAL PRIMARY KEY,
  planner_id  INT REFERENCES "planner" (id) NOT NULL,
  menu_id     INT REFERENCES "menu" (id) NOT NULL
)
