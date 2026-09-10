package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
) // gorm salah satu ORM untuk database di Golang

// buat func koneksi db
func ConnectDB() (*gorm.DB, error) {

	// load file .env disini
	loadenv := godotenv.Load()

	if loadenv != nil {
		return nil,
			fmt.Errorf("Gagal load file .env : %w", loadenv)
	}

	// konfigurasi db sesuai .env
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil

}
