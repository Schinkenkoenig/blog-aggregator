-- +goose Up 
-- add feeds tables 
CREATE TABLE feeds (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL, 
  url TEXT NOT NULL UNIQUE,
  user_id UUID REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
-- delete feeds table
DROP TABLE feeds;
