package main

import "fmt"

func main() {
	word := "hello"
	println("Cantidad de letras: ", len(word))
	println("Letras: ")
	for i := 0; i < len(word); i++ {
		fmt.Printf("%c\n", word[i])
	}
}
