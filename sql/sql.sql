create database if not exists devbook;
use devbook;
drop table if exists users;
drop table if exists follows;

create table users (
    id int not null auto_increment primary key,
    name varchar(255) not null,
    nick varchar(255) not null unique,
    email varchar(255) not null unique,
    password varchar(255) not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,
    deleted_at  DATETIME  NULL
) ENGINE=INNODB;

CREATE TABLE IF NOT EXISTS follows (
    user_id int not null,
    foreign key (user_id) 
    references users(id) 
    on delete cascade,
    following_id int not null,
    foreign key (following_id)
    references users(id),
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,
    deleted_at  DATETIME  NULL,
    primary key (user_id, following_id)
) ENGINE=INNODB;