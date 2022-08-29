create database if not exists devbook;
use devbook;
drop table if exists users;
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