CREATE TABLE users
(
    id            int    not null primary key,
    refresh_token text,
    exp           bigint not null,
    ip            varchar(32),
    email         varchar(64)
)
