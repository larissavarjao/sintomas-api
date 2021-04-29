package pacient

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type UseCase interface {
	GetAll() ([]*Pacient, error)
}

type Service struct {
	DB *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{
		DB: db,
	}
}

func (s *Service) GetAll() ([]*Pacient, error) {
	var result []*Pacient

	rows, err := s.DB.Query("select id, first_name, last_name, type, gender, date_of_birth, user_id from pacients")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p Pacient

		err := rows.Scan(&p.ID, &p.FirstName, &p.LastName, &p.Type, &p.Gender, &p.DateOfBirth, &p.UserID)
		if err != nil {
			return nil, err
		}
		
		result = append(result, &p)
	}

	return result, nil
}