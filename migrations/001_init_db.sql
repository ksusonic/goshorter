create table shortened
(
    hash_id     varchar(10) primary key,
    url         text                     not null,
    create_time timestamp with time zone not null default now()
)