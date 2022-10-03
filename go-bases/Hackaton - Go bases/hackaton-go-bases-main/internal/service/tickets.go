package service

import "fmt"

type Bookings interface {
	// Create create a new Ticket
	Create(t Ticket) (bookings, error)
	// Read read a Ticket by id
	Read(id int) (Ticket, error)
	// Update update values of a Ticket
	Update(id int, t Ticket) (bookings, error)
	// Delete delete a Ticket by id
	Delete(id int) (bookings, error)
}

type bookings struct {
	Tickets []Ticket
}

type Ticket struct {
	Id          int    `csv:"ticket_id"`
	Names       string `csv:"ticket_name"`
	Email       string `csv:"ticket_email"`
	Destination string `csv:"ticket_destination"`
	Date        string `csv:"ticket_date"`
	Price       int    `csv:"ticket_price"`
}

// NewBookings creates a new bookings service
func NewBookings(Tickets []Ticket) Bookings {
	return &bookings{Tickets: Tickets}
}

func (b *bookings) Create(t Ticket) (bookings, error) {
	b.Tickets = append(b.Tickets, t)
	return *b, nil
}

func (b *bookings) Read(id int) (Ticket, error) {
	ticketFound := Ticket{}
	for _, ticket := range b.Tickets {
		if ticket.Id == id {
			ticketFound = ticket
		}
	}
	return ticketFound, nil
}

func (b *bookings) Update(id int, t Ticket) (bookings, error) { //NO funciona con forrange (¿por que?)
	for i := 0; i < len(b.Tickets); i++ {
		if b.Tickets[i].Id == id {
			b.Tickets[i].Id = t.Id
			b.Tickets[i].Names = t.Names
			b.Tickets[i].Email = t.Email
			b.Tickets[i].Destination = t.Destination
			b.Tickets[i].Date = t.Date
			b.Tickets[i].Price = t.Price
		}
	}
	fmt.Printf("%v", b)
	return *b, nil
}

func (b *bookings) Delete(id int) (bookings, error) {
	////Remove record: OPTION 1
	//Set the record to remove (with id = -1)
	copyTickets := b.Tickets
	for i := 0; i < len(b.Tickets); i++ {
		if b.Tickets[i].Id == id {
			copyTickets[i].Id = -1
		}
	}

	b.Tickets = nil

	//Copy all the records and exclude the record with id = -1
	for i := 0; i < len(copyTickets); i++ {
		if copyTickets[i].Id != -1 {
			b.Tickets = append(b.Tickets, copyTickets[i])
		}
	}

	////Remove record: OPTION 2
	//b.Tickets = append(b.Tickets[0:index], b.Tickets[index+1:]...) me duplica el ultimo!

	return *b, nil
}
