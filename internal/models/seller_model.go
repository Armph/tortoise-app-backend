package models

// Seller model from seller collection
type Seller struct {
	ID      string   `json:"seller_id"`
	Name    string   `json:"name"`
	Surname string   `json:"surname"`
	Pets    []string `json:"pets"`
}
