-- +goose Up 
-- add feeds follow tabet 
CREATE TABLE feeds_users (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL, 
  user_id UUID NOT NULL,
  feed_id UUID NOT NULL,
  UNIQUE (user_id, feed_id)
);

-- +goose Down
-- delete feeds table
DROP TABLE feeds_users;
