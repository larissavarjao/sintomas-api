package user

/* 1. BANCO DE DADOS - TABELA DE USUARIOS

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    first_name varchar(100) NOT NULL,
    last_name varchar(100),
		password varchar(255),
    email varchar(100) UNIQUE NOT NULL,
		avatar varchar(255)
);

*/

type User struct {
	ID    string     `json:"id"`
	FirstName  string    `json:"firstName"`
	LastName  string  `json:"lastName"`
	Password  string  `json:"password"`
	Email string `json:"email"`
	Avatar string `json:"avatar"`
}