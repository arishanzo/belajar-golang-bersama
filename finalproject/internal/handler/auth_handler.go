package handler

import (
	"encoding/json"
	"finalproject/internal/service"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// buat jwt screet
// Supaya aman kita simpan di .env key jwtnya
var jwtkey = []byte(os.Getenv("JWT_SCRET"))

// buat struct authHandler untuk menyimpan service
type AuthHandler struct {
	service service.AuthService // panggil service
}

// inject service ke handler
func NewAuthHandler(service service.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

// register handler
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req map[string]string
	json.NewDecoder(r.Body).Decode(&req)

	// masukan  req.boy email dan password
	user, err := h.service.Register(req["name"], req["email"], req["password"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(user)

}

// Login handler
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req map[string]string
	json.NewDecoder(r.Body).Decode(&req)

	// masukan  req.boy email dan password
	user, err := h.service.Login(req["email"], req["password"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// generate token JWT disini
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": user.Name,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, errToken := token.SignedString(jwtkey)

	if errToken != nil {
		http.Error(w, "Gagal generate Token", http.StatusInternalServerError)
		return
	}

	// cetak token dengan format json
	json.NewEncoder(w).Encode(map[string]string{
		"token": tokenString,
	})

}
