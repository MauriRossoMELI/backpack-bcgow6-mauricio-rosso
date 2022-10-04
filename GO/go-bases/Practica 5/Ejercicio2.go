package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("./testFileExercise1.txt")
	if err != nil {
		println("An error was occured: ", err.Error())
		panic(err)
	}
	stringData := string(data)
	fieldList := strings.Split(stringData, ",")

	//var totalAmount float64
	fmt.Printf("%s\t\t%10s\t\t%10s\n", "Id Producto", "Precio", "Cantidad")
	fmt.Printf("%s\t\t\t%10s\t\t%10s\n", fieldList[0], fieldList[1], fieldList[2])
	fmt.Printf("%s\t\t\t%10s\t\t%10s\n", fieldList[3], fieldList[4], fieldList[5])
}
