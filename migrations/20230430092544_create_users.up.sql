CREATE TABLE users (
    id serial not null primary key,
    email varchar(50) not null unique,
    encrypted_password varchar(50) not null
); 
  
