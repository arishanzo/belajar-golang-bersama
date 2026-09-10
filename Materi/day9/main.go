package day9

import (
	"errors"
	"fmt"
)

// di go menggunakan explicit untuk error handling
// contohnya error handling di go

func kalkulator(a, b float64) (float64, error) {

	// fungsi ini menerima 2 nilai a dan b dan mengembalikan nilai serta error apabila kondisi tidak sesuai

	// melakukan pengecekan
	if b == 0 {
		// sekarang menambahkan wrapping error di go disini
		// supaya error lebih informatif dan bisa dilacak asalnya

		baseErr := errors.New("Angka tidak bisa dibagi 0 atau hasil 0")

		// mengembalikan error
		// bungkus dengan konteks tambahan %f %w
		return 0, fmt.Errorf("Gagal Menghitung karena %f / %f: %w", a, b, baseErr)
	}

	return a / b, nil // nil di go itu artinya kosong / tidak ada nilai

}

// kita ke error handling dengan interface dan cutom

// definisi eror custom dengan struct
type validasiError struct {
	field string
	pesan string
}

// implementasi interface error
func (e *validasiError) Error() string {
	return fmt.Sprintf("Kesalahan validasi pada field '%s' : %s", e.field, e.pesan)
}

func validasiUmur(umur int) error {
	if umur < 18 {
		return &validasiError{
			field: "Umur",
			pesan: "Minimal 18 tahun",
		}
	}
	return nil // nill atau kosong di go
}

func Main() {
	// kita panggil disini

	// jika umurnya < 18 maka erorr kalau > 18 umur sesuai
	if err := validasiUmur(20); err != nil {
		fmt.Println("Terjadi Error :", err)
	} else {
		fmt.Println("Validasi berhasil umur sesuai.")
	}

	fmt.Println("...............................................")

	// kita eksekusi disini
	result, err := kalkulator(10, 0)

	if err != nil {
		// cetak error
		fmt.Println(err)

		return
	}

	fmt.Println(result)
}
