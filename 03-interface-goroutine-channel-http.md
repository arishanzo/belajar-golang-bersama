# 📙 Part 3 — Advanced Go & Concurrency

Materi:

11. Interface
12. Goroutine
13. Channel
14. HTTP Server

---

# 11. Interface

Interface mendefinisikan behavior.

```go
type Payment interface {
    Pay(amount int)
}
```

Implementasi:

```go
type BankTransfer struct{}

func (b BankTransfer) Pay(amount int) {
    fmt.Println("Pay using bank transfer:", amount)
}
```

Penggunaan:

```go
func processPayment(payment Payment) {
    payment.Pay(100000)
}
```

### Multiple Implementation

```go
type EWallet struct{}

func (e EWallet) Pay(amount int) {
    fmt.Println("Pay using e-wallet:", amount)
}
```

Sekarang `BankTransfer` dan `EWallet` dapat digunakan sebagai `Payment`.

### 🎯 Latihan

Buat interface:

```go
Notification
```

dengan method:

```go
Send(message string)
```

Implementasikan:

```text
EmailNotification
WhatsAppNotification
SMSNotification
```

---

# 12. Goroutine

Goroutine digunakan untuk menjalankan function secara concurrent.

```go
go sayHello()
```

Contoh:

```go
func sayHello() {
    fmt.Println("Hello")
}

func main() {
    go sayHello()

    time.Sleep(time.Second)
}
```

### WaitGroup

Untuk menunggu beberapa goroutine selesai:

```go
var wg sync.WaitGroup

wg.Add(2)

go func() {
    defer wg.Done()
    fmt.Println("Task 1")
}()

go func() {
    defer wg.Done()
    fmt.Println("Task 2")
}()

wg.Wait()
```

### 🎯 Latihan

Buat 5 goroutine yang memproses 5 produk secara bersamaan.

---

# 13. Channel

Channel digunakan untuk komunikasi antar goroutine.

```go
numbers := make(chan int)
```

Mengirim:

```go
numbers <- 10
```

Menerima:

```go
value := <-numbers
```

Contoh:

```go
func calculate(result chan int) {
    result <- 10 + 20
}

func main() {
    result := make(chan int)

    go calculate(result)

    value := <-result

    fmt.Println(value)
}
```

### Buffered Channel

```go
channel := make(chan int, 3)

channel <- 10
channel <- 20
channel <- 30
```

### 🎯 Latihan

Buat producer yang mengirim 10 angka ke channel dan consumer yang membaca semuanya.

---

# 14. HTTP Server

Go memiliki package `net/http`.

Server sederhana:

```go
package main

import (
    "fmt"
    "net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello Go!")
}

func main() {
    http.HandleFunc("/", hello)

    http.ListenAndServe(":8080", nil)
}
```

Jalankan:

```bash
go run main.go
```

Buka:

```text
http://localhost:8080
```

### Routing

```go
http.HandleFunc("/users", users)
http.HandleFunc("/products", products)
```

### 🎯 Latihan

Buat HTTP Server dengan endpoint:

```text
GET /
GET /users
GET /products
```

---

# 🧪 Mini Project Part 3

## Concurrent Product Processor

Buat aplikasi yang:

1. Memiliki 10 produk
2. Memproses produk menggunakan Goroutine
3. Mengirim hasil menggunakan Channel
4. Menggunakan Interface untuk processor
5. Menyediakan endpoint HTTP

Contoh:

```text
GET /products/process
```

Response:

```json
{
  "message": "products processed successfully"
}
```
