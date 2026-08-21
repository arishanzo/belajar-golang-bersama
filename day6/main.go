package day6

import "fmt"

func Main() {

	// Map sering digunakna untuk menyimpan data key unik dan value didalamnya

	// map ini menyimpan data bertipe string
	user := map[string]string{
		"name": "Aris",
		"age":  "25",
	}

	// mari kita cetak
	fmt.Println(user["name"])
	fmt.Println(user["age"])

	// kita coba cek key apakah bertipe false apa true
	// jika ada maka tipenya true

	// value merujuk key di dalam map dan exists merujuk pengecekan bertipe boolean
	value, exists := user["email"]
	// mari kita cetak
	fmt.Println(value, exists)

	//hapus key dalam map
	delete(user, "age")
	// walaupun sudah kita hapus pakai delete key age didalam map masih tetap ada
	// akan tetapi ketika kita run kembali key age sudah tidak ada

	// jumlah elemen map
	fmt.Println(len(user))

	// cara menambahkan key map
	user["city"] = "Lamongan"

	fmt.Println(user)

	fmt.Println("--------------------------------------------")
	fmt.Println("Latihan")

	// kita  buat produk

	products := map[string]interface{}{
		"Laptop":   10000000,
		"Mouse":    1500000,
		"keyboard": 50000,
	}

	// kita coba looping
	for name, price := range products {
		// jadi hasil dari looping ini mencetak name sebagai key, dan price sebagai value dari map
		fmt.Println(name, price)
	}

	// Selanjutnya pakai type interface
	// interface bis menampung tipe apa aja sebagi contoh

	// dengan type assertion
	keyboard := products["keyboard"].(int)
	fmt.Println(keyboard * 10000)

}
