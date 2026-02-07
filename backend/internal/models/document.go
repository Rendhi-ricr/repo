package models

import "time"

// Document mewakili struktur dokumen dalam database
type Document struct {
	ID        string    `json:"id"`
	Judul     string    `json:"judul"`
	Penulis   string    `json:"penulis"`
	JenisFile string    `json:"jenis_file"`
	FilePath  string    `json:"file_path,omitempty"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

// CreateDocumentRequest adalah request body untuk membuat dokumen baru
type CreateDocumentRequest struct {
	Title    string `json:"title"`
	Author   string `json:"author"`
	Category string `json:"category"`
	Status   string `json:"status"`
}

// UpdateDocumentRequest adalah request body untuk update dokumen
type UpdateDocumentRequest struct {
	Title    string `json:"title"`
	Author   string `json:"author"`
	Category string `json:"category"`
	Status   string `json:"status"`
}
