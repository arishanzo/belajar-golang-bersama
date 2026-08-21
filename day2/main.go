package day2

import "fmt"

func Main() {

	// fmt.Println("....... Kondisi If Else .............")

	// var age int = 10

	// jika umur lebih dari 18 maka dewasa kalau tidak anda belum dewasa
	// if age >= 18 {
	// 	fmt.Println("Anda Dewasa")
	// } else {
	// 	fmt.Println("Anda Belum Dewasa")
	// }

	// fmt.Println("...........................................")
	// fmt.Println("...........Kondisi Pakai Switch............")

	// var day string = ""

	// switch day {
	// case "Monday":
	// 	fmt.Println("Ya Hari Ini Senin")
	// case "Tuesday":
	// 	fmt.Println("Ya Hari Ini Selesa")
	// default:
	// 	fmt.Println("Anda Belum Memasukan Nama Hari")
	// }

	fmt.Println("...........................................")
	fmt.Println("...........Latihan .......................")
	fmt.Println("...........................................")

	var score int = 50

	if score >= 90 && score <= 100 {
		fmt.Println("Nilai: A")
	} else if score >= 80 {
		fmt.Println("Nilai: B")
	} else if score >= 70 {
		fmt.Println("Nilai: C")
	} else if score >= 60 {
		fmt.Println("Nilai: D")
	} else {
		fmt.Println("Nilai: E")
	}

}
