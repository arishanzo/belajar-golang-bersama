package main

import (
	"finalproject/config"
	"finalproject/repository"
	"fmt"
	"net/http"
)

func main() {
	db, err := config.ConnectDB() // koneksi database

	// handle error
	if err != nil {
		panic("Gagal Koneksi Di Database: " + err.Error())
	}

	// Pakai muc untuk http dan mapping URL Di GOlang
	mux := http.NewServeMux()

	// konfigurasi auth
	userRepo := repository.NewUserRepository(db) // fungsi ini buat dipakai untuk urusan query ke tabel user seperti create, find, update, delete dll

	//  Todo : Routing API disini

	fmt.Println("Server Berjalan di htpp://localhost:5000")
	http.ListenAndServe(":5000", mux)

}
