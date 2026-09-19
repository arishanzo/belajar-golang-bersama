package main

import (
	"finalproject/config"
	"finalproject/internal/model"
	"fmt"
)

func main() {
	db, err := config.ConnectDB() // koneksi database

	// handle error
	if err != nil {
		panic("Gagal Koneksi Di Database: " + err.Error())
	}

	// auto migration semua tabel disini
	errMigration := db.AutoMigrate(&model.User{}, &model.Order{},
		&model.OrderItem{}, &model.Product{}, &model.Category{})

	// Handle error
	if errMigration != nil {
		fmt.Println("Database Migrasi Erorr: ", err)

	} else {
		fmt.Println("Database Migrasi Berhasil: ")

	}

}
