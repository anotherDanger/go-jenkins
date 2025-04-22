create database if not exists perpustakaan_new;

use perpustakaan_new;

create table if not exists(
    id int auto_increment primary key ,
    author varchar(255),
    title varchar(255)
);