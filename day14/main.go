package day14

import (
	"fmt"
	"net/http"
)

// Bikin Server http sederhana di go

// w http.ResponseWriter = objek untuk menulis ke klient
// r *http.Request = objek untuk membaca request dari client

func root(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Halaman Root") // menulis halaman root di browser

	// sekarang kita ke http request di htpp request ada

	// 1. Method yang digunakan client mengakses halaman
	fmt.Fprintln(w, "Method:", r.Method)

	// 2. Path Url / Url sekarang diakses user
	fmt.Fprintln(w, "Path : ", r.URL.Path)

	// 3. Header mengambil informasi browser atau alat yang digunakan client
	fmt.Fprintln(w, "Header:", r.Header.Get("User-Agent"))

}

// halaman user
func users(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Halaman User") // menulis halaman root di browser

	// 1. Method yang digunakan client mengakses halaman
	fmt.Fprintln(w, "Method:", r.Method)

	// 2. Path Url / Url sekarang diakses user
	fmt.Fprintln(w, "Path : ", r.URL.Path)

	// 3. Header mengambil informasi browser atau alat yang digunakan client
	fmt.Fprintln(w, "Header:", r.Header.Get("User-Agent"))

}

// halaman user
func products(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Halaman Products") // menulis halaman root di browser

}

func Main() {

	// konfigurasi Http Server disini
	http.HandleFunc("/", root) //mengambil func root atau berisikan func root

	// bikin router agar klient bisa mengakses lebih dari alamat atau halaman contohnya
	http.HandleFunc("/users", users)       // menampilkan halaman user
	http.HandleFunc("/products", products) // menampilkan halaman products

	fmt.Println("Server Berjalan di http://localhost:5000") //server akan berjalan di localhost:5000

	err := http.ListenAndServe(":5000", nil) // akan berhalan di port 5000

	// cek eror server atau handle error
	if err != nil {
		fmt.Println("Server Erorr:", err)
	}

}
