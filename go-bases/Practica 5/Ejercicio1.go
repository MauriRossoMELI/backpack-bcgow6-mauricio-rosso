package main

import (
	"fmt"
	"os"
)

func main() {
	values := fmt.Sprintf("%v, %.2f, %v,\n%v, %.2f, %v,", 333, 84.5, 6, 444, 105.6, 4)
	err := os.WriteFile("./testFileExercise1.txt", []byte(values), 777)
	if err != nil {
		println("An error was occured: ", err.Error())
		panic(err)
	}
	println("File saved successfully.")
}
