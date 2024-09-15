CREATE TABLE users (
  user_id   text PRIMARY KEY,
  mail text UNIQUE NOT NULL,
  name text      NOT NULL,
  hashed_password text NOT NULL
);

CREATE TABLE eisa_files (
  user_id text PRIMARY KEY REFERENCES users(user_id),
  file_path text NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at timestamp DEFAULT NULL
);

CREATE TABLE color_codes (
  user_id text PRIMARY KEY REFERENCES users(user_id),
  color_code_1 text NOT NULL,
  color_code_2 text NOT NULL
);