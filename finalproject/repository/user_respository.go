package repository

import (
	"finalproject/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	// buat kontrak method untuk membuat user baru
	Create(user model.User) (model.User, error)

	// buat kontrak method untuk mencari Name dan Email
	FindByName(name string) (model.User, error)
	// findByEmail(email string) (model.User, error)
}

// interface userrepository ini untuk mendefinisikan aturan yang harus diikuti oleh implementasi Repository

// buat struct userRepo
type userRepo struct {
	db *gorm.DB // struct userRepo Menyimpan koneksi database (pointer ke GORM DB)
}

// buat func new user repository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{db: db} // bikin object userRepo dengan koneksi DB,
	// Lalu mengembalikan sebagai func Userrepository
}

func (r *userRepo) Create(user model.User) (model.User, error) {

	err := r.db.Create(&user).Error // eksekusi query insert ke tabel users.

	return user, err // mengembalikan user yang dibuat dan erorr handle jika ada
}

func (r *userRepo) FindByName(name string) (model.User, error) {

	var user model.User // membuat variabel user bertipe model.user

	err := r.db.Where("name = ?", name).First(&user).Error

	return user, err
}

// method findbyemail untuk mencari berdasarkan kolom email
func (r *userRepo) FindByEmail(email string) (model.User, error) {

	var user model.User // membuat variabel user bertipe model.user

	err := r.db.Where("email = ?", email).First(&user).Error

	return user, err
}
