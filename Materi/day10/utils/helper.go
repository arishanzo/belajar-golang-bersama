package utils // nama package utils
import "fmt"

// buat func say hello
func SayHello(name string) {

	// kita panggil disini
	formatted := formatName(name)
	fmt.Println("Hello, ", formatted)
}

// bikin formatName
func formatName(name string) string {
	// func formatname menerima name string dan mengembalikan tipe string
	return "Mr./Ms" + name
}
