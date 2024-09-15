CREATE TABLE IF NOT EXISTS  users (
  user_id   text PRIMARY KEY,
  mail text UNIQUE NOT NULL,
  name text      NOT NULL,
  hashed_password text NOT NULL
);
