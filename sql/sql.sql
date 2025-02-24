CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS posts;
DROP TABLE IF EXISTS followers;
DROP TABLE IF EXISTS users;

CREATE TABLE users(
    id int auto_increment primary key,
    name varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique, 
    pwd varchar(100) not null,
    createdAt timestamp default current_timestamp()
) ENGINE=INNODB;

CREATE TABLE followers(
    user_id int not null foreign key (user_id) references users(id) on delete cascade,
    follower_id int not null foreign key (follower_id) references users(id) on delete cascade,
    createdAt timestamp default current_timestamp()
    primary key(user_id, follower_id)
)ENGINE=INNODB;

CREATE TABLE posts(
    id int auto_increment primary key,
    title varchar(50) not null,
    content varchar(300) not null,
    author_id int not null,
    foreign key (author_id)
    references users(id)
    on delete cascade,

    likes int default 0,
    createdAt timestamp default current_timestamp()
)ENGINE=INNODB;