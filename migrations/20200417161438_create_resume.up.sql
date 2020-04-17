create table resume
(
    id          serial    not null
        constraint resume_pk
            primary key,
    created     timestamp default now(),
    updated     timestamp not null,
    name        varchar,
    money       money,
    description text      not null,
    active      boolean   not null,
    id_client   integer   not null
);

alter table resume
    owner to test;

create unique index resume_id_uindex
    on resume (id);

