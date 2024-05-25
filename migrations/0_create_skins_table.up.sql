create table if not exists skins
(
    id           varchar(255) not null
        primary key,
    name         varchar(255),
    image        varchar(510),
    inspect_link varchar(510),
    float        varchar(255),
    new          boolean,
    price        varchar(255),
    tradable     boolean,
    url          varchar(255)
);

alter table skins
    owner to skin_user;

