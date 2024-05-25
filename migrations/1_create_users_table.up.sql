create table if not exists users
(
    id            integer not null
        primary key,
    name          varchar(255),
    user_name     varchar(255),
    language_code varchar(10),
    is_premium    boolean,
    is_bot        boolean,
    active        boolean,
    date_add      integer
);

alter table users
    owner to skin_user;

