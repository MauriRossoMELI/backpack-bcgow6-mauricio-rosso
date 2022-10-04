package file

import (
	"os"

	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
	"github.com/gocarina/gocsv"
)

type File struct {
	path string
}

func (f *File) Read() ([]service.Ticket, error) {
	data, err := os.OpenFile("./tickets.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer data.Close()

	tickets := []service.Ticket{}

	if err != nil { //I check again because I check the panic and keep the execution alive.
		panic(err)
	} else {
		err = gocsv.UnmarshalFile(data, &tickets)
		//csv.NewReader(data) option 2
	}

	//Option "IF" easier (check)
	// if err := gocsv.UnmarshalFile(data, &tickets); err != nil { // Load clients from file
	// 	panic(err)
	// }

	return tickets, err
}

func (f *File) Write(ticketsUpdated []service.Ticket) error {
	ticketsFile, err := os.OpenFile("./tickets.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer ticketsFile.Close()

	if err != nil { //I check again because I check the panic and keept the execution alive.
		panic(err)
	}

	tickets := ticketsUpdated
	//fmt.Printf("%v", &tickets)

	err = gocsv.MarshalFile(&tickets, ticketsFile) // I use this to save the CSV back to the file
	if err != nil {
		panic(err)
	}

	return err
}
