package pacient

import "time"

/* 2. BANCO DE DADOS - TABELA DE PACIENTES

CREATE TABLE pacients (
    id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    first_name varchar(100) NOT NULL,
    last_name varchar(100),
		type INT,
		gender INT,
		date_of_birth TIMESTAMP,
		CONSTRAINT user_id
				FOREIGN KEY(id)
						REFERENCES users(id)
);

*/

type Pacient struct {
	ID    int64     `json:"id"`
	FirstName  string    `json:"firstName"`
	LastName  string  `json:"lastName"`
	UserID  string  `json:"userID"`
	Type PacientType `json:"pacientType"`
	Gender GenderType `json:"gender"`
	DateOfBirth time.Time `json:"dateOfBirth"`
}

type PacientType int

const (
	Human = iota + 1
	Dog
	Cat
)

func (p PacientType) String() string {
	switch p {
	case Human:
		return "Humano"
	case Dog:
		return "Cachorro"
	case Cat:
		return "Gato"
	}

	return "Outro"
}

type GenderType int

const (
	Female = iota + 1
	Male
	NotDeclare
)

func (g GenderType) String() string {
	switch g {
	case Female:
		return "Feminino"
	case Male:
		return "Masculino"
	case NotDeclare:
		return "Prefere não informar"
	}

	return "Não precisa declarar mesmo"
}