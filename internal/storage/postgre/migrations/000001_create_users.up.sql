create table users (
    id varchar primary key,
    name varchar(255) not null,
    surname varchar(255) not null,
    login varchar(255) not null unique,
    password varchar(255) not null
)