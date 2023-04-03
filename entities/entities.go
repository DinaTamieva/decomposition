package entities

import "fmt"

type Student struct {
	Name  string
	Age   int
	Grade int
}

func (s *Student) String() string {
	return fmt.Sprintf("Имя %v, возраст %v, оценка %v", s.Name, s.Age, s.Grade)
}
