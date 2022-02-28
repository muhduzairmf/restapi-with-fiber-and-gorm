package models

import "time"

type Order struct {
	ID        	  uint `json:"id" gorm:"primaryKey"`
	CreatedAt 	  time.Time `json:"createdAt"`
	ProductId 	  uint `json:"productId"`
	Product 	  Product `gorm:"foreignKey:ProductId"`
	AmountToOrder uint `json:"amountToOrder"`
	UserId 		  uint `json:"userId"`
	User 		  User `gorm:"foreignKey:UserId"`
}