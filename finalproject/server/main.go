package main

import (
	"finalproject/config"
	"fmt"
	"net/http"
)

func main() {
	_, err := config.ConnectDB() // koneksi database

	// handle error
	if err != nil {
		panic("Gagal Koneksi Di Database: " + err.Error())
	}

	// Pakai muc untuk http dan mapping URL Di GOlang
	mux := http.NewServeMux()

	//  Todo : Routing API disini

	fmt.Println("Server Berjalan di htpp://localhost:5000")
	http.ListenAndServe(":5000", mux)

	// brarti koneksi database mysql sudah berhasil

}
