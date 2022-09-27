package main

import "fmt"

func main() {
	salary := calculateSalary(9600, CategoryA)
	fmt.Printf("%f", salary)
}

const (
	CategoryA = "A"
	CategoryB = "B"
	CategoryC = "C"
)

func calculateSalary(minutes int, category string) float64 {
	workedHours := float64(minutes / 60)
	var salary float64
	switch category {
	case CategoryA:
		salary = workedHours * 3000
		additional := (salary * 50) / 100
		salary += additional
		return salary
	case CategoryB:
		salary = workedHours * 1500
		additional := (salary * 20) / 100
		salary += additional
		return salary
	case CategoryC:
		salary = workedHours * 1000
		return salary
	}
	return 0
}
