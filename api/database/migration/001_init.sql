-- +goose Up
create table if not exists Locations (
    id            serial primary key,
    NumberLocation        int not null,
    NameLocation          varchar not null,
    City          varchar not null,
    StateLocation         varchar not null,
    Country       varchar not null,
    Postcode      varchar not null,
    Latitude       varchar not null,
    Longitude     varchar not null,
    OffsetLocation        varchar not null,
    Description   varchar not null
);

-- +goose Down
drop table if exists Locations;
