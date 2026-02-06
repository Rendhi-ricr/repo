# Sistem Manajemen Dokumen - CRUD

Sistem manajemen dokumen dengan fitur CRUD (Create, Read, Update, Delete) lengkap.

## 🚀 Fitur

### ✅ CRUD Dokumen Lengkap
- **Create**: Tambah dokumen baru dengan upload file
- **Read**: Lihat daftar semua dokumen dengan search & filter
- **Update**: Edit metadata dokumen dan ganti file (opsional)
- **Delete**: Hapus dokumen beserta file fisiknya

### 🎨 Fitur Tambahan
- **Search**: Cari dokumen berdasarkan judul atau penulis
- **Filter**: Filter dokumen berdasarkan status (draft/publish)
- **Download**: Download file dokumen
- **Status Management**: Ubah status dokumen (draft/publish)
- **Modern UI**: Desain modern dengan gradient, animasi smooth, dan responsive

## 📋 API Endpoints

### Documents Collection
- `GET /api/documents` - Ambil semua dokumen
- `POST /api/documents` - Buat dokumen baru

### Single Document
- `GET /api/documents/:id` - Ambil dokumen berdasarkan ID
- `PUT /api/documents/:id` - Update dokumen
- `DELETE /api/documents/:id` - Hapus dokumen

### Download
- `GET /download/:id` - Download file dokumen

## 🛠️ Cara Menjalankan

### Backend (Go)
```bash
cd backend
go run main.go
```
Server akan berjalan di `http://localhost:8080`

### Frontend (Svelte)
```bash
cd frontend
npm install
npm run dev
```
Frontend akan berjalan di `http://localhost:5173`

## 📊 Database Schema

```sql
CREATE TABLE documents (
    id UUID PRIMARY KEY,
    judul VARCHAR(255) NOT NULL,
    penulis VARCHAR(255) NOT NULL,
    jenis_file VARCHAR(50) NOT NULL,
    file_path VARCHAR(500) NOT NULL,
    status VARCHAR(20) DEFAULT 'draft',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## 🎯 Teknologi

- **Backend**: Go (Golang) dengan PostgreSQL
- **Frontend**: Svelte dengan SPA Router
- **Styling**: Custom CSS dengan gradient modern
- **Font**: Inter (Google Fonts)

## 📝 Catatan

- File yang diupload disimpan di folder `backend/uploads/`
- Saat menghapus dokumen, file fisik juga akan terhapus
- Saat mengupdate dokumen, file lama akan diganti jika file baru diupload
- Semua operasi menggunakan CORS untuk mendukung development

## 🎨 Desain UI

- Gradient colors (Blue to Purple)
- Card-based layout dengan hover effects
- Smooth animations dan transitions
- Responsive design untuk mobile dan desktop
- Custom scrollbar dengan gradient
- Modern typography dengan Inter font
