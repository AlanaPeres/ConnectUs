CREATE DATABASE IF NOT EXISTS connectUs; 
USE connectUs;
DROP TABLE IF EXISTS usuarios; 

CREATE TABLE usuarios(
    id int auto_increment primary key,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(10) not null,
    CriadoEm timestamp default current_timestamp()
)ENGINE=INNODB;  