package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/CagataySert/library-system/internal/database"
	"github.com/CagataySert/library-system/internal/handlers"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("🚀 Kütüphane Otomasyonu Başlatiliyor...")
	database.Connect()
	database.Migrate() // 🛠️ GORM Migration işlemi

	r := mux.NewRouter()

	// API Rotaları
	r.HandleFunc("/books", handlers.GetBooks).Methods("GET")
	r.HandleFunc("/books", handlers.AddBook).Methods("POST")
	r.HandleFunc("/books/{id}", handlers.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", handlers.DeleteBook).Methods("DELETE")

	fmt.Println("📡 Sunucu 8080 portunda çalişiyor...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
