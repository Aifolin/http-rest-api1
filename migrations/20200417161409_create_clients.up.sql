create table clients
(
    id           serial                  not null
        constraint client_pk
            primary key,
    created      timestamp default now() not null,
    updated      timestamp default now() not null,
    name         varchar                 not null,
    email        varchar,
    number_phone varchar                 not null,
    active       boolean                 not null
);

alter table clients
    owner to test;

create unique index client_id_uindex
    on clients (id);

