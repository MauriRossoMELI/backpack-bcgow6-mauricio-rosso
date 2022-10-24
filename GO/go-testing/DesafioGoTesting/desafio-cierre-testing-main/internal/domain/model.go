package domain

type Product struct {
	ID          string
	SellerID    string `json:"seller_id"`
	Description string
	Price       float64
}
