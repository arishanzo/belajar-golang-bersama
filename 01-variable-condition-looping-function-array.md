# 📘 Part 1 — Go Fundamentals

Materi:

1. Variable & Data Type
2. Condition
3. Looping
4. Function
5. Array & Slice

---

## 01. Variable & Data Type

### Variable

```go
package main

import "fmt"

func main() {
    var name string = "Aris"
    age := 25

    fmt.Println(name)
    fmt.Println(age)
}
```

### Data Type

Go memiliki beberapa tipe data dasar:

```go
string
int
float64
bool
```

Contoh:

```go
var name string = "Golang"
var age int = 10
var price float64 = 99.99
var active bool = true
```

### Constant

```go
const appName = "Go Learning"

fmt.Println(appName)
```

### 🎯 Latihan

Buat program yang menyimpan:

- Nama
- Umur
- Kota
- Saldo
- Status aktif

Kemudian tampilkan semuanya.

---

# 02. Condition

Condition digunakan untuk membuat keputusan.

```go
age := 20

if age >= 18 {
    fmt.Println("Dewasa")
} else {
    fmt.Println("Belum dewasa")
}
```

### Else If

```go
score := 85

if score >= 90 {
    fmt.Println("A")
} else if score >= 80 {
    fmt.Println("B")
} else {
    fmt.Println("C")
}
```

### Switch

```go
day := "Monday"

switch day {
case "Monday":
    fmt.Println("Hari Senin")
case "Tuesday":
    fmt.Println("Hari Selasa")
default:
    fmt.Println("Hari lainnya")
}
```

### 🎯 Latihan

Buat program menentukan grade:

```text
90 - 100 = A
80 - 89  = B
70 - 79  = C
60 - 69  = D
< 60     = E
```

---

# 03. Looping

Go menggunakan `for` sebagai looping utama.

```go
for i := 1; i <= 5; i++ {
    fmt.Println(i)
}
```

### Infinite Loop

```go
for {
    fmt.Println("Running...")
}
```

### Loop dengan Condition

```go
i := 1

for i <= 5 {
    fmt.Println(i)
    i++
}
```

### Range

```go
names := []string{"Aris", "Budi", "Citra"}

for index, name := range names {
    fmt.Println(index, name)
}
```

### 🎯 Latihan

Buat program untuk menampilkan angka:

```text
1
2
3
...
100
```

Kemudian tampilkan hanya angka genap.

---

# 04. Function

Function digunakan untuk membuat kode yang reusable.

```go
func sayHello() {
    fmt.Println("Hello Go")
}
```

Pemanggilan:

```go
sayHello()
```

### Parameter

```go
func greet(name string) {
    fmt.Println("Hello", name)
}
```

### Return Value

```go
func add(a int, b int) int {
    return a + b
}

result := add(10, 20)
fmt.Println(result)
```

### Multiple Return

```go
func calculate(a int, b int) (int, int) {
    return a + b, a * b
}

sum, multiply := calculate(10, 5)
```

### 🎯 Latihan

Buat function:

```text
calculateDiscount(price, discount)
```

Function mengembalikan harga setelah diskon.

---

# 05. Array & Slice

## Array

Array memiliki ukuran tetap.

```go
var numbers [3]int

numbers[0] = 10
numbers[1] = 20
numbers[2] = 30
```

Atau:

```go
numbers := [3]int{10, 20, 30}
```

## Slice

Slice lebih fleksibel.

```go
numbers := []int{10, 20, 30}

numbers = append(numbers, 40)
```

### Mengakses Slice

```go
fmt.Println(numbers[0])
fmt.Println(numbers[1:3])
```

### Length & Capacity

```go
fmt.Println(len(numbers))
fmt.Println(cap(numbers))
```

### 🎯 Latihan

Buat slice daftar produk:

```text
Laptop
Mouse
Keyboard
Monitor
```

Tambahkan satu produk baru menggunakan `append`.

---

# 🧪 Mini Project Part 1

## Simple Student Grade

Buat program CLI sederhana:

```text
=== Student Grade ===

Nama  : Aris
Nilai : 85

Grade : B
Status: LULUS
```

Gunakan:

- Variable
- Data Type
- Condition
- Function
- Slice

### ⭐ Challenge

Buat agar program dapat menyimpan minimal 5 siswa dan menghitung rata-rata nilai.
