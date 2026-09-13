package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model // pakai gorm.model

	// tak perlu mendefinisikan id tiap tabel gorm akan membuat kolom id otomatis tiap tabel

	// kolom tabel
	Name string `gorm:"type:varchar(255);not null"` // wajib diisi
	Desc string

	// foreignkey
	Products []Product `gorm:"foreignKey:CategoryID"`
}
