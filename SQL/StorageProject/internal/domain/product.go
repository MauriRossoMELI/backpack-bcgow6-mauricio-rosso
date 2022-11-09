package domain

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	Count       int     `json:"count"`
	Price       float32 `json:"price"`
	IdWarehouse int     `json:"id_warehouse"`
}
