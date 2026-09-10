package day15

import (
	"belajar-golang-bersama/day15/controllers"
	"fmt"
	"net/http"
)

func Main() {

	// konfigurasi disini
	http.HandleFunc("/produk", controllers.GetProduct)           // GET
	http.HandleFunc("/produk/filter", controllers.GetByProduct)  // GET
	http.HandleFunc("/produk/create", controllers.CreateProduct) // POST
	http.HandleFunc("/produk/update", controllers.UpdateProduct) // PUT
	http.HandleFunc("/produk/delete", controllers.DeleteProduct) // DELETE

	fmt.Println("Server Berjalan di http://localhost:5000")
	err := http.ListenAndServe(":5000", nil)

	// cek erorr server
	if err != nil {
		fmt.Println("Server error : ", err)
	}

	// anda bisa pakai aplikasi postman untuk tes cek produk by id dan tambah data
	// disini saya pakai echo api
}
