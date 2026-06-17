package main

import "fmt"

type Date struct {
	Day, Month, Year int
}

type YoungPerson struct {
	FullName  string
	BirthDate Date
	IsStudent bool
	// Школьник
	SchoolName string
	Grade      int
	// Студент
	UniversityName string
	Faculty        string
	GroupName      string
}

func inputPerson() YoungPerson {
	var p YoungPerson
	fmt.Print("ФИО: ")
	fmt.Scan(&p.FullName)
	fmt.Print("Дата рождения (dd mm yyyy): ")
	fmt.Scan(&p.BirthDate.Day, &p.BirthDate.Month, &p.BirthDate.Year)
	fmt.Print("Студент? (1 - да, 0 - нет): ")
	var choice int
	fmt.Scan(&choice)
	p.IsStudent = choice == 1
	if p.IsStudent {
		fmt.Print("Название вуза: ")
		fmt.Scan(&p.UniversityName)
		fmt.Print("Факультет: ")
		fmt.Scan(&p.Faculty)
		fmt.Print("Группа: ")
		fmt.Scan(&p.GroupName)
	} else {
		fmt.Print("Название школы: ")
		fmt.Scan(&p.SchoolName)
		fmt.Print("Класс: ")
		fmt.Scan(&p.Grade)
	}
	return p
}

func main() {
	p := inputPerson()
	fmt.Printf("%+v\n", p)
}
