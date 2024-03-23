
-- +goose Up 
-- add column apiKey 
alter table users 
add column apiKey varchar(64) not null unique
default encode(sha256(random()::text::bytea), 'hex');

-- +goose Down
-- delete apiKey column
alter table users drop column apiKey;
