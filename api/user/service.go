package user

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type UseCase interface {
	GetAll() ([]*User, error)
	Register(*User) (*User, error)
	Login(email *string, password *string) (*User, error)
	Me() (*User, error)
}

type Service struct {
	DB *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{
		DB: db,
	}
}

func (s *Service) GetAll() ([]*User, error) {
	var result []*User

	rows, err := s.DB.Query("select * from users")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Password, &u.Email, &u.Avatar)
		if err != nil {
			return nil, err
		}

		result = append(result, &u)
	}

	return result, nil
}

func (s *Service) Register(newUser *User) (*User, error) {
	return nil, nil
}

func (s *Service) Login(email *string, password *string) (*User, error) {
	return nil, nil
}

func (s *Service) Me() (*User, error) {
	return nil, nil
}