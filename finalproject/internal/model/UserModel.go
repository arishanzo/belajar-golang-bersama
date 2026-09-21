package model

import "gorm.io/gorm"

type User struct {
	gorm.Model // pakai gorm.model

	// tak perlu mendefinisikan id tiap tabel gorm akan membuat kolom id otomatis tiap tabel

	// kolom tabel
	Name     string `gorm:"type:varchar(255);not null"`             // wajib diisi
	Email    string `gorm:"type:varchar(255);uniqueIndex;not null"` // wajib diisi dan email harus berbeda dengan menambahkan uniqueindex
	Password string `gorm:"type:varchar(255);not null"`              // panjang karakter 8 karakter dan wajib disi

	// foreignkey
	Orders []Order `gorm:"foreignKey:UserID"` // User mempunyai banyak data Order atau relasi user > order
}
