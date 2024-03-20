create table jobs
(
    id      char(36)    not null
        primary key,
    payload mediumtext  not null,
    result  mediumtext  null,
    status  varchar(16) not null,
    type    varchar(16) not null,
    user_id int         not null
);

create table users
(
    id         int auto_increment
        primary key,
    first_name varchar(32)                not null,
    last_name  varchar(32)                not null,
    email      varchar(32)                not null,
    password   varchar(64)                not null,
    token      varchar(64)                null,
    balance    decimal(8, 2) default 0.00 not null,
    constraint users_email_uindex
        unique (email)
);
