create table chat
(
    id_message serial                  not null
        constraint chat_pk
            primary key,
    created    timestamp default now() not null,
    updated    timestamp default now() not null,
    message    text                    not null,
    id_respond integer                 not null,
    id_client  integer                 not null
);

alter table chat
    owner to test;

create unique index chat_id_message_uindex
    on chat (id_message);

