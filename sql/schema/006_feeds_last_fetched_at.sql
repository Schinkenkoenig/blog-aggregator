
-- +goose Up 
-- add last fetched at 
ALTER TABLE feeds 
ADD COLUMN last_fetched_at TIMESTAMP;

-- +goose Down
-- drop last fetched at
ALTER TABLE feeds
DROP COLUMN last_fetched_at;
