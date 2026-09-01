package day13

import "fmt"

// channel di go adalah alat utama untuk concurrency yang aman, sederhana, dan idiomatik
// biasa digunakan untuk komunkasi antar goroutine.

// ada dua tipe channel
// channel unbuffered yaitu setiap kirim harus ada penerima langsung
// channel buffered bisa menampung beberapa data dulu, penerima bisa ambil belakangan

//  buat func disini
func calculate(result chan int) {
	result <- 10 + 20 // mengirim berisikan 10 tambah 20 = 30
}

func Main() {

	// channel unbuffered
	result := make(chan int) // membuat channel

	go calculate(result) // menjalankan goroutine

	value := <-result // menerima pesan channel

	fmt.Println(value) // fmt.println bisa jalan karena da mekanisme sinkronisasi otomati dari channel

	// channel buffer
	ch := make(chan int, 3) // channel bisa menampung 3 buffer

	//kirim data  ke channel sebanyak 3 data
	ch <- 10
	ch <- 20
	ch <- 30

	// ambil data dari channel sebanyak 3 data
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// data yang keluar bakal sesuai urutan masuk fifo (first in, first out)

	// Latihan
	fmt.Println("--------------------------------------------------------------")
	// kita buat channel dengan buffer isi 10

	chLatihan := make(chan int, 10)

	//produser kirim 10 angka
	go func() {
		for i := 1; i <= 10; i++ {
			chLatihan <- i
			fmt.Println("Produser kirim sebanyak:", i)
		}
		close(chLatihan) // tutup channel setelah selesai looping
	}()

	// consumer baca semua angka
	for val := range chLatihan {
		fmt.Println("Consumer terima sebanyak:", val)
	}

}
