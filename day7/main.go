package day7

import "fmt"

/* ## Day 7 – Struct

   ### Penjelasan
   struct di Golang digunakan untuk membuat tipe data bentukan
   yang bisa menampung beberapa field dengan tipe berbeda.
   Cocok untuk merepresentasikan objek nyata seperti User, Product, atau Book.

   Keuntungan:
   - Data lebih terorganisir
   - Bisa gabungkan berbagai tipe data
   - Lebih aman dibanding map[string]interface{}
*/

// definisikan struct
type User struct {
	Name   string
	Age    int
	Active bool
}

func Main() {

	// membuat instance struct
	u := User{
		Name:   "Aris",
		Age:    27,
		Active: true,
	}

	// cara aksesnya
	fmt.Println("Nama :", u.Name)
	fmt.Println("Umur :", u.Age)
	fmt.Println("Status :", u.Active)

	// update field struct
	// u.Age = 80
	fmt.Println("Umur Setelah Update :", u.Age)

	u.Greet()

}

// Struct di pakai bersama method
func (u User) Greet() {
	fmt.Println("Hallo, Nama Saya ", u.Name)
	fmt.Println("Umur Saya Saat Ini  ", u.Age)
}
