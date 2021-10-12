SET timezone = 'Europe/Moscow';
-- SHOW TIMEZONE;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users
(
    id serial not null unique,
    username varchar(50) not null,
    first_name varchar(300) not null,
    last_name varchar(300) not null,
    email varchar(300) not null,
    password varchar(200) not null,
    phone varchar(50) not null,
    address text null
);
-- DROP TABLE IF EXISTS users CASCADE;

CREATE TABLE IF NOT EXISTS loyalty_card
(
    id serial not null unique,
    user_id int references users (id) on delete cascade on update cascade not null,
    code uuid NOT NULL DEFAULT uuid_generate_v4(),
    create_date timestamp without time zone default now(),
    status smallint NOT NULL DEFAULT 0
);
-- DROP TABLE IF EXISTS loyalty_card CASCADE;

CREATE TABLE IF NOT EXISTS notification
(
    id serial not null unique,
    user_id int references users (id) on delete cascade on update cascade not null,
    send_date timestamp without time zone default now(),
    notification_program_type smallint NOT NULL DEFAULT 0,
    domain_event_type smallint NOT NULL DEFAULT 0,
    status smallint NOT NULL DEFAULT 0
);
-- DROP TABLE IF EXISTS notification CASCADE;
