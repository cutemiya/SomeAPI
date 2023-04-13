-- +goose Up
create table if not exists Logins (
    id            serial primary key,
    Uuid          varchar not null,
    Username      varchar not null,
    Password      varchar not null,
    Salt          varchar not null,
    Md5           varchar not null,
    Sha1          varchar not null,
    Sha256        varchar not null
);

-- +goose Down
drop table if exists Logins;