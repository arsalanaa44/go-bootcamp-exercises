CREATE TABLE users (
    id          int primary key AUTO_INCREMENT,
    name        varchar(255) not null default "h",
    phone_number varchar(255) not null unique ,
    password text not null,
    created_at timestamp default current_timestamp
);

// for migration
    alter table users add column password text not null;
