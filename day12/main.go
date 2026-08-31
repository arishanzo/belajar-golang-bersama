package day12

import (
	"fmt"
	"sync"
	"time"
)

// goroutine digunakan untuk menjalankan function secara concurrent / bersamaan

func sayHello() {
	fmt.Println("Hello World")
}

// sekarang di dunia nyata kita gak bisa pakai time.sleep kita gak tau seberapa lama kita menuggu untuk di jalankan
// solusinya kita pakai waitGroup

// apa itu wait group yaitu untuk menunggu beberapa goroutine selesai
func sayHelloWaitGroup(id int, wg *sync.WaitGroup) {
	defer wg.Done()                         // dibuat untuk mengurangi gorountine 1 tiap dipanngil
	fmt.Println("Hello dari goroutine", id) // mencetak sayhello ke id berapa
}

func Main() {

	// kita buat variabel wg disini
	var wg sync.WaitGroup
	wg.Add(2) // artinya menambahkan 2 goroutine

	// karena ada 2 goroutine
	go sayHelloWaitGroup(1, &wg) // &wg mengambil alamat pointer
	go sayHelloWaitGroup(2, &wg)

	wg.Wait() // artinya menunggu sampai jumlah goroutine jadi 0 atau selesai
	fmt.Println("Semua goroutine Selesai")

	fmt.Println("............................................")

	go sayHello() // keyword go membuat fungsi sayHello() dijalankan sebagai goroutine.

	// terjadi kesalahan kenapa sayHello() tidak dicetak karena kita belum memberi jeda atau time.sleep
	time.Sleep(time.Second) // memberi jeda 1 detik
	// sangat penting karena goroutine sayHello() punya kesempayan untuk dijalankan

	fmt.Println("Selesai")

}

// jadi pakai waitGroup jadi solusi terbaik daripada pakai time.sleep jika di dunia nyata

// perlu diingat bahwa method dari type sync.Waitgroup punya add, done, wait

// add untuk menambahkan jumlah goroutine yang akan di tunggu
//  done untuk mengurangi jumlah goroutine 1 sampai 0 atau selesai
//  wait untuk menunggu jumlah goroutine sampai 0 atau selesai setelah itu bisa dieksekusi bari code selanjutnya
