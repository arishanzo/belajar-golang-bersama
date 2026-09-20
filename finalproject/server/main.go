package main

import (
	"finalproject/config"
	"finalproject/internal/handler"
	"finalproject/internal/repository"
	"finalproject/internal/service"
	"fmt"
	"net/http"
)

func main() {
	db, err := config.ConnectDB() // koneksi database

	// handle error
	if err != nil {
		panic("Gagal Koneksi Di Database: " + err.Error())
	}

	// konfigurasi auth
	userRepo := repository.NewUserRepository(db) // fungsi ini buat dipakai untuk urusan query ke tabel user seperti create, find, update, delete dll
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	// Pakai mux untuk http dan mapping URL Di GOlang
	mux := http.NewServeMux()

	//  Todo : Routing API disini
	mux.HandleFunc("/register", authHandler.Register)
	mux.HandleFunc("/login", authHandler.Login)

	fmt.Println("Server Berjalan di http://localhost:5000")
	http.ListenAndServe(":5000", mux)

}
