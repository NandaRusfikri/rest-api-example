package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLogin(t *testing.T) {
	// Membuat instance objek router Gin
	r := gin.Default()

	// Menggunakan mode test untuk menghindari logging
	gin.SetMode(gin.TestMode)

	// Menggunakan endpoint yang sama dengan aplikasi utama
	r.POST("/login", login)

	// Membuat request dengan payload JSON
	payload := `{"username": "admin", "password": "admin123"}`
	req, err := http.NewRequest("POST", "/login", strings.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}

	// Menyimpan response dari endpoint
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Memeriksa status code response
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Memeriksa body response
	expectedBody := `{"message":"Login successful"}`
	if w.Body.String() != expectedBody {
		t.Errorf("Expected response body %s, got %s", expectedBody, w.Body.String())
	}
}
