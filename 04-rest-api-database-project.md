# 📕 Part 4 — REST API, Database & Final Project

Materi:

15. REST API
16. Database
17. Project

---

# 15. REST API

REST API digunakan agar aplikasi dapat berkomunikasi melalui HTTP.

## HTTP Method

| Method | Fungsi |
|---|---|
| GET | Mengambil data |
| POST | Membuat data |
| PUT | Mengubah data |
| PATCH | Update sebagian data |
| DELETE | Menghapus data |

Contoh endpoint:

```text
GET    /api/products
GET    /api/products/1
POST   /api/products
PUT    /api/products/1
DELETE /api/products/1
```

---

## JSON

Response:

```json
{
  "id": 1,
  "name": "Laptop",
  "price": 10000000
}
```

Struct:

```go
type Product struct {
    ID    int     `json:"id"`
    Name  string  `json:"name"`
    Price float64 `json:"price"`
}
```

Encoding:

```go
json.NewEncoder(w).Encode(product)
```

Decoding:

```go
json.NewDecoder(r.Body).Decode(&product)
```

---

## HTTP Status Code

Gunakan status code yang sesuai:

```text
200 OK
201 Created
400 Bad Request
401 Unauthorized
403 Forbidden
404 Not Found
500 Internal Server Error
```

### 🎯 Latihan

Buat CRUD API Product.

---

# 16. Database

Go dapat digunakan dengan PostgreSQL, MySQL, SQLite, dan database lainnya.

Konsep dasar:

```text
Application
    ↓
Repository
    ↓
Database Driver
    ↓
PostgreSQL / MySQL
```

## Database Connection

Contoh menggunakan `database/sql`:

```go
db, err := sql.Open("postgres", connectionString)

if err != nil {
    log.Fatal(err)
}
```

Cek connection:

```go
err = db.Ping()

if err != nil {
    log.Fatal(err)
}
```

## Query

```go
rows, err := db.Query(`
    SELECT id, name, price
    FROM products
`)
```

Insert:

```go
_, err := db.Exec(`
    INSERT INTO products (name, price)
    VALUES ($1, $2)
`, product.Name, product.Price)
```

> Gunakan parameter query, jangan memasukkan input user langsung ke SQL untuk menghindari SQL Injection.

---

# 🏗️ Struktur Project Modern

Untuk project API yang lebih besar, gunakan struktur yang jelas.

```text
go-api/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── handler/
│   ├── service/
│   ├── repository/
│   ├── model/
│   └── middleware/
├── config/
├── migrations/
├── pkg/
├── go.mod
└── README.md
```

Alur:

```text
HTTP Request
     ↓
Handler
     ↓
Service
     ↓
Repository
     ↓
Database
```

---

# 17. Final Project

## 🚀 E-Commerce REST API

Buat backend E-Commerce menggunakan Go.

### Authentication

Endpoint:

```text
POST /api/register
POST /api/login
POST /api/logout
```

### Products

```text
GET    /api/products
GET    /api/products/:id
POST   /api/products
PUT    /api/products/:id
DELETE /api/products/:id
```

### Categories

```text
GET    /api/categories
POST   /api/categories
PUT    /api/categories/:id
DELETE /api/categories/:id
```

### Orders

```text
GET  /api/orders
GET  /api/orders/:id
POST /api/orders
```

### Database

Minimal table:

```text
users
categories
products
orders
order_items
```

Relasi:

```text
users
  │
  └── orders
        │
        └── order_items
              │
              └── products
                    │
                    └── categories
```

---

# 🎯 Final Project Requirements

Project harus memiliki:

- [ ] REST API
- [ ] CRUD
- [ ] Database
- [ ] Validation
- [ ] Error Handling
- [ ] Authentication
- [ ] Middleware
- [ ] Password Hashing
- [ ] Pagination
- [ ] Search
- [ ] Filtering
- [ ] JSON Response
- [ ] HTTP Status Code
- [ ] Environment Variables
- [ ] Logging
- [ ] Clean Project Structure

---

# 🧪 Testing API

Gunakan:

- Postman
- Bruno
- curl

Contoh:

```bash
curl http://localhost:8080/api/products
```

---

# 🐳 Optional: Docker

Tambahkan:

```text
Dockerfile
docker-compose.yml
```

Service:

```text
Go API
PostgreSQL
```

---

# 🚀 Deployment

Setelah project selesai:

```text
Local Development
       ↓
Git
       ↓
GitHub
       ↓
Docker
       ↓
Cloud Server
       ↓
Production API
```

---

# 🏆 Graduation Challenge

Buat aplikasi **Inventory Management API** dengan:

```text
Authentication
Products
Categories
Suppliers
Stock In
Stock Out
Stock History
Users
Reports
```

Target akhirnya:

```text
Frontend
   ↓
REST API Go
   ↓
Service Layer
   ↓
Repository
   ↓
PostgreSQL
```

Jika mampu menyelesaikan project ini secara mandiri, kamu sudah memiliki fondasi yang kuat untuk bekerja sebagai **Go Backend Developer**.
