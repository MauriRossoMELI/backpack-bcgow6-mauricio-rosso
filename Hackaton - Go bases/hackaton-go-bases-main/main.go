package main

import (
	"fmt"

	"github.com/bootcamp-go/hackaton-go-bases/internal/file"
	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

func main() {
	var tickets []service.Ticket
	// Funcion para obtener tickets del archivo csv
	service.NewBookings(tickets)

	fileManager := file.File{}
	tickets, err := fileManager.Read()

	if err != nil {
		panic(err.Error())
	}

	//READ a ticket by id
	ticket, err := service.NewBookings(tickets).Read(3)
	if err != nil {
		panic("An error was occured getting the specified ticket!")
	}
	fmt.Printf("TICKET 3: %v \n", ticket)

	//WRITE a ticket in the file
	// newTicket := service.Ticket{Id: 1001, Names: "TESTNAME", Email: "test@test.com", Destination: "Argentina", Date: "23:43", Price: 333}
	// ticketsEdited, err := service.NewBookings(tickets).Create(newTicket)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// errWrite := fileManager.Write(ticketsEdited.Tickets)
	// if errWrite != nil {
	// 	panic(err.Error())
	// }

	//UPDATE a ticket in the file
	// newTicketToUpdate := service.Ticket{Id: 1000, Names: "TESTUPDATE", Email: "testupd@testupd.com", Destination: "ArgentinaUpdated", Date: "11:44", Price: 123}
	// ticketsEdited, err := service.NewBookings(tickets).Update(1000, newTicketToUpdate)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// errUpdate := fileManager.Write(ticketsEdited.Tickets)
	// if errUpdate != nil {
	// 	panic(err.Error())
	// }

	//DELETE a ticket in the file
	// ticketsEdited, err := service.NewBookings(tickets).Delete(50)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// errUpdate := fileManager.Write(ticketsEdited.Tickets)
	// if errUpdate != nil {
	// 	panic(err.Error())
	// }
}
