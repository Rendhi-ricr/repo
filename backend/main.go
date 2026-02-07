/*
Repository UN - Backend Server
==============================

PERHATIAN: File ini adalah entry point lama.
Gunakan: go run cmd/server/main.go

File ini tetap dipertahankan untuk backward compatibility.
*/

package main

import (
	"fmt"
	"net/http"

	"repository-un/internal/config"
	"repository-un/internal/handlers"
	"repository-un/internal/middleware"
)

func main() {
	// Koneksi ke database
	config.ConnectDB()

	// --- Auth Routes (Public) ---
	http.HandleFunc("/api/auth/login", handlers.LoginHandler)
	http.HandleFunc("/api/auth/register", handlers.RegisterHandler)
	http.HandleFunc("/api/auth/me", middleware.AuthMiddleware(handlers.GetMeHandler))

	// --- User Routes (Admin Only) ---
	http.HandleFunc("/api/users", middleware.AdminMiddleware(handlers.UsersHandler))
	http.HandleFunc("/api/users/", middleware.AdminMiddleware(handlers.UserByIdHandler))

	// --- Document Routes ---
	http.HandleFunc("/uploads", handlers.UploadHandler)
	http.HandleFunc("/api/documents", handlers.DocumentsHandler)
	http.HandleFunc("/api/documents/", handlers.DocumentByIdHandler)
	http.HandleFunc("/api/documents/pages/", handlers.DocumentPagesHandler)

	// --- File Routes ---
	http.HandleFunc("/download/", handlers.DownloadHandler)
	http.HandleFunc("/preview/split/", handlers.PreviewSplitHandler)
	http.Handle("/split/", http.StripPrefix("/split/", http.FileServer(http.Dir("uploads/split"))))

	fmt.Println("========================================")
	fmt.Println("  Repository UN - Backend Server")
	fmt.Println("========================================")
	fmt.Println("Server running at http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}
