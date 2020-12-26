create extension if not exists pg_trgm;

create table if not exists verified_messages
(
    id              character varying(36)    not null,
    checked         boolean default false,
    explanation     text                     null,
    first_appear    timestamp with time zone not null,
    is_fake         boolean,
    link            character varying(255)   null,
    text            text                     not null,
    text_normalized text                     not null
);