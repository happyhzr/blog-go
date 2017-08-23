CREATE DATABASE IF NOT EXISTS blog;
USE blog;

DROP TABLE IF EXISTS user;
CREATE TABLE user(
	id bigint NOT NULL auto_increment,
    username varchar(30),
    password varchar(30),
    PRIMARY KEY (id)
);

DROP TABLE IF EXISTS post;
CREATE TABLE post(
	id bigint NOT NULL auto_increment,
    title varchar(30),
    content text,
    created_at bigint,
    updated_at bigint,
    PRIMARY KEY(id)
);