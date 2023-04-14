CREATE TABLE users (
    id          int primary key AUTO_INCREMENT,
    name        varchar(255) not null default "h",
    phone_number varchar(255) not null unique ,
    created_at timestamp default current_timestamp
);