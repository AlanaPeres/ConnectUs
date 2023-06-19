CREATE DATABASE IF NOT EXISTS connectUs;
USE connectUs;
DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios(
    id int auto_increment primary key,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(10) not null unique,
    CriadoEm timestamp default current_timestamp()
)ENGINE=INNODB;  /*é usada para especificar o mecanismo de armazenamento (storage engine) que será usado para essa tabela. O InnoDB é um mecanismo de armazenamento transacional que oferece recursos como ACID (Atomicidade, Consistência, Isolamento e Durabilidade) 
e suporte a chave estrangeira (foreign key).*/