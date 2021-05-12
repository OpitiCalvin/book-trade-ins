CREATE TABLE IF NOT EXISTS book(
    id serial PRIMARY KEY,
    title varchar not null,
    author varchar not null,
    created_at      timestamp with time zone default now() not null,
    updated_at      timestamp with time zone,
    user_id integer not null,
    FOREIGN KEY (user_id) REFERENCES "user" (id)
);