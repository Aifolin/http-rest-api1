create table vacancies
(
    id          serial    not null
        constraint vacancies_pk
            primary key,
    created     timestamp default now(),
    updated     timestamp not null,
    name        varchar   not null,
    title       text,
    description varchar   not null,
    of_salary   money,
    to_salary   money,
    active      boolean   not null,
    id_client   integer   not null,
    active_data timestamp
);

alter table vacancies
    owner to test;

create unique index vacancies_id_uindex
    on vacancies (id);

