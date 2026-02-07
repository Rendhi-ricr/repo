/*
Repository UN - Backend Server
==============================
Server API untuk aplikasi repository dokumen.

Struktur folder:
  - cmd/server/     : Entry point aplikasi
  - internal/       : Kode internal aplikasi
    - config/       : Konfigurasi (database, dll)
    - handlers/     : HTTP handlers untuk setiap endpoint
    - middleware/   : Middleware (auth, cors, dll)
    - models/       : Struktur data (User, Document, dll)
    - utils/        : Fungsi utilitas (PDF processing, dll)
  - migrations/     : SQL migration files
  - uploads/        : File yang diupload

Cara menjalankan:
  go run cmd/server/main.go

Atau build dan run:
  go build -o server.exe cmd/server/main.go
  ./server.exe
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

	// ============================================
	// ROUTES
	// ============================================

	// --- Auth Routes (Public) ---
	// Login dan register tidak perlu token
	http.HandleFunc("/api/auth/login", handlers.LoginHandler)
	http.HandleFunc("/api/auth/register", handlers.RegisterHandler)
	http.HandleFunc("/api/auth/me", middleware.AuthMiddleware(handlers.GetMeHandler))

	// --- User Routes (Admin Only) ---
	// Hanya admin yang bisa mengelola user
	http.HandleFunc("/api/users", middleware.AdminMiddleware(handlers.UsersHandler))
	http.HandleFunc("/api/users/", middleware.AdminMiddleware(handlers.UserByIdHandler))

	// --- Document Routes ---
	// CRUD dokumen
	http.HandleFunc("/uploads", handlers.UploadHandler)
	http.HandleFunc("/api/documents", handlers.DocumentsHandler)
	http.HandleFunc("/api/documents/", handlers.DocumentByIdHandler)
	http.HandleFunc("/api/documents/pages/", handlers.DocumentPagesHandler)

	// --- File Routes ---
	// Download dan preview file
	http.HandleFunc("/download/", handlers.DownloadHandler)
	http.HandleFunc("/preview/split/", handlers.PreviewSplitHandler)

	// Static file server untuk split PDF
	http.Handle("/split/", http.StripPrefix("/split/", http.FileServer(http.Dir("uploads/split"))))

	// ============================================
	// START SERVER
	// ============================================
	fmt.Println("========================================")
	fmt.Println("  Repository UN - Backend Server")
	fmt.Println("========================================")
	fmt.Println("Server running at http://localhost:8080")
	fmt.Println("")
	fmt.Println("Endpoints:")
	fmt.Println("  POST /api/auth/login     - Login")
	fmt.Println("  POST /api/auth/register  - Register")
	fmt.Println("  GET  /api/auth/me        - Get current user")
	fmt.Println("  GET  /api/users          - List users (admin)")
	fmt.Println("  GET  /api/documents      - List documents")
	fmt.Println("  POST /api/documents      - Create document")
	fmt.Println("")

	http.ListenAndServe(":8080", nil)
}
