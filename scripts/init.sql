CREATE DATABASE IF NOT EXISTS devbook;

USE devbook;

CREATE TABLE usuarios (
  id int auto_increment primary key,
  nome varchar(50) not null,
  email varchar(50) not null
) ENGINE=InnoDB;

INSERT INTO usuarios (nome, email) VALUES ('Charles', 'swirfneblin@gmail');