package repository

import (
	"decomposition/app"
	"decomposition/entities"
)

type StabStorage struct{}

var _ app.Storage = (*StabStorage)(nil)

func (s *StabStorage) GetAll() ([]*entities.Student, error) {
	result := []*entities.Student{
		{
			Name:  "Имя Фамилия 1",
			Age:   10,
			Grade: 1,
		},
		{
			Name:  "Имя Фамилия 2",
			Age:   20,
			Grade: 2,
		},
	}
	return result, nil
}

func (s *StabStorage) Put(student *entities.Student) error {
	return nil
}
