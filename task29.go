package main

import "fmt"

type Date struct {
	Day, Month, Year int
}

type Employee struct {
	FullName  string
	BirthDate Date
	IsWorker  bool
	// Рабочий
	Professions [5]string
	Categories  [5]int
	ProfCount   int
	// Программист
	University  string
	GradYear    int
	Languages   [5]string
	Proficiency [5]int
	LangCount   int
}

func inputEmployee() Employee {
	var e Employee
	fmt.Print("ФИО: ")
	fmt.Scan(&e.FullName)
	fmt.Print("Дата рождения (dd mm yyyy): ")
	fmt.Scan(&e.BirthDate.Day, &e.BirthDate.Month, &e.BirthDate.Year)
	fmt.Print("Рабочий? (1 - да / 0 - программист): ")
	var choice int
	fmt.Scan(&choice)
	e.IsWorker = choice == 1
	if e.IsWorker {
		fmt.Print("Количество профессий (не более 5): ")
		fmt.Scan(&e.ProfCount)
		for i := 0; i < e.ProfCount; i++ {
			fmt.Printf("Профессия %d: ", i+1)
			fmt.Scan(&e.Professions[i])
			fmt.Printf("Категория %d: ", i+1)
			fmt.Scan(&e.Categories[i])
		}
	} else {
		fmt.Print("Название вуза: ")
		fmt.Scan(&e.University)
		fmt.Print("Год окончания: ")
		fmt.Scan(&e.GradYear)
		fmt.Print("Количество языков (не более 5): ")
		fmt.Scan(&e.LangCount)
		for i := 0; i < e.LangCount; i++ {
			fmt.Printf("Язык %d: ", i+1)
			fmt.Scan(&e.Languages[i])
			fmt.Printf("Уровень владения %d: ", i+1)
			fmt.Scan(&e.Proficiency[i])
		}
	}
	return e
}

func main() {
	e := inputEmployee()
	fmt.Printf("%+v\n", e)
}
