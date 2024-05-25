create table if not exists stickers
(
    instanceid            integer not null primary key,
    name                  varchar(255),
    hash_name             varchar(255),
    sell_price            bigint,
    sell_price_text       varchar(255),
    is_custom_sell_price  boolean
);

alter table stickers
    owner to skin_user;