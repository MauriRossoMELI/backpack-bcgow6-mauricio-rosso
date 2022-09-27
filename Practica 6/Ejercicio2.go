package main

import "fmt"

func main() {
	Product := newProduct("bottle (original)", 50.35)
	newUser := User{Name: "Test name", Surname: "Test surname", Email: "testemail@testing.com"}
	fmt.Printf("User BEFORE add the product: %v\n", newUser)
	addProduct(&newUser, &Product)
	fmt.Printf("User AFTER add the product: %v\n", newUser)
	fmt.Printf("User BEFORE delete the product: %v\n", newUser)
	deleteProduct(&newUser)
	fmt.Printf("User AFTER delete the product: %v\n", newUser)
}

type User struct {
	Name     string
	Surname  string
	Email    string
	Products []Product
}

type Product struct {
	Name     string
	Price    float64
	Quantity int
}

func newProduct(name string, price float64) Product {
	newProd := Product{Name: name, Price: price, Quantity: 33}
	return *&newProd
}

func addProduct(user *User, product *Product) {
	*&user.Products = append(*&user.Products, *product)
}

func deleteProduct(user *User) {
	*&user.Products = nil
}
