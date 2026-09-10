package day4

import "fmt"

func Main() {

	// lalu kita panggil disini
	SayHello()
	// lalu kita isi parameter disini
	var name string = "Aris"
	greet(name)

	fmt.Println("...........................")

	// kita isi parameter disini dengan inisiasi variabel langsung
	result := add(10, 20)
	// result berisikan function add bernilai 10 dan 20

	fmt.Println(result)

	// sekarang kita bikin kalkulator
	fmt.Println("...........................")
	fmt.Println("..........Kalkulator..............")

	tambah, kali := kalkulator(2, 5)

	fmt.Println("Hasil dari Tambah", tambah)
	fmt.Println("Hasil dari Kali", kali)

	// sekarang kita bikin kalkulator
	fmt.Println("...........................")
	fmt.Println("..........Bikin Diskon..............")

	price := 100000.0
	discount := 20.0

	totalHarga := Diskon(price, discount)

	fmt.Println("Harga Awal :", price)
	fmt.Println("Diskon  : ", discount, "%")
	fmt.Println("Harga Akhir  :", totalHarga)
}

// function diskon
func Diskon(price float64, discount float64) float64 {
	totaldiskon := price * discount / 100
	return price - totaldiskon
}

func kalkulator(a int, b int) (int, int) {
	// dimana function kalkulator ini
	// menerima paarameter a dan b dan mengembalikan nilai int
	return a + b, a * b
}

// KITA Bikin function add atau tambah
func add(a int, b int) int {
	return a + b
}

// kita bikin function sayhello
func SayHello() {
	fmt.Println("Hallo")
}

// sekarang kita kasih parameter di function
func greet(name string) {
	fmt.Println("Nama Saya", name)
}
