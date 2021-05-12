CREATE TABLE IF NOT EXISTS "user"(
    id serial PRIMARY KEY,
    username varchar unique not null,
    email varchar unique not null,
    password varchar not null,
    fname varchar(30) not null,
    surname varchar(50) not null,
    created_at      timestamp with time zone default now() not null,
    updated_at      timestamp with time zone    
);