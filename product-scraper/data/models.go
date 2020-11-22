package data

// Model for a Product
type Product struct {
	URL  string  `json:"url"`
	Name	string	`json:"name"`
	ImageURL	string	`json:"imageURL"`
	Description string `json:"description"`
	Price	string `json:"price"`
	TotalReviews int `json:"totalReviews"`
}

// Model for a Product along with URL
// type ProductDetail struct {
// 	URL  string  `json:"url" validate:"required" sql:"url"`
// 	Product Product `json:"product" validate:"required"`
// }