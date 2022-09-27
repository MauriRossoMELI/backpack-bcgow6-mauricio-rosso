package main

import "fmt"

func main() {
	storeAdded := newStore([]product{newProduct(small, "Bottle", 65.3)})
	fmt.Printf("%f", storeAdded.Total())
	fmt.Printf()
}

type store struct {
	products []product
}

type product struct {
	productType string
	name        string
	price       float64
}

type Product interface {
	CalculateCost() float64
}

type Ecommerce interface {
	Total() float64
	Add()
}

func newProduct(productType string, name string, price float64) product {
	newProd := product{
		productType: productType,
		name:        name,
		price:       price,
	}
	return newProd
}

func newStore(s []product) Ecommerce {

	return Ecommerce
}

func Total(productPrices []product) float64 {
	totalCost := 0.0
	for _, prod := range productPrices {
		switch prod.productType {
		case "small":
			totalCost += prod.price
		case "medium":
			aditionalCost := (prod.price * 3) / 100
			totalCost += prod.price + aditionalCost
		case "big":
			aditionalCost := (prod.price*6)/100 + 2500
			totalCost += prod.price + aditionalCost
		}
	}
	return totalCost
}

func Add(prod product) {
	productList := []product{
		prod.name,
		prod.price,
		prod.productType,
	}
}

const (
	small  = "Small"
	medium = "Medium"
	big    = "Big"
)
