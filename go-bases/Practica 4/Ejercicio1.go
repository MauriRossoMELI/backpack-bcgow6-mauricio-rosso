package main

import (
	"time"
)

func main() {
	student1 := Student{
		Name:    "Lionel",
		Surname: "Messi",
		Id:      35854345,
		Date:    time.Now(),
	}
	detail(student1)
}

type Student struct {
	Name    string
	Surname string
	Id      int
	Date    time.Time
}

func detail(student Student) {
	println("Student info: \n")
	println("Name: ", student.Name)
	println("Surname: ", student.Surname)
	println("Id Card: ", student.Id)
	println("Birthdate: ", student.Date.Format("01-02-2006"))
}
