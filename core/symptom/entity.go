package symptom

import "time"

/* 3. BANCO DE DADOS - TABELA DE TIPOS DE SINTOMAS

CREATE TABLE symptoms (
    id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    name varchar(255) NOT NULL,
		description varchar(255),
		type INT NOT NULL,
		happened_at timestamp NOT NULL,
		default boolean,
		duration int,
		observation varchar(255),
		CONSTRAINT fk_customer
				FOREIGN KEY(id)
						REFERENCES pacients(id)
);

*/

type Symptom struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Type SymptomType `json:"type"`
	Default bool `json:"default"`
	HappenedAt time.Time `json:"happenedAt"`
	Duration time.Duration `json:"duration"`
	Observation string `json:"observation"`
	PacientID string `json:"pacientID"`
}

type SymptomType int

const (
	GenericType = 1
  NeurologicalType = 2
  OphthalmologicalType = 3
  GastrointestinalType = 4
  CardiovascularType = 5
  UrologicalType = 6
  PulmonaryType = 7
  IntegumentaryType = 8
	Other = 9
)

func (s SymptomType) String() string {
	switch s {
	case GenericType:
		return "Genérico"
	case NeurologicalType:
		return "Neurológico"
	case OphthalmologicalType:
		return "Oftalmológico"
	case GastrointestinalType:
		return "Gastrointestinal"
	case CardiovascularType:
		return "Cardiovascular"
	case UrologicalType:
		return "Urológico"
	case PulmonaryType:
		return "Pulmonar"
	case IntegumentaryType:
		return "Inter"
	case Other:
		return "Outro"
	}

	return "Outro"
}