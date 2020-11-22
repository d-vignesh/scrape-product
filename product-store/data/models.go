package data

import (
	"time"
)

// Model for a Product	
type Product struct {
	ID		 string `json:"id" sql:"id"`
	URL 	 string `json:"url" validate:"required" sql:"url"`
	Name 	 string	`json:"name" validate:"required" sql:"name"`
	ImageURL string	`json:"imageURL" validate:"required" sql:"imageURL"`
	Description string `json:"description" validate:"required" sql:"description"`
	Price	 string  `json:"price" validate:"required" sql:"price"`
	TotalReviews int `json:"totalReviews" validate:"required" sql:"totalReviews"`
	CreatedAt time.Time `json:"createdAt" sql:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" sql:"updatedAt"`
}

// type ProductDetail struct {
// 	ID	 string  `json:"id" sql:"id"`
// 	URL  string  `json:"url" validate:"required" sql:"url"`
// 	Product Product `json:"product" validate:"required"`
// }