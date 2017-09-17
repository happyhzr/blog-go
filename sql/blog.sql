DROP TABLE IF EXISTS users;

CREATE TABLE users(
    id serial4 primary key,
    name text,
    password text
);

DROP TABLE IF EXISTS posts;

CREATE TABLE posts(
    id serial4 primary key,
    title text,
    content text,
    created_at int8,
    created_by int4
);
