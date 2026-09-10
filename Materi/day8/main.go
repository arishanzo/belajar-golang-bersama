package day8

import (
	"fmt"
)

func Main() {

	// pointer adalah menyimpan alamat memory suatu value / nilai.

	x := 10
	p := &x // p menyimpan alamat memory dari x / & digunakan untuk mendapatkan address

	fmt.Println("Nilai x :", x)
	fmt.Println("Alamat x:", p) // pointer
	fmt.Println("Nilai p:", *p) // * digunakan untuk mendapatkan value dari pointer

	// lalu cara ubah nilai lewat pointer yaitu
	*p = 20 // * ambil nilai value
	fmt.Println("Nilai Baru x :", x)

	fmt.Println(".........................................")

	b := 10
	changePointer(&b)
	fmt.Println("Hasil changePinter, nilai b:", b)

	a := 10
	changeBiasa(a)
	fmt.Println("Hasil changeBiasa, nilai a:", a)

	fmt.Println(".........................................")
	person := Person{
		Name: "Aris",
	}

	// kita panggil function Tadi
	person.Rename("Budi")
	fmt.Println("Nama Baru :", person.Name)
}

// lalu perbedaan function biasa tanpa pointer
func changeBiasa(x int) {
	x = 99 // hanya ubah salinan
}

// pointer di function
func changePointer(p *int) {
	*p = 99 // ubah nilai disini
}

// sekarang kombinasi pointer dengan struct
type Person struct {
	Name string
}

// nilai pointer menerima type string
func (p *Person) Rename(newName string) {
	p.Name = newName
}
