CREATE TABLE users(
    id bigserial not null primary key,
    email varchar not null enique,
    encrypted_password varchar not null
);