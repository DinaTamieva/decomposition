package repository

import (
	"decomposition/app"
	"decomposition/entities"
	"errors"
)

type MemStorage struct {
	studentsByName map[string]*entities.Student
}

var _ app.Storage = (*MemStorage)(nil)

func NewMemStorage() *MemStorage {
	return &MemStorage{
		studentsByName: make(map[string]*entities.Student),
	}
}

func (ms *MemStorage) GetAll() ([]*entities.Student, error) {
	var students []*entities.Student
	for _, v := range ms.studentsByName {
		students = append(students, v)
	}
	return students, nil
}

func (ms *MemStorage) Put(student *entities.Student) error {
	if _, found := ms.studentsByName[student.Name]; !found {
		ms.studentsByName[student.Name] = student
		return nil
	}
	return errors.New("студент с таким именем уже зачислен")
}
