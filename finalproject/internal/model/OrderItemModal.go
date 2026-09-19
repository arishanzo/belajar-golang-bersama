package model

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model // pakai gorm.model

	// tak perlu mendefinisikan id tiap tabel gorm akan membuat kolom id otomatis tiap tabel

	// kolom tabel
	OrderID   uint `gorm:"index;not null"`
	ProductID int  `gorm:"index;not null"`

	Quantity int     `gorm:"not null"`
	Price    float64 `gorm:"not null"`

	// foreignkey
	Order   Order   `gorm:"constraint:OnUpdate:CASCADE, onDelete:CASCADE"`
	Product Product `gorm:"constraint:OnUpdate:CASCADE, onDelete:CASCADE"`

	// OnUpdate:CASCADE, onDelete:CASCADE artinya ketika Tabel Order dan Product dihapus atau update
	// OrderItem bakal kehapus atau update otomatis sesuai dengan id masing masing tabel

}
