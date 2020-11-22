package data

type Product struct {
	Name 	 string	`json:"name" validate:"required" sql:"name"`
	ImageURL string	`json:"imageURL" validate:"required" sql:"imageURL"`
	Description string `json:"description" validate:"required" sql:"description"`
	Price	 string  `json:"price" validate:"required" sql:"price"`
	TotalReviews int `json:"totalReviews" validate:"required" sql:"totalReviews"`
}

type ProductDetail struct {
	ID	 string  `json:"id" sql:"id"`
	URL  string  `json:"url" validate:"required" sql:"url"`
	Product Product `json:"product" validate:"required"`
}