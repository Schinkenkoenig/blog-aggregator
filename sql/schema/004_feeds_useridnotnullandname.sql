-- +goose Up 
-- add name column and make user id not nulladd feeds tables 
--

ALTER TABLE feeds 
ALTER COLUMN user_id SET NOT NULL,
ADD COLUMN name TEXT NOT NULL;

-- +goose Down
-- delete feeds table
ALTER TABLE feeds 
DROP COLUMN name
ALTER COLUMN user_id SET NULL;

