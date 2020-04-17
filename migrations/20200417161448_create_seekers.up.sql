create table seekers
(
    id_client integer not null,
    gender    varchar(10)
);

alter table seekers
    owner to test;

