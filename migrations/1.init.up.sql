CREATE TABLE IF NOT EXISTS users
(
    id integer PRIMARY KEY,
    first_name text NOT NULL,
    last_name text NOT NULL,
    email text NOT NULL,
    birthdate timestamp NOT NULL
    want_notifications boolean NOT NULL DEFAULT true
);

