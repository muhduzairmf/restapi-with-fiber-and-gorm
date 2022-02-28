package models

import "time"

type User struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"createdAt"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email 	  string `json:"email" gorm:"unique"`
}