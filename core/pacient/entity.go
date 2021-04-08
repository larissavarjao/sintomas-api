package pacient

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

/* 2. BANCO DE DADOS - TABELA DE PACIENTES

CREATE TABLE pacients (
    id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    first_name varchar(100) NOT NULL,
    last_name varchar(100) NOT NULL,
    type INT NOT NULL,
    gender INT NOT NULL,
    date_of_birth TIMESTAMP NOT NULL,
    user_id uuid NOT NULL,
    CONSTRAINT user_id
        FOREIGN KEY(user_id)
            REFERENCES users(id)
);

INSERT INTO symptoms(name, description, type, happened_at, custom, duration, observation, pacient_id) values (
    'Caquexia',
    'Description',
    1,
    NOW(),
    false,
    10,
    'Observation',
    '71bcd242-be3b-43fb-b3a9-e5280037b0f0'
);

*/

type Pacient struct {
	ID    uuid.UUID     `json:"id"`
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