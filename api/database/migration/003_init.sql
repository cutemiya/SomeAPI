-- +goose Up
create table if not exists Pictures (
    id            serial primary key,
    Large         varchar not null,
    Medium        varchar not null,
    Thumbnail     varchar not null
);

-- +goose Down
drop table if exists Pictures;