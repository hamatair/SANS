package entity

import "github.com/google/uuid"

type Product struct {
	ID          uuid.UUID `json:"id" gorm:"primary_key;"`
	Name        string    `json:"name" gorm:"type:varchar(255);not null;" binding:"required"`
	Description string    `json:"description" gorm:"type:varchar(255);not null;" binding:"required"`
	Price       float64   `json:"price" gorm:"type:not null;" binding:"required"`
	CategoryID  int       `json:"category_id"`
	Stock       int       `json:"stock"`
	ImageURL    string    `json:"image_url" gorm:"type:varchar(200);"`
	Rating      float64   `json:"rating"`
	Weight      float64   `json:"weight"`
	Discount    float64   `json:"discount"`
}
