package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model // pakai gorm.model

	// tak perlu mendefinisikan id tiap tabel gorm akan membuat kolom id otomatis tiap tabel

	// kolom tabel
	Name       string `gorm:"type:varchar(255);not null"` // wajib diisi
	Desc       string
	Price      float64 `gorm:"not null"`           // wajib diisi
	Stock      int     `gorm:"not null;default:0"` // nilai default 0
	CategoryID uint

	// foreignkey
	Category   Category    // child ke category
	OrderItems []OrderItem `gorm:"foreignKey:ProductID"` // Product punya banyak data OrderItem
}
