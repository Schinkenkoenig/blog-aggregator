-- +goose Up 
-- add name column  
--

ALTER TABLE feeds 
ADD COLUMN name TEXT NOT NULL;

-- +goose Down
-- delete feeds table
ALTER TABLE feeds 
DROP COLUMN name;

