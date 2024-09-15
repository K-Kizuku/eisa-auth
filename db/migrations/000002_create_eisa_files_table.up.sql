CREATE TABLE IF NOT EXISTS eisa_files (
  user_id text PRIMARY KEY REFERENCES users(user_id),
  file_path text NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at timestamp DEFAULT NULL
);
