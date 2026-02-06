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

	"repository-un/utils"

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
	http.HandleFunc("/api/documents", documentsHandler)
	http.HandleFunc("/api/documents/", documentByIdHandler)
	http.HandleFunc("/download/", downloadDocumentHandler)
	// Serve split files statically
	http.Handle("/split/", http.StripPrefix("/split/", http.FileServer(http.Dir("uploads/split"))))
	http.HandleFunc("/api/documents/pages/", documentPagesHandler)

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

func documentsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		enableCORS(w)
		return
	}

	enableCORS(w)

	switch r.Method {
	case http.MethodGet:
		listDocuments(w, r)
	case http.MethodPost:
		createDocument(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func listDocuments(w http.ResponseWriter, r *http.Request) {
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

func createDocument(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20) // 10 MB

	judul := r.FormValue("title")
	penulis := r.FormValue("author")
	jenisFile := r.FormValue("category")
	status := r.FormValue("status")

	if judul == "" || penulis == "" || jenisFile == "" {
		http.Error(w, "Metadata tidak lengkap", http.StatusBadRequest)
		return
	}

	if status == "" {
		status = "draft"
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

	// Validasi PDF jika file adalah PDF
	if strings.ToLower(ext) == ".pdf" {
		if err := utils.ValidatePDF(filePath); err != nil {
			os.Remove(filePath) // Hapus file corrupt
			http.Error(w, "File PDF rusak atau tidak valid: "+err.Error(), http.StatusBadRequest)
			return
		}

		// Split PDF per halaman
		splitDir := filepath.Join("uploads", "split", id.String())
		if err := utils.SplitPDF(filePath, splitDir); err != nil {
			fmt.Println("Gagal memecah PDF:", err)
			// Lanjut saja, ini fitur tambahan
		}
	}

	query := `
		INSERT INTO documents (id, judul, penulis, jenis_file, file_path, status)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err = config.DB.Exec(context.Background(), query,
		id, judul, penulis, jenisFile, filePath, status,
	)

	if err != nil {
		http.Error(w, "Gagal menyimpan metadata", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":         id,
		"judul":      judul,
		"penulis":    penulis,
		"jenis_file": jenisFile,
		"status":     status,
	})
}

func documentByIdHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		enableCORS(w)
		return
	}

	enableCORS(w)

	id := strings.TrimPrefix(r.URL.Path, "/api/documents/")
	if id == "" {
		http.Error(w, "ID tidak valid", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		getDocumentById(w, r, id)
	case http.MethodPut:
		updateDocument(w, r, id)
	case http.MethodDelete:
		deleteDocument(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getDocumentById(w http.ResponseWriter, r *http.Request, id string) {
	var d Document
	err := config.DB.QueryRow(context.Background(),
		`SELECT id, judul, penulis, jenis_file, status, created_at
		 FROM documents WHERE id = $1`, id).Scan(
		&d.ID,
		&d.Judul,
		&d.Penulis,
		&d.JenisFile,
		&d.Status,
		&d.CreatedAt,
	)

	if err != nil {
		http.Error(w, "Dokumen tidak ditemukan", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(d)
}

func updateDocument(w http.ResponseWriter, r *http.Request, id string) {
	r.ParseMultipartForm(10 << 20) // 10 MB

	judul := r.FormValue("title")
	penulis := r.FormValue("author")
	jenisFile := r.FormValue("category")
	status := r.FormValue("status")

	if judul == "" || penulis == "" || jenisFile == "" {
		http.Error(w, "Metadata tidak lengkap", http.StatusBadRequest)
		return
	}

	if status == "" {
		status = "draft"
	}

	// Check if new file is uploaded
	file, header, err := r.FormFile("file")
	var filePath string

	if err == nil {
		// New file uploaded
		defer file.Close()

		// Get old file path to delete it
		var oldFilePath string
		config.DB.QueryRow(context.Background(),
			`SELECT file_path FROM documents WHERE id = $1`, id).Scan(&oldFilePath)

		// Delete old file
		if oldFilePath != "" {
			os.Remove(oldFilePath)
		}

		// Save new file
		ext := filepath.Ext(header.Filename)
		storedName := uuid.New().String() + ext
		filePath = "uploads/" + storedName

		dst, err := os.Create(filePath)
		if err != nil {
			http.Error(w, "Gagal menyimpan file", http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		io.Copy(dst, file)

		// Validasi PDF jika file adalah PDF
		if strings.ToLower(ext) == ".pdf" {
			if err := utils.ValidatePDF(filePath); err != nil {
				os.Remove(filePath) // Hapus file corrupt
				http.Error(w, "File PDF rusak atau tidak valid: "+err.Error(), http.StatusBadRequest)
				return
			}

			// Split PDF per halaman
			splitDir := filepath.Join("uploads", "split", id)
			if err := utils.SplitPDF(filePath, splitDir); err != nil {
				fmt.Println("Gagal memecah PDF:", err)
			}
		}

		// Update with new file
		query := `
			UPDATE documents
			SET judul = $1, penulis = $2, jenis_file = $3, status = $4, file_path = $5
			WHERE id = $6
		`
		_, err = config.DB.Exec(context.Background(), query,
			judul, penulis, jenisFile, status, filePath, id)

		if err != nil {
			http.Error(w, "Gagal update dokumen", http.StatusInternalServerError)
			return
		}
	} else {
		// No new file, update metadata only
		query := `
			UPDATE documents
			SET judul = $1, penulis = $2, jenis_file = $3, status = $4
			WHERE id = $5
		`
		_, err = config.DB.Exec(context.Background(), query,
			judul, penulis, jenisFile, status, id)

		if err != nil {
			http.Error(w, "Gagal update dokumen", http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":         id,
		"judul":      judul,
		"penulis":    penulis,
		"jenis_file": jenisFile,
		"status":     status,
	})
}

func deleteDocument(w http.ResponseWriter, r *http.Request, id string) {
	// Get file path from DB
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

	// Delete physical file (Best Effort)
	if filePath != "" {
		err = os.Remove(filePath)
		if err != nil {
			// Log error but continue if file just doesn't exist
			fmt.Printf("Warning: Failed to delete file %s: %v\n", filePath, err)
		}
	}

	// Delete split pages directory if exists
	splitDir := filepath.Join("uploads", "split", id)
	os.RemoveAll(splitDir)

	// Delete from DB
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

func documentPagesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		enableCORS(w)
		return
	}
	enableCORS(w)

	// Parse ID: /api/documents/pages/{id}
	id := strings.TrimPrefix(r.URL.Path, "/api/documents/pages/")
	if id == "" {
		http.Error(w, "ID tidak valid", http.StatusBadRequest)
		return
	}

	splitDir := filepath.Join("uploads", "split", id)
	files, err := os.ReadDir(splitDir)
	if err != nil {
		// Dir tidak ditemukan = belum di-split atau bukan PDF
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("[]"))
		return
	}

	var pages []string
	for _, f := range files {
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".pdf") {
			pages = append(pages, f.Name())
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pages)
}
