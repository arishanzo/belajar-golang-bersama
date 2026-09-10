package day5

import "fmt"

func Main() {

	// ada dua cara membuat array di go
	// cara pertama

	// numbers berisikan 3 array yang type datanya Integer
	// var numbers [3]int

	// kita masukan array disini
	// numbers[0] = 10
	// numbers[1] = 20
	// numbers[2] = 30

	// cara kedua bisa langsung begini

	// arr := [3]int{10, 20, 30}

	// fmt.Println("Array", numbers)
	// fmt.Println("Array 2", arr)

	// kita masuk ke bagian slice
	// slice itu seperti kotak transparan yang menyoroti sebagian rak array,

	// arraySlice := [5]int{10, 20, 30, 40, 50}

	// menggunakan slice dimulai index 1 sapai akhir
	// nilai := arraySlice[1:]

	// ambil array 1-4 saja
	// nilai2 := arraySlice[1:4]

	// fmt.Println("slice :", nilai)

	// fmt.Println("slice2 :", nilai2)

	// lalu perbaddan array sama slice gimana?
	// kalau array tanpa :
	// nilaiArray := arraySlice[0]

	// nilaiSlice := arraySlice[1:]

	// fmt.Println("Nilai array :", nilaiArray)

	// fmt.Println("Nilai slice :", nilaiSlice)

	// latihan menambahkan array

	numbers := []int{10, 20, 30}

	numbers = append(numbers, 40)
	numbers = append(numbers, 50)

	// kita coba print
	fmt.Println(numbers)

	// kita coba tambahin array bertipe string

	products := []string{
		"Laptop",
		"Mouse",
		"Keyboard",
		"Monitor",
	}

	fmt.Println("Produk sebelum ditambahkan")
	fmt.Println(products)

	products = append(products, "Headset")

	fmt.Println("\n Produk Setelah Ditambahkan")

	// range ini menghasillkan dua yaitu index sama value dari array produk
	for i, product := range products {
		fmt.Println(`index ke -`, i, product)
	}

}
