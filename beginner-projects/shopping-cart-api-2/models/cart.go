package models

type Item struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

// Cart represent the full shopping cart
type Cart struct {
	Items []Item `json:"items"`
}
