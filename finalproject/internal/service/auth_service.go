package service

import (
	"errors"
	"finalproject/internal/model"
	"finalproject/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

// buat interface AuthService = kontrak method yang harus ada
type AuthService interface {
	Register(name, email, password string) (model.User, error)
	Login(email, password string) (model.User, error)
}

// buat struct authService untuk menyimpan repository
type authService struct {
	repo repository.UserRepository
}

// memasukkan repository ke dalam service di func NewAuthService
func NewAuthService(repo repository.UserRepository) AuthService {
	return &authService{repo: repo}
}

// method register untuk daftar user baru
func (s *authService) Register(name, email, password string) (model.User, error) {

	// cek apakah name sudah ada
	_, errName := s.repo.FindByName(name)
	if errName == nil {
		return model.User{}, errors.New("Name sudah terdaftar")
	}

	// cek apakah email sudah ada
	_, errEmail := s.repo.FindByEmail(email)
	if errEmail == nil {
		return model.User{}, errors.New("Name sudah terdaftar")
	}

	// Hash Password dengan bcrypt
	hashed, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	// handle error password
	if err != nil {
		return model.User{}, err
	}

	// jika tidak buat user baru
	user := model.User{Name: name, Email: email, Password: string(hashed)}
	return s.repo.Create(user) // buat data user baru

}

// buat method login validasi login user
func (s *authService) Login(email, password string) (model.User, error) {
	user, err := s.repo.FindByEmail(email) // cari email user

	if err != nil {
		return model.User{}, errors.New("User tidak ditemukan, silahkan buat akun")
	}

	// cek juga password dengan bcrypt
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return model.User{}, errors.New("Password Salah")
	}

	return user, nil

}
