create table respond
(
    id         serial    not null
        constraint respond_pk
            primary key,
    created    timestamp not null,
    updated    timestamp,
    status     integer,
    id_vacancy integer   not null,
    id_resume  integer   not null
);

alter table respond
    owner to test;

create unique index respond_id_uindex
    on respond (id);