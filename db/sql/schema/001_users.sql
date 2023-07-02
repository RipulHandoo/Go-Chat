-- +goose Up
CREATE TABLE users(
    id bigserial,
    Email varchar(250) not null UNIQUE,
    primary key(id),
    password varchar(250) not null,
    username varchar not null
);

-- +goose Down
DROP TABLE users;