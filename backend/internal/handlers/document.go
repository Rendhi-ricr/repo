package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"repository-un/internal/config"
	"repository-un/internal/middleware"
	"repository-un/internal/models"
	"repository-un/internal/utils"

	"github.com/google/uuid"
)

// DocumentsHandler menangani operasi list dan create dokumen
// GET /api/documents - List semua dokumen
// POST /api/documents - Upload dokumen baru
func DocumentsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		middleware.EnableCORS(w)
		return
	}
	middleware.EnableCORS(w)

	switch r.Method {
	case http.MethodGet:
		listDocuments(w, r)
	case http.MethodPost:
		createDocument(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// DocumentByIdHandler menangani operasi pada dokumen tertentu
// GET /api/documents/:id - Get dokumen by ID
// PUT /api/documents/:id - Update dokumen
// DELETE /api/documents/:id - Delete dokumen
func DocumentByIdHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		middleware.EnableCORS(w)
		return
	}
	middleware.EnableCORS(w)

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

// listDocuments mengambil semua dokumen dari database
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

	documents := []models.Document{}

	for rows.Next() {
		var d models.Document
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

// getDocumentById mengambil dokumen berdasarkan ID
func getDocumentById(w http.ResponseWriter, r *http.Request, id string) {
	var d models.Document
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

// createDocument membuat dokumen baru dengan upload file
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

		// Split PDF per halaman untuk preview
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

// updateDocument mengupdate dokumen
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

	// Cek apakah ada file baru
	file, header, err := r.FormFile("file")
	var filePath string

	if err == nil {
		// Ada file baru diupload
		defer file.Close()

		// Ambil path file lama untuk dihapus
		var oldFilePath string
		config.DB.QueryRow(context.Background(),
			`SELECT file_path FROM documents WHERE id = $1`, id).Scan(&oldFilePath)

		// Hapus file lama
		if oldFilePath != "" {
			os.Remove(oldFilePath)
		}

		// Simpan file baru
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

		// Update dengan file baru
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
		// Tidak ada file baru, update metadata saja
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

// deleteDocument menghapus dokumen dan filenya
func deleteDocument(w http.ResponseWriter, r *http.Request, id string) {
	// Ambil file path dari DB
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

	// Hapus file fisik
	if filePath != "" {
		err = os.Remove(filePath)
		if err != nil {
			fmt.Printf("Warning: Failed to delete file %s: %v\n", filePath, err)
		}
	}

	// Hapus direktori split pages jika ada
	splitDir := filepath.Join("uploads", "split", id)
	os.RemoveAll(splitDir)

	// Hapus dari database
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

// DownloadHandler menangani download dokumen
// GET /download/:id
func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		middleware.EnableCORS(w)
		return
	}
	middleware.EnableCORS(w)

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

// PreviewSplitHandler menangani preview halaman PDF
// GET /preview/split/:id/:page.pdf
func PreviewSplitHandler(w http.ResponseWriter, r *http.Request) {
	middleware.EnableCORS(w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Expected path: /preview/split/{id}/{page.pdf}
	relPath := strings.TrimPrefix(r.URL.Path, "/preview/split/")
	if relPath == "" {
		http.Error(w, "Path not specified", http.StatusBadRequest)
		return
	}

	// Cegah directory traversal
	cleanRel := filepath.Clean(relPath)
	if strings.HasPrefix(cleanRel, "..") {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		return
	}

	filePath := filepath.Join("uploads", "split", cleanRel)

	// Cek file exists
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		fmt.Printf("File not found: %s, error: %v\n", filePath, err)
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	// Cek apakah file kosong
	if fileInfo.Size() == 0 {
		fmt.Printf("Warning: File is empty: %s\n", filePath)
		http.Error(w, "File is empty", http.StatusInternalServerError)
		return
	}

	fmt.Printf("Serving file: %s (size: %d bytes)\n", filePath, fileInfo.Size())

	// Set headers
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "inline")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("Cache-Control", "public, max-age=3600")

	http.ServeFile(w, r, filePath)
}

// DocumentPagesHandler mengembalikan list halaman PDF
// GET /api/documents/pages/:id
func DocumentPagesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		middleware.EnableCORS(w)
		return
	}
	middleware.EnableCORS(w)

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

// UploadHandler menangani upload file legacy
// POST /uploads
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		middleware.EnableCORS(w)
		return
	}
	middleware.EnableCORS(w)

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
