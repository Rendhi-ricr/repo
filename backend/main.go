package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"repository-un/config"
	"strings"
	"time"

	"github.com/google/uuid"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		enableCORS(w)
		return
	}

	enableCORS(w)

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	judul := r.FormValue("judul")
	penulis := r.FormValue("penulis")
	jenisFile := r.FormValue("jenis_file")

	if judul == "" || penulis == "" || jenisFile == "" {
		http.Error(w, "Metadata tidak lengkap", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "File tidak ditemukan", http.StatusBadRequest)
		return
	}
	defer file.Close()

	ext := filepath.Ext(header.Filename)
storedName := uuid.New().String() + ext
filePath := "uploads/" + storedName

	dst, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Gagal menyimpan file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	io.Copy(dst, file)

	id := uuid.New()

	query := `
		INSERT INTO documents (id, judul, penulis, jenis_file, file_path, status)
		VALUES ($1, $2, $3, $4, $5, 'draft')
	`

	_, err = config.DB.Exec(context.Background(), query,
		id, judul, penulis, jenisFile, filePath,
	)

	if err != nil {
		http.Error(w, "Gagal menyimpan metadata", http.StatusInternalServerError)
		return
	}
	r.ParseMultipartForm(10 << 20) // 10 MB

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{
		"id": "%s",
		"judul": "%s",
		"penulis": "%s",
		"jenis_file": "%s",
		"status": "draft"
	}`, id, judul, penulis, jenisFile)
}




func main() {
	config.ConnectDB()

	http.HandleFunc("/uploads", uploadHandler)
	http.HandleFunc("/documents", listDocumentsHandler)
	http.HandleFunc("/download/", downloadDocumentHandler)
	http.HandleFunc("/publish/", publishDocumentHandler)
	http.HandleFunc("/delete/", deleteDocumentHandler)

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}


type Document struct {
	ID        string    `json:"id"`
	Judul     string    `json:"judul"`
	Penulis   string    `json:"penulis"`
	JenisFile string    `json:"jenis_file"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}


func listDocumentsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		enableCORS(w)
		return
	}

	enableCORS(w)

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	rows, err := config.DB.Query(context.Background(),
		`SELECT id, judul, penulis, jenis_file, status, created_at
		 FROM documents
		 ORDER BY created_at DESC`)
	if err != nil {
		http.Error(w, "Gagal mengambil data", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	documents := []Document{}

	for rows.Next() {
		var d Document
		err := rows.Scan(
			&d.ID,
			&d.Judul,
			&d.Penulis,
			&d.JenisFile,
			&d.Status,
			&d.CreatedAt,
		)
		if err != nil {
			http.Error(w, "Gagal membaca data", http.StatusInternalServerError)
			return
		}
		documents = append(documents, d)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(documents)
}

func publishDocumentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		enableCORS(w)
		return
	}

	enableCORS(w)

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/publish/")
	if id == "" {
		http.Error(w, "ID tidak valid", http.StatusBadRequest)
		return
	}

	query := `
		UPDATE documents
		SET status = 'publish'
		WHERE id = $1
	`

	cmd, err := config.DB.Exec(context.Background(), query, id)
	if err != nil {
		http.Error(w, "Gagal update status", http.StatusInternalServerError)
		return
	}

	if cmd.RowsAffected() == 0 {
		http.Error(w, "Dokumen tidak ditemukan", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{
		"id": "%s",
		"status": "publish"
	}`, id)
}

func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}


func downloadDocumentHandler(w http.ResponseWriter, r *http.Request) {
	
	if r.Method == http.MethodOptions {
		enableCORS(w)
		return
	}
	enableCORS(w)

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Ambil ID dari URL
	path := strings.TrimPrefix(r.URL.Path, "/download/")
	id := strings.TrimSuffix(path, "/download")

	if id == "" {
		http.Error(w, "ID tidak valid", http.StatusBadRequest)
		return
	}

	// Ambil file_path dari DB
	var filePath string
	err := config.DB.QueryRow(
		context.Background(),
		`SELECT file_path FROM documents WHERE id = $1`,
		id,
	).Scan(&filePath)

	if err != nil {
		http.Error(w, "File tidak ditemukan", http.StatusNotFound)
		return
	}

	// Kirim file
	w.Header().Set("Content-Disposition", "attachment")
	http.ServeFile(w, r, filePath)
}

func deleteDocumentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		enableCORS(w)
		w.WriteHeader(http.StatusOK)
		return
	}

	enableCORS(w)

	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/delete/")
	if id == "" {
		http.Error(w, "ID tidak valid", http.StatusBadRequest)
		return
	}

	// 1️⃣ Ambil file_path dari DB
	var filePath string
	err := config.DB.QueryRow(
		context.Background(),
		`SELECT file_path FROM documents WHERE id = $1`,
		id,
	).Scan(&filePath)

	if err != nil {
		http.Error(w, "Dokumen tidak ditemukan", http.StatusNotFound)
		return
	}

	// 2️⃣ Hapus file fisik
	err = os.Remove(filePath)
	if err != nil {
		http.Error(w, "Gagal menghapus file", http.StatusInternalServerError)
		return
	}

	// 3️⃣ Hapus data DB
	_, err = config.DB.Exec(
		context.Background(),
		`DELETE FROM documents WHERE id = $1`,
		id,
	)
	if err != nil {
		http.Error(w, "Gagal menghapus data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message":"Dokumen berhasil dihapus"}`))
}







