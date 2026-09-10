package controllers

import (
	"belajar-golang-bersama/day15/models"
	"encoding/json"
	"fmt"
	"net/http"
)

// buat data produk disini
var products = []models.Product{
	{ID: 1, Name: "Mouse", Price: 10000},
	{ID: 2, Name: "Laptop", Price: 30000},
}

// menampilkan semua produk
func GetProduct(w http.ResponseWriter, r *http.Request) {

	// kirim respon lewat json ke klien
	w.Header().Set("Content-Type", "application/json") // type json
	json.NewEncoder(w).Encode(products)                // mengirim atau menerima data respon ke klient

}

// menampilkan produk sesuai id
func GetByProduct(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("ID") // ambil path url ID

	// validasi ID KOSONG
	if id == "" {
		http.Error(w, "ID Kosong", http.StatusBadRequest)
		return
	}

	// cari produk berdasarkan ID
	for _, p := range products {
		if fmt.Sprintf("%d", p.ID) == id {
			// kirim respon lewat json ke klien
			w.Header().Set("Content-Type", "application/json") // type json
			json.NewEncoder(w).Encode(p)                       // mengirim atau menerima data respon ke klient
			return
		}
	}

	// handle error
	http.Error(w, "Produk Tidak Ditemukan", http.StatusNotFound)

}

// menambakaan produk
func CreateProduct(w http.ResponseWriter, r *http.Request) {

	var product models.Product

	err := json.NewDecoder(r.Body).Decode(&product) // membaca request body yang dikirimkan lewat POST

	// handle error
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Simpan Produk KE Slice
	products = append(products, product) // menambah produk baru

	// kirim respon lewat json ke klien
	w.Header().Set("Content-Type", "application/json") // type json
	json.NewEncoder(w).Encode(product)                 // mengirim atau menerima data respon ke klient

}

// buat func PUt atau update
func UpdateProduct(w http.ResponseWriter, r *http.Request) {

	var updated models.Product

	err := json.NewDecoder(r.Body).Decode(&updated)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// cari produk berdasarkan ID
	for i, p := range products {
		if p.ID == updated.ID {
			products[i] = updated

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updated)
			return
		}
	}

	http.Error(w, "Produk Tidak ditemukan / Gagal Update", http.StatusNotFound)
}

// buat func Delete
func DeleteProduct(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("ID")

	// VALIDASI ID KOSONG
	if id == "" {
		http.Error(w, "ID Produk Kosong Isi ID terlebih dahulu", http.StatusBadRequest)
		return
	}

	// cari produk berdasarkan ID
	for i, p := range products {
		if fmt.Sprintf("%d", p.ID) == id {
			// hapus produk
			products = append(products[:i], products[i+1:]...)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": fmt.Sprintf("Produk Dengan ID %s berhasil dihapus", id),
			})
			return
		}
	}

	http.Error(w, "Produk Tidak ditemukan / Gagal Delete", http.StatusNotFound)
}
