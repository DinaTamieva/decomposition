package app

import (
	"bufio"
	"decomposition/entities"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Storage interface {
	GetAll() ([]*entities.Student, error)
	Put(student *entities.Student) error
}

type App struct {
	storage Storage
}

func New(storage Storage) *App {
	return &App{
		storage: storage,
	}
}

func (app *App) Run() {
	app.enteringStudents()
	app.printStudents()
}

func (app *App) enteringStudents() {
	fmt.Println("Enter student's name, age, and grade separated by spaces or type 'exit' to quit:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text() // получаем строку
		if line == "exit" {
			fmt.Println("Final students data:")
			app.printStudents()
		} else {
			fmt.Println()
			fmt.Println("Enter student's name, age, and grade separated by spaces or type 'exit' to quit:")
			student, err := populate(line)

			if err != nil {
				fmt.Println("Неверные данные", err)
				continue
			}

			err = app.storage.Put(student)
			if err != nil {
				fmt.Println("Не удалось зачислить студента на курс:", err)
				continue

			}
			fmt.Println("Зачислен студент:", student)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка ввода вывода")
	}
}

func (app *App) printStudents() {
	students, err := app.storage.GetAll()
	if err != nil {
		fmt.Println("Ошибка при получении данных:", err)
		return
	}

	for _, v := range students {
		fmt.Println(v)
	}
}

func populate(line string) (*entities.Student, error) {
	arr := strings.Split(line, " ")

	if len(arr) != 3 {
		return nil, errors.New("неправильный ввод данных")
	}

	name := arr[0]

	age, err := strconv.Atoi(arr[1])
	if err != nil {
		return nil, err
	}

	grade, err := strconv.Atoi(arr[1])
	if err != nil {
		return nil, err
	}

	student := &entities.Student{
		Name:  name,
		Age:   age,
		Grade: grade,
	}
	return student, nil
}
