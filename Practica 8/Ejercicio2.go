package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	// fileNumber, err := GenerateFileNumber()
	// if err != nil {
	// 	panic("NUMERO DE LEGAJO NULO: " + fileNumber)
	// }
	defer func() { //EL DEFER SIEMPRE VA LUEGO DE UNA ACCION CRITICA, INMEDIATAMENTE.
		err := recover()
		if err != nil {
			fmt.Println(err, "")
		}
	}()

	CheckCustomer()
	cust := GenerateCustomerData()
	_, err := ValidateNonZeroValues(&cust)
	if err != nil {
		panic(err)
	}
}

func CheckCustomer() {
	_, err := os.OpenFile("./customers.txt", os.O_RDONLY, 0644)

	if err != nil {
		panic("El archivo indicado no fue encontrado o está dañado.")
	}
}

func ValidateNonZeroValues(customer *Customer) (bool, error) {
	if customer.Name != "" && customer.Surname != "" && customer.Address != "" && customer.FileNumber != "" && customer.IdCard != 0 && customer.PhoneNumber != "" {
		return true, nil
	}
	return false, errors.New("Existen valores nulos!")
}

func GenerateCustomerData() Customer {
	return Customer{Name: "Pablo", Surname: "", IdCard: 39500667, PhoneNumber: "3493559988", Address: "Los robles 36"}
}

func GenerateFileNumber() (string, error) {
	fileNum := "F333"
	if fileNum == "F333" {
		return "", nil
	}
	return fileNum, nil
}

type Customer struct {
	FileNumber  string
	Name        string
	Surname     string
	IdCard      int
	PhoneNumber string
	Address     string
}
