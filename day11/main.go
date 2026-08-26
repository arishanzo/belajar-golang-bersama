package day11

import "fmt"

// interface mendifinisikan kontrak yang harus ada method didalamnya contoh Pay(amount int)

type Payment interface {
	Pay(amount int)
}

// kita buat struct fungsinya menyediakan method sesuai kontrak
type BankTransfer struct{}

// kita buat func yang dimana ada (b BankTransfer) yang kita sebut reciever.
// fungsi receiver ini membuat fungsi Pay menjadi method miliknya struct BankTransfer

func (b BankTransfer) Pay(amount int) {

	fmt.Println("Pembayaran menggunakan bank transfer :", amount)
}

// kita buat func proses Pembayaran
func processPayment(p Payment, amount int) {
	p.Pay(amount)
}

// selanjutnya kita mencoba Polymorphism dengan slice

type Notification interface {
	Send(message string) // pesan / send bertipe string
}

// buat struct disini
type WaNotification struct{}

func (w WaNotification) Send(message string) {
	fmt.Println("Notifikasi dikirim lewat Wa :", message)
}

type EmailNotification struct{}

func (w EmailNotification) Send(message string) {
	fmt.Println("Notifikasi dikirim lewat Email :", message)
}

// buat fungsi proses notifikasi
func processNotification(n Notification, message string) {
	n.Send(message)
}

func Main() {

	// kita eksesuki disini
	bank := BankTransfer{}
	processPayment(bank, 50000)

	// alurnya kita beri nilai amount di processPayment dan Bank berisi funnc BankTransfer
	// di Struct BankTransfer mengambil method miliknya Payment yaitu Pay

	//  jadi ketika kita cetak atau print akan diproses pembayaran menggunakan bank transfer sebesar nilai kita masukan tadi

	fmt.Println(".................................................")

	// KITA eksesukis disini mengunakan slice
	notifications := []Notification{
		WaNotification{},
		EmailNotification{},
	}

	// kita looping datanya jika lebih dari 1 datanya
	for _, n := range notifications {
		// jadi ini isi dari pesan
		n.Send("Sitem akan kirim dalam waktu 10 jam !")
	}
}
