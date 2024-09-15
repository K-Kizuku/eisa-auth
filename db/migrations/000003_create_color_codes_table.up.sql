CREATE TABLE IF NOT EXISTS color_codes (
  user_id text PRIMARY KEY REFERENCES users(user_id),
  color_code_1 text NOT NULL,
  color_code_2 text NOT NULL
);