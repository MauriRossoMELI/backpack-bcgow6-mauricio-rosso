package main

func main() {
	channel1 := make(chan int)
	channel2 := make(chan int)
	channel3 := make(chan int)
	var products []Products
	products = append(products, Products{Name: "ProductTest", Price: 23, Quantity: 1})
	go sumProducts(products, 1, channel1)

	var services []Services
	services = append(services, Services{Name: "ServiceTest", Price: 55, WorkedMinutes: 35})
	go sumServices(services, 1, channel2)

	var maintenances []Maintenance
	maintenances = append(maintenances, Maintenance{Name: "MaintenanceTest", Price: 33})
	go sumMaintenance(maintenances, 1, channel3)
	<-channel1
	<-channel2
	<-channel3
	println("Execution ended")
}

type Products struct {
	Name     string
	Price    float64
	Quantity int
}

type Services struct {
	Name          string
	Price         float64
	WorkedMinutes int
}

type Maintenance struct {
	Name  string
	Price float64
}

func sumProducts(products []Products, i int, c chan int) float64 {
	total := 0.0
	for _, value := range products {
		total += value.Price * float64(value.Quantity)
	}
	c <- i
	println("TERMINO PRODUCTS")
	return total
}

func sumServices(services []Services, i int, c chan int) float64 {
	total := 0.0
	for _, value := range services {
		if value.WorkedMinutes < 30 {
			total += value.Price
		} else {
			total = value.Price * float64(value.WorkedMinutes)
		}
	}
	c <- i
	println("TERMINO SERVICES")
	return total
}

func sumMaintenance(maintenance []Maintenance, i int, c chan int) float64 {
	total := 0.0
	for _, value := range maintenance {
		total += value.Price
	}
	c <- i
	println("TERMINO MAINTENANCES")
	return total
}
