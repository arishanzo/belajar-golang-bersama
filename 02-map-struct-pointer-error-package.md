# 📗 Part 2 — Data Structure & Code Organization

Materi:

6. Map
7. Struct
8. Pointer
9. Error Handling
10. Package & Module

---

# 06. Map

Map digunakan untuk menyimpan data key-value.

```go
user := map[string]string{
    "name":  "Aris",
    "email": "aris@example.com",
}
```

Mengakses:

```go
fmt.Println(user["name"])
```

Menambahkan:

```go
user["city"] = "Surabaya"
```

Menghapus:

```go
delete(user, "city")
```

### Cek Key

```go
value, exists := user["name"]

if exists {
    fmt.Println(value)
}
```

### 🎯 Latihan

Buat map produk:

```text
laptop -> 10000000
mouse  -> 150000
keyboard -> 500000
```

---

# 07. Struct

Struct digunakan untuk membuat tipe data sendiri.

```go
type User struct {
    Name  string
    Email string
    Age   int
}
```

Membuat object:

```go
user := User{
    Name:  "Aris",
    Email: "aris@example.com",
    Age:   25,
}
```

Akses:

```go
fmt.Println(user.Name)
```

### Method

```go
func (u User) SayHello() {
    fmt.Println("Hello", u.Name)
}
```

Pemakaian:

```go
user.SayHello()
```

### 🎯 Latihan

Buat struct:

```text
Product
- ID
- Name
- Price
- Stock
```

Kemudian buat method untuk menampilkan informasi produk.

---

# 08. Pointer

Pointer menyimpan alamat memory suatu value.

```go
number := 10

pointer := &number

fmt.Println(pointer)
fmt.Println(*pointer)
```

`&` digunakan untuk mendapatkan address.

`*` digunakan untuk mendapatkan value dari pointer.

### Pointer pada Function

```go
func changeValue(value *int) {
    *value = 100
}

number := 10

changeValue(&number)

fmt.Println(number)
```

### 🎯 Latihan

Buat function:

```go
updateStock(stock *int, amount int)
```

Function harus mengubah stock secara langsung.

---

# 09. Error Handling

Go menggunakan explicit error handling.

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }

    return a / b, nil
}
```

Penggunaan:

```go
result, err := divide(10, 0)

if err != nil {
    fmt.Println(err)
    return
}

fmt.Println(result)
```

### Custom Error

```go
errors.New("user not found")
```

### 🎯 Latihan

Buat function:

```go
findUser(id int)
```

Jika user tidak ditemukan, return error.

---

# 10. Package & Module

Package digunakan untuk mengorganisasi kode.

Struktur:

```text
myapp/
├── go.mod
├── main.go
└── utils/
    └── helper.go
```

Inisialisasi module:

```bash
go mod init myapp
```

Package:

```go
package utils

func Add(a int, b int) int {
    return a + b
}
```

Import:

```go
import "myapp/utils"
```

Penggunaan:

```go
result := utils.Add(10, 20)
```

### 🎯 Latihan

Buat struktur:

```text
myapp/
├── main.go
├── go.mod
├── math/
│   └── calculator.go
└── user/
    └── user.go
```

Pisahkan logic berdasarkan tanggung jawabnya.

---

# 🧪 Mini Project Part 2

## Product Management CLI

Buat aplikasi CLI:

```text
=== Product Management ===

1. List Product
2. Add Product
3. Find Product
4. Delete Product
5. Exit
```

Gunakan:

- Map
- Struct
- Pointer
- Error Handling
- Package
