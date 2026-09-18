package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model

	UserID     uint    `gorm:"index;not null"`
	TotalPrice float64 `gorm:"not null"`

	// foreignKey
	User  User
	Items []OrderItem `gorm:"foreignKey:OrderID"` // oRDER punya banyak data OrderItem atau relasi ke OrderItem
}
