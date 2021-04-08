package symptom

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type UseCase interface {
	GetAll() ([]*Symptom, error)
	Create(symptom *Symptom) (*Symptom, error)
}

type Service struct {
	DB *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{
		DB: db,
	}
}

func (s *Service) GetAll() ([]*Symptom, error) {
	var result []*Symptom

	rows, err := s.DB.Query("select * from symptoms")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var sy Symptom
		err := rows.Scan(&sy.ID, &sy.Name, &sy.Description, &sy.Type, &sy.HappenedAt, &sy.Custom, &sy.Duration, &sy.Observation, &sy.PacientID)
		if err != nil {
			return nil, err
		}
		result = append(result, &sy)
	}

	return result, nil
}

func (s *Service) Create(sy *Symptom) (*Symptom, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		return nil, err
	}

	smtx, err := s.DB.Prepare("insert into symptoms(name, description, type, custom, happened_at, duration, observation, pacient_id) values (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	defer smtx.Close()

	_, err = smtx.Exec(&sy.Name, &sy.Description, &sy.Type, &sy.HappenedAt, &sy.Custom, &sy.Duration, &sy.Observation, &sy.PacientID)
	if err != nil {
		return nil, err
	}

	tx.Commit()
	return nil, nil
}