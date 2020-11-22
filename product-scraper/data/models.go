package data

type Product struct {
	Name	string	`json:"name"`
	ImageURL	string	`json:"imageURL"`
	Description string `json:"description"`
	Price	string `json:"price"`
	TotalReviews int `json:"totalReviews"`
}