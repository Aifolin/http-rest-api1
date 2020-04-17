create table employers
(
    id_client integer not null,
    about     text
);

alter table employers
    owner to test;