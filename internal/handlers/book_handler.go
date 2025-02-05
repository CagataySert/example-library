package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/CagataySert/library-system/internal/database"
	"github.com/CagataySert/library-system/internal/models"
	"github.com/gorilla/mux"
)

// Kitapları Listele
func GetBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book
	if err := database.DB.Find(&books).Error; err != nil {
		http.Error(w, "Veritabani hatasi", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(books); err != nil {
		http.Error(w, "JSON encode hatası", http.StatusInternalServerError)
		return
	}
}

// Yeni Kitap Ekle
func AddBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Geçersiz JSON", http.StatusBadRequest)
		return
	}

	if err := database.DB.Create(&book).Error; err != nil {
		http.Error(w, "Kitap eklenirken hata oluştu", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(book); err != nil {
		http.Error(w, "JSON encode hatası", http.StatusInternalServerError)
		return
	}
}

// Kitap Güncelle
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Geçersiz ID", http.StatusBadRequest)
		return
	}

	var book models.Book
	if err := database.DB.First(&book, id).Error; err != nil {
		http.Error(w, "Kitap bulunamadı", http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Geçersiz JSON", http.StatusBadRequest)
		return
	}

	database.DB.Save(&book)
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(book); err != nil {
		http.Error(w, "JSON encode hatası", http.StatusInternalServerError)
		return
	}
}

// Kitap Sil
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Geçersiz ID", http.StatusBadRequest)
		return
	}

	if err := database.DB.Delete(&models.Book{}, id).Error; err != nil {
		http.Error(w, "Kitap silinirken hata oluştu", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Kitap silindi: %d", id)
}
