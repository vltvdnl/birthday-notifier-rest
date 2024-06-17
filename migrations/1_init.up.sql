CREATE TABLE IF NOT EXISTS users
(
    id integer PRIMARY KEY,
    first_name text NOT NULL,
    last_name text NOT NULL,
    email text NOT NULL,
    birthdate timestamp NOT NULL,
    want_notifications boolean NOT NULL DEFAULT true
);

CREATE TABLE IF NOT EXISTS users_subsciptions
(  
    follower_id integer PRIMARY KEY REFERENCES users(id),
    user_id integer REFERENCES users(id)
);

CREATE INDEX IF NOT EXISTS idx_userid ON users_subsciptions(user_id);