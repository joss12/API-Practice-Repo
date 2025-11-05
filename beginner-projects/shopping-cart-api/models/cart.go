package models

// Item represent a single product in the shopping cart
type Item struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

// cart represent the full sshopping cart (mock session-wide cart)
type Cart struct {
	Items []Item `json:"items"`
}
