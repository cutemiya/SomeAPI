-- +goose Up
create table if not exists Users (
    id                  serial primary key,
    Gender              varchar not null,
    Title               varchar not null,
    First               varchar not null,
    Last                varchar not null,
    Email               varchar not null,
    Date                varchar not null,
    Age                 int not null,
    RegisteredDate      varchar not null,
    RegisteredAge       int not null,
    Phone               varchar not null,
    Cell                varchar not null,
    IdName              varchar not null,
    IdValue             varchar not null,
    Nat                 varchar not null,
    LocationId          int not null,
    LoginId             int not null,
    PictureId           int not null,

    foreign key (LocationId) references Locations(id),
    foreign key (LoginId) references Logins(id),
    foreign key (PictureId) references Pictures(id)
);

-- +goose Down
drop table if exists Users;